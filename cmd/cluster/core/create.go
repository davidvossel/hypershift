package core

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/wait"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"

	apifixtures "github.com/openshift/hypershift/api/fixtures"
	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	"github.com/openshift/hypershift/cmd/cluster/common"
	"github.com/openshift/hypershift/cmd/log"
	"github.com/openshift/hypershift/cmd/util"
	"github.com/openshift/hypershift/cmd/version"
	hyperapi "github.com/openshift/hypershift/support/api"
)

// ApplyPlatformSpecifics can be used to create platform specific values as well as enriching the fixure with additional values
type ApplyPlatformSpecifics = func(ctx context.Context, fixture *apifixtures.ExampleOptions, options *CreateOptions) error

type CreateOptions struct {
	Annotations                      []string
	AutoRepair                       bool
	ControlPlaneAvailabilityPolicy   string
	ControlPlaneOperatorImage        string
	EtcdStorageClass                 string
	FIPS                             bool
	GenerateSSH                      bool
	InfrastructureAvailabilityPolicy string
	InfrastructureJSON               string
	InfraID                          string
	Name                             string
	Namespace                        string
	BaseDomain                       string
	NetworkType                      string
	NodePoolReplicas                 int32
	PullSecretFile                   string
	ReleaseImage                     string
	Render                           bool
	SSHKeyFile                       string
	ServiceCIDR                      string
	PodCIDR                          string
	NonePlatform                     NonePlatformCreateOptions
	KubevirtPlatform                 KubevirtPlatformCreateOptions
	AWSPlatform                      AWSPlatformOptions
	AgentPlatform                    AgentPlatformCreateOptions
	AzurePlatform                    AzurePlatformOptions
	Wait                             bool
	Timeout                          time.Duration
}

type AgentPlatformCreateOptions struct {
	ServicesPublishOpts *common.ServicesPublishOptions
	AgentNamespace      string
}

type NonePlatformCreateOptions struct {
	ServicesPublishOpts *common.ServicesPublishOptions
}

type KubevirtPlatformCreateOptions struct {
	ServicesPublishOpts *common.ServicesPublishOptions
	Memory              string
	Cores               uint32
	ContainerDiskImage  string
}

type AWSPlatformOptions struct {
	AWSCredentialsFile string
	AdditionalTags     []string
	IAMJSON            string
	InstanceType       string
	IssuerURL          string
	PrivateZoneID      string
	PublicZoneID       string
	Region             string
	RootVolumeIOPS     int64
	RootVolumeSize     int64
	RootVolumeType     string
	EndpointAccess     string
	Zones              []string
	EtcdKMSKeyARN      string
}

type AzurePlatformOptions struct {
	CredentialsFile string
	Location        string
	InstanceType    string
}

func createCommonFixture(opts *CreateOptions) (*apifixtures.ExampleOptions, error) {
	if len(opts.ReleaseImage) == 0 {
		defaultVersion, err := version.LookupDefaultOCPVersion()
		if err != nil {
			return nil, fmt.Errorf("release image is required when unable to lookup default OCP version: %w", err)
		}
		opts.ReleaseImage = defaultVersion.PullSpec
	}

	annotations := map[string]string{}
	for _, s := range opts.Annotations {
		pair := strings.SplitN(s, "=", 2)
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid annotation: %s", s)
		}
		k, v := pair[0], pair[1]
		annotations[k] = v
	}

	if len(opts.ControlPlaneOperatorImage) > 0 {
		annotations[hyperv1.ControlPlaneOperatorImageAnnotation] = opts.ControlPlaneOperatorImage
	}

	pullSecret, err := ioutil.ReadFile(opts.PullSecretFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read pull secret file: %w", err)
	}
	var sshKey, sshPrivateKey []byte
	if len(opts.SSHKeyFile) > 0 {
		if opts.GenerateSSH {
			return nil, fmt.Errorf("--generate-ssh and --ssh-key cannot be specified together")
		}
		key, err := ioutil.ReadFile(opts.SSHKeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read ssh key file: %w", err)
		}
		sshKey = key
	} else if opts.GenerateSSH {
		sshKey, sshPrivateKey, err = generateSSHKeys()
		if err != nil {
			return nil, fmt.Errorf("failed to generate ssh keys: %w", err)
		}
	}

	return &apifixtures.ExampleOptions{
		InfraID:                          opts.InfraID,
		Annotations:                      annotations,
		AutoRepair:                       opts.AutoRepair,
		ControlPlaneAvailabilityPolicy:   hyperv1.AvailabilityPolicy(opts.ControlPlaneAvailabilityPolicy),
		FIPS:                             opts.FIPS,
		InfrastructureAvailabilityPolicy: hyperv1.AvailabilityPolicy(opts.InfrastructureAvailabilityPolicy),
		Namespace:                        opts.Namespace,
		Name:                             opts.Name,
		NetworkType:                      hyperv1.NetworkType(opts.NetworkType),
		NodePoolReplicas:                 opts.NodePoolReplicas,
		PullSecret:                       pullSecret,
		ReleaseImage:                     opts.ReleaseImage,
		SSHPrivateKey:                    sshPrivateKey,
		SSHPublicKey:                     sshKey,
		EtcdStorageClass:                 opts.EtcdStorageClass,
		ServiceCIDR:                      opts.ServiceCIDR,
		PodCIDR:                          opts.PodCIDR,
	}, nil
}

