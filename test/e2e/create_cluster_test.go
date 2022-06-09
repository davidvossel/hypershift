//go:build e2e
// +build e2e

package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	e2eutil "github.com/openshift/hypershift/test/e2e/util"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// TestCreateCluster implements a test that creates a cluster with the code under test
// vs upgrading to the code under test as TestUpgradeControlPlane does.
func TestCreateCluster(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	ctx, cancel := context.WithCancel(testContext)
	defer cancel()

	client, err := e2eutil.GetClient()
	g.Expect(err).NotTo(HaveOccurred(), "failed to get k8s client")

	clusterOpts := globalOpts.DefaultClusterOptions()
	clusterOpts.ControlPlaneAvailabilityPolicy = string(hyperv1.SingleReplica)

	hostedCluster := e2eutil.CreateCluster(t, ctx, client, &clusterOpts, globalOpts.Platform, globalOpts.ArtifactDir)

	// Sanity check the cluster by waiting for the nodes to report ready
	t.Logf("Waiting for guest client to become available")
	guestClient := e2eutil.WaitForGuestClient(t, testContext, client, hostedCluster)

	// Wait for Nodes to be Ready
	numNodes := int32(globalOpts.configurableClusterOptions.NodePoolReplicas * len(clusterOpts.AWSPlatform.Zones))
	e2eutil.WaitForNReadyNodes(t, testContext, guestClient, numNodes)

	// Wait for the rollout to be complete
	t.Logf("Waiting for cluster rollout. Image: %s", globalOpts.LatestReleaseImage)
	e2eutil.WaitForImageRollout(t, testContext, client, guestClient, hostedCluster, globalOpts.LatestReleaseImage)
	err = client.Get(testContext, crclient.ObjectKeyFromObject(hostedCluster), hostedCluster)
	g.Expect(err).NotTo(HaveOccurred(), "failed to get hostedcluster")

	// TODO BEGIN EXPERIMENT
	guestNodes := &corev1.NodeList{}

	err = guestClient.List(ctx, guestNodes)
	g.Expect(err).NotTo(HaveOccurred(), "failed to list guest nodes")

	// CREATE ECHO POD AND NODEPORT on Guest
	nodePort := 32700

	echoPod := &corev1.Pod{}
	echoPod.Name = "http-echo"
	echoPod.Namespace = "default"
	echoPod.ObjectMeta.Labels = map[string]string{
		"app": "http-echo",
	}
	echoPod.Spec.Containers = []corev1.Container{
		{
			Name:  "echo-pod",
			Image: "hashicorp/http-echo:0.2.3",
			Args: []string{
				"-text=echo",
			},
		},
	}

	service := &corev1.Service{}
	service.Name = "echo-service"
	service.Namespace = "default"
	service.Spec = corev1.ServiceSpec{
		Type: "NodePort",
		Selector: map[string]string{
			"app": "http-echo",
		},
		Ports: []corev1.ServicePort{
			{
				// default port for the echo container
				Port:     5678,
				NodePort: int32(nodePort),
			},
		},
	}

	// CREATE CURL JOB On Guest from HostNet
	template := corev1.PodSpec{
		HostNetwork:   true,
		Containers:    []corev1.Container{},
		RestartPolicy: "Never",
		// This makes sure we schedule the curl pod on a node separate from the echo pod.
		Affinity: &corev1.Affinity{
			PodAntiAffinity: &corev1.PodAntiAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: []corev1.PodAffinityTerm{
					{
						LabelSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      "app",
									Values:   []string{"http-echo"},
									Operator: metav1.LabelSelectorOpIn,
								},
							},
						},
						TopologyKey: "kubernetes.io/hostname",
					},
				},
			},
		},
	}
	for i, node := range guestNodes.Items {
		ip := ""
		for _, addr := range node.Status.Addresses {
			if addr.Type == corev1.NodeInternalIP {
				ip = addr.Address
			}
		}
		g.Expect(ip).NotTo(Equal(""), fmt.Sprintf("no internal ip found for guest node %s", node.Name))

		template.Containers = append(template.Containers, corev1.Container{
			Name:  fmt.Sprintf("curl-%d", i),
			Image: "fedora:35",
			Command: []string{
				"curl",
				fmt.Sprintf("%s:%d", ip, nodePort),
			},
		})
	}
	backoff := int32(4)
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "nodeport-curl-test-hostnetwork",
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backoff,
			Template: corev1.PodTemplateSpec{
				Spec: template,
			},
		},
	}

	err = guestClient.Create(ctx, echoPod)
	g.Expect(err).NotTo(HaveOccurred(), "failed to create echo pod")
	err = guestClient.Create(ctx, service)
	g.Expect(err).NotTo(HaveOccurred(), "failed to create echo nodeport")
	err = guestClient.Create(ctx, job)
	g.Expect(err).NotTo(HaveOccurred(), "failed to create echo curl test job")

	err = wait.PollImmediateWithContext(ctx, 10*time.Second, 5*time.Minute, func(ctx context.Context) (done bool, err error) {
		updatedJob := &batchv1.Job{}
		err = guestClient.Get(ctx, crclient.ObjectKeyFromObject(job), updatedJob)
		if err != nil {
			t.Errorf("Failed to get job: %v", err)
			return false, nil
		}

		for _, condition := range updatedJob.Status.Conditions {
			if condition.Type == batchv1.JobComplete && condition.Status == corev1.ConditionTrue {
				t.Logf("Guest NodePort connectivity test passed from guest host network")
				return true, nil
			}
		}

		return false, nil
	})
	g.Expect(err).NotTo(HaveOccurred(), "curl pod failed connectivity tests for NodePort from guest node host network")
	// TODO END EXPERIMENT

	e2eutil.EnsureNodeCountMatchesNodePoolReplicas(t, testContext, client, guestClient, hostedCluster.Namespace)
	e2eutil.EnsureNoCrashingPods(t, ctx, client, hostedCluster)
}

func TestNoneCreateCluster(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)

	ctx, cancel := context.WithCancel(testContext)
	defer cancel()

	client, err := e2eutil.GetClient()
	g.Expect(err).NotTo(HaveOccurred(), "failed to get k8s client")

	clusterOpts := globalOpts.DefaultClusterOptions()
	clusterOpts.ControlPlaneAvailabilityPolicy = "SingleReplica"

	hostedCluster := e2eutil.CreateCluster(t, ctx, client, &clusterOpts, hyperv1.NonePlatform, globalOpts.ArtifactDir)

	// Wait for the rollout to be reported complete
	t.Logf("Waiting for cluster rollout. Image: %s", globalOpts.LatestReleaseImage)
	// Since the None platform has no workers, CVO will not have expectations set,
	// which in turn means that the ClusterVersion object will never be populated.
	// Therefore only test if the control plane comes up (etc, apiserver, ...)
	e2eutil.WaitForConditionsOnHostedControlPlane(t, ctx, client, hostedCluster, globalOpts.LatestReleaseImage)

	// etcd restarts for me once always and apiserver two times before running stable
	// e2eutil.EnsureNoCrashingPods(t, ctx, client, hostedCluster)
}