func generateSSHKeys() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEMBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}
	privatePEM := pem.EncodeToMemory(&privatePEMBlock)

	publicRSAKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicBytes := ssh.MarshalAuthorizedKey(publicRSAKey)

	return publicBytes, privatePEM, nil
}

func apply(ctx context.Context, exampleOptions *apifixtures.ExampleOptions, render bool, waitForRollout bool) error {

	exampleObjects := exampleOptions.Resources().AsObjects()
	switch {
	case render:
		for _, object := range exampleObjects {
			err := hyperapi.YamlSerializer.Encode(object, os.Stdout)
			if err != nil {
				return fmt.Errorf("failed to encode objects: %w", err)
			}
			fmt.Println("---")
		}
	default:
		client := util.GetClientOrDie()
		var hostedCluster *hyperv1.HostedCluster
		for _, object := range exampleObjects {
			key := crclient.ObjectKeyFromObject(object)
			object.SetLabels(map[string]string{util.AutoInfraLabelName: exampleOptions.InfraID})
			var err error
			if object.GetObjectKind().GroupVersionKind().Kind == "HostedCluster" {
				hostedCluster = &hyperv1.HostedCluster{ObjectMeta: metav1.ObjectMeta{Namespace: object.GetNamespace(), Name: object.GetName()}}
				err = client.Create(ctx, object)
			} else {
				err = client.Patch(ctx, object, crclient.Apply, crclient.ForceOwnership, crclient.FieldOwner("hypershift-cli"))
			}
			if err != nil {
				return fmt.Errorf("failed to apply object %q: %w", key, err)
			}
			log.Log.Info("Applied Kube resource", "kind", object.GetObjectKind().GroupVersionKind().Kind, "namespace", key.Namespace, "name", key.Name)
		}

		if waitForRollout {
			log.Log.Info("Waiting for cluster rollout")
			return wait.PollInfiniteWithContext(ctx, 30*time.Second, func(ctx context.Context) (bool, error) {
				hostedCluster := hostedCluster.DeepCopy()
				if err := client.Get(ctx, crclient.ObjectKeyFromObject(hostedCluster), hostedCluster); err != nil {
					return false, fmt.Errorf("failed to get hostedcluster %s: %w", crclient.ObjectKeyFromObject(hostedCluster), err)
				}
				rolledOut := len(hostedCluster.Status.Version.History) > 0 && hostedCluster.Status.Version.History[0].CompletionTime != nil
				if !rolledOut {
					log.Log.Info("Cluster rollout not finished yet, checking again in 30 seconds...")
				}
				return rolledOut, nil
			})
		}

		return nil
	}
	return nil
}

func Validate(ctx context.Context, opts *CreateOptions) error {
	if !opts.Render {
		client := util.GetClientOrDie()
		// Validate HostedCluster with this name doesn't exists in the namespace
		cluster := &hyperv1.HostedCluster{ObjectMeta: metav1.ObjectMeta{Namespace: opts.Namespace, Name: opts.Name}}
		if err := client.Get(ctx, crclient.ObjectKeyFromObject(cluster), cluster); err == nil {
			return fmt.Errorf("hostedcluster %s already exists", crclient.ObjectKeyFromObject(cluster))
		} else if !apierrors.IsNotFound(err) {
			return fmt.Errorf("hostedcluster doesn't exist validation failed with error: %w", err)
		}
	}

	return nil
}

func CreateCluster(ctx context.Context, opts *CreateOptions, platformSpecificApply ApplyPlatformSpecifics) error {
	if opts.Wait && opts.NodePoolReplicas < 1 {
		return errors.New("--wait requires --node-pool-replicas > 0")
	}

	if opts.InfraID == "" {
		opts.InfraID = fmt.Sprintf("%s-%s", opts.Name, utilrand.String(5))
	}

	exampleOptions, err := createCommonFixture(opts)
	if err != nil {
		return err
	}

	// Apply platform specific options and create platform specific resources
	if err := platformSpecificApply(ctx, exampleOptions, opts); err != nil {
		return err
	}

	return apply(ctx, exampleOptions, opts.Render, opts.Wait)
}
