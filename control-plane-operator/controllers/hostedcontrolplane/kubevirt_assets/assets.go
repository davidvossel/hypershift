package kubevirt_assets

import (
	"bytes"
	"embed"
	"io"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/hoisie/mustache"
	crclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func ptrbool(p bool) *bool {
	return &p
}

func ptrHostPathType(p corev1.HostPathType) *corev1.HostPathType {
	return &p
}

func ptrMountPropagationMode(p corev1.MountPropagationMode) *corev1.MountPropagationMode {
	return &p
}

// var KubevirtCsiNodeServiceAccount = &corev1.ServiceAccount{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ServiceAccount",
// 		APIVersion: "v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi-node-sa",
// 		Namespace: "openshift-cluster-csi-drivers",
// 	},
// }

// var KubevirtCsiControllerServiceAccount = &corev1.ServiceAccount{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ServiceAccount",
// 		APIVersion: "v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi-controller-sa",
// 		Namespace: "openshift-cluster-csi-drivers",
// 	},
// }

// var KubevirtCsiControllerClusterRole = &rbacv1.ClusterRole{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ClusterRole",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "kubevirt-csi-controller-cr",
// 	},
// 	Rules: []rbacv1.PolicyRule{
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"list",
// 				"create",
// 			},
// 			APIGroups: []string{
// 				"apiextensions.k8s.io",
// 			},
// 			Resources: []string{
// 				"customresourcedefinitions",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"create",
// 				"delete",
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"persistentvolumes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"secrets",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"persistentvolumeclaims",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"persistentvolumeclaims/status",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"nodes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumeattachments",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"storageclasses",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"create",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"csi.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"csidrivers",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"list",
// 				"watch",
// 				"create",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"events",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"snapshot.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumesnapshotclasses",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"create",
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"delete",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"snapshot.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumesnapshotcontents",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"snapshot.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumesnapshots",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"update",
// 			},
// 			APIGroups: []string{
// 				"snapshot.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumesnapshots/status",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumeattachments/status",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"csinodes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"use",
// 			},
// 			APIGroups: []string{
// 				"security.openshift.io",
// 			},
// 			Resources: []string{
// 				"securitycontextconstraints",
// 			},
// 			ResourceNames: []string{
// 				"privileged",
// 			},
// 		},
// 	},
// }

// var KubevirtCsiNodeClusterRole = &rbacv1.ClusterRole{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ClusterRole",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "kubevirt-csi-node-cr",
// 	},
// 	Rules: []rbacv1.PolicyRule{
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"create",
// 				"delete",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"persistentvolumes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"persistentvolumeclaims",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"storageclasses",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"nodes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"csi.storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"csinodeinfos",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"csinodes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumeattachments",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"list",
// 				"watch",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{
// 				"storage.k8s.io",
// 			},
// 			Resources: []string{
// 				"volumeattachments/status",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"list",
// 				"watch",
// 				"create",
// 				"update",
// 				"patch",
// 			},
// 			APIGroups: []string{},
// 			Resources: []string{
// 				"events",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"use",
// 			},
// 			APIGroups: []string{
// 				"security.openshift.io",
// 			},
// 			Resources: []string{
// 				"securitycontextconstraints",
// 			},
// 			ResourceNames: []string{
// 				"privileged",
// 			},
// 		},
// 	},
// }

// var KubevirtCsiControllerBinding = &rbacv1.ClusterRoleBinding{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ClusterRoleBinding",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "kubevirt-csi-controller-binding",
// 	},
// 	Subjects: []rbacv1.Subject{
// 		rbacv1.Subject{
// 			Kind:      "ServiceAccount",
// 			Name:      "kubevirt-csi-controller-sa",
// 			Namespace: "openshift-cluster-csi-drivers",
// 		},
// 	},
// 	RoleRef: rbacv1.RoleRef{
// 		APIGroup: "rbac.authorization.k8s.io",
// 		Kind:     "ClusterRole",
// 		Name:     "kubevirt-csi-controller-cr",
// 	},
// }

// var KubevirtCsiNodeBinding = &rbacv1.ClusterRoleBinding{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ClusterRoleBinding",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "kubevirt-csi-node-binding",
// 	},
// 	Subjects: []rbacv1.Subject{
// 		rbacv1.Subject{
// 			Kind:      "ServiceAccount",
// 			Name:      "kubevirt-csi-node-sa",
// 			Namespace: "openshift-cluster-csi-drivers",
// 		},
// 	},
// 	RoleRef: rbacv1.RoleRef{
// 		APIGroup: "rbac.authorization.k8s.io",
// 		Kind:     "ClusterRole",
// 		Name:     "kubevirt-csi-node-cr",
// 	},
// }

// var InfraServiceAccount = &corev1.ServiceAccount{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ServiceAccount",
// 		APIVersion: "v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi",
// 		Namespace: "clusters-isaac-infra",
// 	},
// }

// var KubevirtCsiRole = &rbacv1.Role{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "Role",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi",
// 		Namespace: "clusters-isaac-infra",
// 	},
// 	Rules: []rbacv1.PolicyRule{
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"get",
// 				"create",
// 				"delete",
// 			},
// 			APIGroups: []string{
// 				"cdi.kubevirt.io",
// 			},
// 			Resources: []string{
// 				"datavolumes",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"list",
// 			},
// 			APIGroups: []string{
// 				"kubevirt.io",
// 			},
// 			Resources: []string{
// 				"virtualmachineinstances",
// 			},
// 		},
// 		rbacv1.PolicyRule{
// 			Verbs: []string{
// 				"update",
// 			},
// 			APIGroups: []string{
// 				"subresources.kubevirt.io",
// 			},
// 			Resources: []string{
// 				"virtualmachineinstances/addvolume",
// 				"virtualmachineinstances/removevolume",
// 			},
// 		},
// 	},
// }

// var KubevirtCsiRoleBinding = &rbacv1.RoleBinding{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "RoleBinding",
// 		APIVersion: "rbac.authorization.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi",
// 		Namespace: "clusters-isaac-infra",
// 	},
// 	Subjects: []rbacv1.Subject{
// 		rbacv1.Subject{
// 			Kind:      "ServiceAccount",
// 			Name:      "kubevirt-csi",
// 			Namespace: "clusters-isaac-infra",
// 		},
// 	},
// 	RoleRef: rbacv1.RoleRef{
// 		APIGroup: "rbac.authorization.k8s.io",
// 		Kind:     "Role",
// 		Name:     "kubevirt-csi",
// 	},
// }

// func ptrint32(p int32) *int32 {
// 	return &p
// }

// var KubevirtCsiControllerDeployment = &appsv1.Deployment{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "Deployment",
// 		APIVersion: "apps/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi-controller",
// 		Namespace: "clusters-isaac-infra",
// 	},
// 	Spec: appsv1.DeploymentSpec{
// 		Replicas: ptrint32(1),
// 		Selector: &metav1.LabelSelector{
// 			MatchLabels: map[string]string{
// 				"app": "kubevirt-csi-driver",
// 			},
// 		},
// 		Template: corev1.PodTemplateSpec{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Labels: map[string]string{
// 					"app": "kubevirt-csi-driver",
// 				},
// 			},
// 			Spec: corev1.PodSpec{
// 				Volumes: []corev1.Volume{
// 					corev1.Volume{
// 						Name: "socket-dir",
// 						VolumeSource: corev1.VolumeSource{
// 							EmptyDir: &corev1.EmptyDirVolumeSource{},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "tenantcluster",
// 						VolumeSource: corev1.VolumeSource{
// 							Secret: &corev1.SecretVolumeSource{
// 								SecretName: "tenant-controller-kubeconfig",
// 							},
// 						},
// 					},
// 				},
// 				Containers: []corev1.Container{
// 					corev1.Container{
// 						Name:  "csi-driver",
// 						Image: "quay.io/isaacdorfman/kubevirt-csi-driver:latest",
// 						Args: []string{
// 							"--endpoint=$(CSI_ENDPOINT)",
// 							"--namespace=kubevirt-csi-driver",
// 							"--infra-cluster-namespace=$(INFRACLUSTER_NAMESPACE)",
// 							"--tenant-cluster-kubeconfig=/var/run/secrets/tenantcluster/kubeconfig",
// 							"--infra-cluster-labels=$(INFRACLUSTER_LABELS)",
// 							"--run-controller-service",
// 							"--v=5",
// 						},
// 						Ports: []corev1.ContainerPort{
// 							corev1.ContainerPort{
// 								Name:          "healthz",
// 								HostPort:      0,
// 								ContainerPort: 10301,
// 								Protocol:      corev1.Protocol("TCP"),
// 							},
// 						},
// 						Env: []corev1.EnvVar{
// 							corev1.EnvVar{
// 								Name:  "CSI_ENDPOINT",
// 								Value: "unix:///var/lib/csi/sockets/pluginproxy/csi.sock",
// 							},
// 							corev1.EnvVar{
// 								Name: "KUBE_NODE_NAME",
// 								ValueFrom: &corev1.EnvVarSource{
// 									FieldRef: &corev1.ObjectFieldSelector{
// 										FieldPath: "spec.nodeName",
// 									},
// 								},
// 							},
// 							corev1.EnvVar{
// 								Name: "INFRACLUSTER_NAMESPACE",
// 								ValueFrom: &corev1.EnvVarSource{
// 									ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
// 										LocalObjectReference: corev1.LocalObjectReference{
// 											Name: "driver-config",
// 										},
// 										Key: "infraClusterNamespace",
// 									},
// 								},
// 							},
// 							corev1.EnvVar{
// 								Name: "INFRACLUSTER_LABELS",
// 								ValueFrom: &corev1.EnvVarSource{
// 									ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
// 										LocalObjectReference: corev1.LocalObjectReference{
// 											Name: "driver-config",
// 										},
// 										Key: "infraClusterLabels",
// 									},
// 								},
// 							},
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(10, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(52428800, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "socket-dir",
// 								MountPath: "/var/lib/csi/sockets/pluginproxy/",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "tenantcluster",
// 								MountPath: "/var/run/secrets/tenantcluster",
// 							},
// 						},
// 						ImagePullPolicy: corev1.PullPolicy("Always"),
// 					},
// 					corev1.Container{
// 						Name:  "csi-provisioner",
// 						Image: "quay.io/openshift/origin-csi-external-provisioner:latest",
// 						Args: []string{
// 							"--csi-address=$(ADDRESS)",
// 							"--default-fstype=ext4",
// 							"--v=5",
// 							"--kubeconfig",
// 							"/var/run/secrets/tenantcluster/kubeconfig",
// 						},
// 						Env: []corev1.EnvVar{
// 							corev1.EnvVar{
// 								Name:  "ADDRESS",
// 								Value: "/var/lib/csi/sockets/pluginproxy/csi.sock",
// 							},
// 						},
// 						Resources: corev1.ResourceRequirements{},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "socket-dir",
// 								MountPath: "/var/lib/csi/sockets/pluginproxy/",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "tenantcluster",
// 								MountPath: "/var/run/secrets/tenantcluster",
// 							},
// 						},
// 					},
// 					corev1.Container{
// 						Name:  "csi-attacher",
// 						Image: "quay.io/openshift/origin-csi-external-attacher:latest",
// 						Args: []string{
// 							"--csi-address=$(ADDRESS)",
// 							"--v=5",
// 							"--kubeconfig",
// 							"/var/run/secrets/tenantcluster/kubeconfig",
// 						},
// 						Env: []corev1.EnvVar{
// 							corev1.EnvVar{
// 								Name:  "ADDRESS",
// 								Value: "/var/lib/csi/sockets/pluginproxy/csi.sock",
// 							},
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(10, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(52428800, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "socket-dir",
// 								MountPath: "/var/lib/csi/sockets/pluginproxy/",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "tenantcluster",
// 								MountPath: "/var/run/secrets/tenantcluster",
// 							},
// 						},
// 					},
// 					corev1.Container{
// 						Name:  "csi-liveness-probe",
// 						Image: "quay.io/openshift/origin-csi-livenessprobe:latest",
// 						Args: []string{
// 							"--csi-address=/csi/csi.sock",
// 							"--probe-timeout=3s",
// 							"--health-port=10301",
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(10, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(52428800, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "socket-dir",
// 								MountPath: "/csi",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "tenantcluster",
// 								MountPath: "/var/run/secrets/tenantcluster",
// 							},
// 						},
// 					},
// 				},
// 				NodeSelector:             map[string]string{},
// 				DeprecatedServiceAccount: "kubevirt-csi",
// 				HostNetwork:              true,
// 				Tolerations: []corev1.Toleration{
// 					corev1.Toleration{
// 						Key:      "CriticalAddonsOnly",
// 						Operator: corev1.TolerationOperator("Exists"),
// 					},
// 					corev1.Toleration{
// 						Key:      "node-role.kubernetes.io/master",
// 						Operator: corev1.TolerationOperator("Exists"),
// 						Effect:   corev1.TaintEffect("NoSchedule"),
// 					},
// 				},
// 				PriorityClassName: "system-cluster-critical",
// 			},
// 		},
// 		Strategy:        appsv1.DeploymentStrategy{},
// 		MinReadySeconds: 0,
// 	},
// }

// var TenantCsiNamespace = &corev1.Namespace{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "Namespace",
// 		APIVersion: "v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "openshift-cluster-csi-drivers",
// 	},
// 	Spec: corev1.NamespaceSpec{},
// }

// var TenantStorageClass = &storagev1.StorageClass{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "StorageClass",
// 		APIVersion: "storage.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "kubevirt",
// 		Annotations: map[string]string{
// 			"storageclass.kubernetes.io/is-default-class": "true",
// 		},
// 	},
// 	Provisioner: "csi.kubevirt.io",
// 	Parameters: map[string]string{
// 		"bus":                   "scsi",
// 		"infraStorageClassName": "standard",
// 	},
// }

// var CsiDriverCRD = &apiextensionsv1.CustomResourceDefinition{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "CustomResourceDefinition",
// 		APIVersion: "apiextensions.k8s.io/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name: "clustercsidrivers.operator.openshift.io",
// 	},
// 	Spec: apiextensionsv1.CustomResourceDefinitionSpec{
// 		Group: "operator.openshift.io",
// 		Names: apiextensionsv1.CustomResourceDefinitionNames{
// 			Plural:   "clustercsidrivers",
// 			Singular: "clustercsidriver",
// 			Kind:     "ClusterCSIDriver",
// 		},
// 		Scope: apiextensionsv1.ResourceScope("Cluster"),
// 		Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
// 			apiextensionsv1.CustomResourceDefinitionVersion{
// 				Name:    "v1",
// 				Served:  true,
// 				Storage: true,
// 				Schema: &apiextensionsv1.CustomResourceValidation{
// 					OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
// 						Description: `ClusterCSIDriver object allows management and configuration of
// a CSI driver operator installed by default in OpenShift.`,
// 						Type: "object",
// 						Required: []string{
// 							"spec",
// 						},
// 						Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 							"apiVersion": apiextensionsv1.JSONSchemaProps{
// 								Description: `APIVersion defines the versioned schema of this representation
// of an object. Servers should convert recognized schemas to the latest
// internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
// 								Type: "string",
// 							},
// 							"kind": apiextensionsv1.JSONSchemaProps{
// 								Description: `Kind is a string value representing the REST resource this
// object represents. Servers may infer this from the endpoint the client
// submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
// 								Type: "string",
// 							},
// 							"metadata": apiextensionsv1.JSONSchemaProps{
// 								Type: "object",
// 								Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 									"name": apiextensionsv1.JSONSchemaProps{
// 										Type: "string",
// 										Enum: []apiextensionsv1.JSON{
// 											apiextensionsv1.JSON{
// 												Raw: []uint8{
// 													34,
// 													99,
// 													115,
// 													105,
// 													46,
// 													107,
// 													117,
// 													98,
// 													101,
// 													118,
// 													105,
// 													114,
// 													116,
// 													46,
// 													105,
// 													111,
// 													34,
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 							"spec": apiextensionsv1.JSONSchemaProps{
// 								Description: "spec holds user settable values for configuration",
// 								Type:        "object",
// 								Required: []string{
// 									"driverConfig",
// 								},
// 								Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 									"driverConfig": apiextensionsv1.JSONSchemaProps{
// 										Description: "CSIDriverConfig is the CSI driver specific configuration",
// 										Type:        "object",
// 										Required: []string{
// 											"driverName",
// 										},
// 										Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 											"driverName": apiextensionsv1.JSONSchemaProps{
// 												Description: "DriverName holds the name of the CSI driver",
// 												Type:        "string",
// 												Enum: []apiextensionsv1.JSON{
// 													apiextensionsv1.JSON{
// 														Raw: []uint8{
// 															34,
// 															99,
// 															115,
// 															105,
// 															46,
// 															107,
// 															117,
// 															98,
// 															101,
// 															118,
// 															105,
// 															114,
// 															116,
// 															46,
// 															105,
// 															111,
// 															34,
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 									"logLevel": apiextensionsv1.JSONSchemaProps{
// 										Description: `logLevel is an intent based logging for an overall component.  It
// does not give fine grained control, but it is a simple way to manage
// coarse grained logging choices that operators have to interpret for
// their operands.`,
// 										Type: "string",
// 									},
// 									"managementState": apiextensionsv1.JSONSchemaProps{
// 										Description: `managementState indicates whether and how the operator
// should manage the component`,
// 										Type:    "string",
// 										Pattern: "^(Managed|Unmanaged|Force|Removed)$",
// 									},
// 									"observedConfig": apiextensionsv1.JSONSchemaProps{
// 										Description: `observedConfig holds a sparse config that controller has
// observed from the cluster state.  It exists in spec because it is
// an input to the level for the operator`,
// 										Type:                   "object",
// 										Nullable:               true,
// 										XPreserveUnknownFields: ptrbool(true),
// 									},
// 									"operatorLogLevel": apiextensionsv1.JSONSchemaProps{
// 										Description: `operatorLogLevel is an intent based logging for the operator
// itself.  It does not give fine grained control, but it is a simple
// way to manage coarse grained logging choices that operators have to
// interpret for themselves.`,
// 										Type: "string",
// 									},
// 									"unsupportedConfigOverrides": apiextensionsv1.JSONSchemaProps{
// 										Description: `unsupportedConfigOverrides holds a sparse config that
// will override any previously set options.  It only needs to be the
// fields to override it will end up overlaying in the following order:
// 1. hardcoded defaults 2. observedConfig 3. unsupportedConfigOverrides`,
// 										Type:                   "object",
// 										Nullable:               true,
// 										XPreserveUnknownFields: ptrbool(true),
// 									},
// 								},
// 							},
// 							"status": apiextensionsv1.JSONSchemaProps{
// 								Description: `status holds observed values from the cluster. They may not
// be overridden.`,
// 								Type: "object",
// 								Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 									"conditions": apiextensionsv1.JSONSchemaProps{
// 										Description: "conditions is a list of conditions and their status",
// 										Type:        "array",
// 										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
// 											Schema: &apiextensionsv1.JSONSchemaProps{
// 												Description: "OperatorCondition is just the standard condition fields.",
// 												Type:        "object",
// 												Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 													"lastTransitionTime": apiextensionsv1.JSONSchemaProps{
// 														Type:   "string",
// 														Format: "date-time",
// 													},
// 													"message": apiextensionsv1.JSONSchemaProps{
// 														Type: "string",
// 													},
// 													"reason": apiextensionsv1.JSONSchemaProps{
// 														Type: "string",
// 													},
// 													"status": apiextensionsv1.JSONSchemaProps{
// 														Type: "string",
// 													},
// 													"type": apiextensionsv1.JSONSchemaProps{
// 														Type: "string",
// 													},
// 												},
// 											},
// 										},
// 									},
// 									"generations": apiextensionsv1.JSONSchemaProps{
// 										Description: `generations are used to determine when an item needs to
// be reconciled or has changed in a way that needs a reaction.`,
// 										Type: "array",
// 										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
// 											Schema: &apiextensionsv1.JSONSchemaProps{
// 												Description: `GenerationStatus keeps track of the generation for a
// given resource so that decisions about forced updates can be made.`,
// 												Type: "object",
// 												Properties: map[string]apiextensionsv1.JSONSchemaProps{
// 													"group": apiextensionsv1.JSONSchemaProps{
// 														Description: "group is the group of the thing you're tracking",
// 														Type:        "string",
// 													},
// 													"hash": apiextensionsv1.JSONSchemaProps{
// 														Description: `hash is an optional field set for resources without
// generation that are content sensitive like secrets and configmaps`,
// 														Type: "string",
// 													},
// 													"name": apiextensionsv1.JSONSchemaProps{
// 														Description: "name is the name of the thing you're tracking",
// 														Type:        "string",
// 													},
// 													"namespace": apiextensionsv1.JSONSchemaProps{
// 														Description: "namespace is where the thing you're tracking is",
// 														Type:        "string",
// 													},
// 													"resource": apiextensionsv1.JSONSchemaProps{
// 														Description: `resource is the resource type of the thing you're
// tracking`,
// 														Type: "string",
// 													},
// 												},
// 											},
// 										},
// 									},
// 									"observedGeneration": apiextensionsv1.JSONSchemaProps{
// 										Description: `observedGeneration is the last generation change you've
// dealt with`,
// 										Type:   "integer",
// 										Format: "int64",
// 									},
// 									"readyReplicas": apiextensionsv1.JSONSchemaProps{
// 										Description: `readyReplicas indicates how many replicas are ready and
// at the desired state`,
// 										Type:   "integer",
// 										Format: "int32",
// 									},
// 									"version": apiextensionsv1.JSONSchemaProps{
// 										Description: "version is the level this availability applies to",
// 										Type:        "string",
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// var KubevirtCsiCR = &unstructured.Unstructured{
// 	Object: map[string]interface{}{
// 		"apiVersion": "operator.openshift.io/v1",
// 		"kind":       "ClusterCSIDriver",
// 		"metadata": map[string]interface{}{
// 			"name": "csi.kubevirt.io",
// 		},
// 		"spec": map[string]interface{}{
// 			"driverConfig": map[string]interface{}{
// 				"driverName": "csi.kubevirt.io",
// 			},
// 			"logLevel":         "Debug",
// 			"managementState":  "Managed",
// 			"operatorLogLevel": "Debug",
// 		},
// 	},
// }

// var DriverConfig = &corev1.ConfigMap{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "ConfigMap",
// 		APIVersion: "v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "driver-config",
// 		Namespace: "openshift-cluster-csi-drivers",
// 	},
// 	Data: map[string]string{
// 		"infraClusterNamespace": "clusters-isaac-infra",
// 	},
// }

// var KubevirtCsiDaemonset = &appsv1.DaemonSet{
// 	TypeMeta: metav1.TypeMeta{
// 		Kind:       "DaemonSet",
// 		APIVersion: "apps/v1",
// 	},
// 	ObjectMeta: metav1.ObjectMeta{
// 		Name:      "kubevirt-csi-node",
// 		Namespace: "openshift-cluster-csi-drivers",
// 	},
// 	Spec: appsv1.DaemonSetSpec{
// 		Selector: &metav1.LabelSelector{
// 			MatchLabels: map[string]string{
// 				"app": "kubevirt-csi-driver",
// 			},
// 		},
// 		Template: corev1.PodTemplateSpec{
// 			ObjectMeta: metav1.ObjectMeta{
// 				Labels: map[string]string{
// 					"app": "kubevirt-csi-driver",
// 				},
// 			},
// 			Spec: corev1.PodSpec{
// 				Volumes: []corev1.Volume{
// 					corev1.Volume{
// 						Name: "infracluster",
// 						VolumeSource: corev1.VolumeSource{
// 							Secret: &corev1.SecretVolumeSource{
// 								SecretName: "infra-kubeconfig",
// 							},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "kubelet-dir",
// 						VolumeSource: corev1.VolumeSource{
// 							HostPath: &corev1.HostPathVolumeSource{
// 								Path: "/var/lib/kubelet",
// 								Type: ptrHostPathType("Directory"),
// 							},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "plugin-dir",
// 						VolumeSource: corev1.VolumeSource{
// 							HostPath: &corev1.HostPathVolumeSource{
// 								Path: "/var/lib/kubelet/plugins/csi.kubevirt.io/",
// 								Type: ptrHostPathType("DirectoryOrCreate"),
// 							},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "registration-dir",
// 						VolumeSource: corev1.VolumeSource{
// 							HostPath: &corev1.HostPathVolumeSource{
// 								Path: "/var/lib/kubelet/plugins_registry/",
// 								Type: ptrHostPathType("Directory"),
// 							},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "device-dir",
// 						VolumeSource: corev1.VolumeSource{
// 							HostPath: &corev1.HostPathVolumeSource{
// 								Path: "/dev",
// 								Type: ptrHostPathType("Directory"),
// 							},
// 						},
// 					},
// 					corev1.Volume{
// 						Name: "udev",
// 						VolumeSource: corev1.VolumeSource{
// 							HostPath: &corev1.HostPathVolumeSource{
// 								Path: "/run/udev",
// 							},
// 						},
// 					},
// 				},
// 				Containers: []corev1.Container{
// 					corev1.Container{
// 						Name:  "csi-driver",
// 						Image: "quay.io/isaacdorfman/kubevirt-csi-driver:latest",
// 						Args: []string{
// 							"--endpoint=unix:/csi/csi.sock",
// 							"--namespace=kubevirt-csi-driver",
// 							"--node-name=$(KUBE_NODE_NAME)",
// 							"--infra-cluster-namespace=$(INFRACLUSTER_NAMESPACE)",
// 							"--infra-cluster-kubeconfig=/var/run/secrets/infracluster/kubeconfig",
// 							"--run-node-service",
// 							"--v=5",
// 						},
// 						Ports: []corev1.ContainerPort{
// 							corev1.ContainerPort{
// 								Name:          "healthz",
// 								HostPort:      0,
// 								ContainerPort: 10300,
// 								Protocol:      corev1.Protocol("TCP"),
// 							},
// 						},
// 						Env: []corev1.EnvVar{
// 							corev1.EnvVar{
// 								Name: "KUBE_NODE_NAME",
// 								ValueFrom: &corev1.EnvVarSource{
// 									FieldRef: &corev1.ObjectFieldSelector{
// 										FieldPath: "spec.nodeName",
// 									},
// 								},
// 							},
// 							corev1.EnvVar{
// 								Name: "INFRACLUSTER_NAMESPACE",
// 								ValueFrom: &corev1.EnvVarSource{
// 									ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
// 										LocalObjectReference: corev1.LocalObjectReference{
// 											Name: "driver-config",
// 										},
// 										Key: "infraClusterNamespace",
// 									},
// 								},
// 							},
// 							corev1.EnvVar{
// 								Name: "INFRACLUSTER_LABELS",
// 								ValueFrom: &corev1.EnvVarSource{
// 									ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
// 										LocalObjectReference: corev1.LocalObjectReference{
// 											Name: "driver-config",
// 										},
// 										Key: "infraClusterLabels",
// 									},
// 								},
// 							},
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(10, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(52428800, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "infracluster",
// 								MountPath: "/var/run/secrets/infracluster",
// 							},
// 							corev1.VolumeMount{
// 								Name:             "kubelet-dir",
// 								MountPath:        "/var/lib/kubelet",
// 								MountPropagation: ptrMountPropagationMode("Bidirectional"),
// 							},
// 							corev1.VolumeMount{
// 								Name:      "plugin-dir",
// 								MountPath: "/csi",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "device-dir",
// 								MountPath: "/dev",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "udev",
// 								MountPath: "/run/udev",
// 							},
// 						},
// 						LivenessProbe: &corev1.Probe{
// 							ProbeHandler: corev1.ProbeHandler{
// 								HTTPGet: &corev1.HTTPGetAction{
// 									Path: "/healthz",
// 									Port: intstr.IntOrString{
// 										Type:   intstr.Type(1),
// 										IntVal: 0,
// 										StrVal: "healthz",
// 									},
// 								},
// 							},
// 							InitialDelaySeconds: 10,
// 							TimeoutSeconds:      3,
// 							PeriodSeconds:       10,
// 							SuccessThreshold:    0,
// 							FailureThreshold:    5,
// 						},
// 						ImagePullPolicy: corev1.PullPolicy("Always"),
// 						SecurityContext: &corev1.SecurityContext{
// 							Privileged:               ptrbool(true),
// 							AllowPrivilegeEscalation: ptrbool(true),
// 						},
// 					},
// 					corev1.Container{
// 						Name:  "csi-node-driver-registrar",
// 						Image: "quay.io/openshift/origin-csi-node-driver-registrar:latest",
// 						Args: []string{
// 							"--csi-address=$(ADDRESS)",
// 							"--kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)",
// 							"--v=5",
// 						},
// 						Env: []corev1.EnvVar{
// 							corev1.EnvVar{
// 								Name:  "ADDRESS",
// 								Value: "/csi/csi.sock",
// 							},
// 							corev1.EnvVar{
// 								Name:  "DRIVER_REG_SOCK_PATH",
// 								Value: "/var/lib/kubelet/plugins/csi.kubevirt.io/csi.sock",
// 							},
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(5, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(20971520, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "plugin-dir",
// 								MountPath: "/csi",
// 							},
// 							corev1.VolumeMount{
// 								Name:      "registration-dir",
// 								MountPath: "/registration",
// 							},
// 						},
// 						Lifecycle: &corev1.Lifecycle{
// 							PreStop: &corev1.LifecycleHandler{
// 								Exec: &corev1.ExecAction{
// 									Command: []string{
// 										"/bin/sh",
// 										"-c",
// 										"rm -rf /registration/csi.kubevirt.io-reg.sock /csi/csi.sock",
// 									},
// 								},
// 							},
// 						},
// 						SecurityContext: &corev1.SecurityContext{
// 							Privileged: ptrbool(true),
// 						},
// 					},
// 					corev1.Container{
// 						Name:  "csi-liveness-probe",
// 						Image: "quay.io/openshift/origin-csi-livenessprobe:latest",
// 						Args: []string{
// 							"--csi-address=/csi/csi.sock",
// 							"--probe-timeout=3s",
// 							"--health-port=10300",
// 						},
// 						Resources: corev1.ResourceRequirements{
// 							Requests: corev1.ResourceList{
// 								"cpu":    *resource.NewQuantity(5, resource.DecimalSI),
// 								"memory": *resource.NewQuantity(20971520, resource.BinarySI),
// 							},
// 						},
// 						VolumeMounts: []corev1.VolumeMount{
// 							corev1.VolumeMount{
// 								Name:      "plugin-dir",
// 								MountPath: "/csi",
// 							},
// 						},
// 					},
// 				},
// 				DeprecatedServiceAccount: "kubevirt-csi-node-sa",
// 				HostNetwork:              true,
// 				Tolerations: []corev1.Toleration{
// 					corev1.Toleration{
// 						Operator: corev1.TolerationOperator("Exists"),
// 					},
// 				},
// 				PriorityClassName: "system-node-critical",
// 			},
// 		},
// 		UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
// 			Type: appsv1.DaemonSetUpdateStrategyType("RollingUpdate"),
// 		},
// 		MinReadySeconds: 0,
// 	},
// }

//go:embed files/*
var resources embed.FS

func getContents(file string) []byte {
	f, err := resources.Open("files/" + file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func GenerateTenantNamespace(namespace string) crclient.Object {
	mustacheRenderTemplate := map[string]string{"NAMESPACE": namespace}

	namespaceBytes := getContents("namespace.yaml")
	renderedNamespace := mustache.Render(string(namespaceBytes), mustacheRenderTemplate)
	namespaceObject := &corev1.Namespace{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedNamespace)), 100).Decode(&namespaceObject); err != nil {
		panic(err)
	}
	return namespaceObject
}

// schema.GroupVersionKind
//unstructured.Unstructured

func GenerateInfraServiceAccountResources(namespace string) (*corev1.ServiceAccount, *rbacv1.Role, *rbacv1.RoleBinding) {
	mustacheRenderTemplate := map[string]string{"INFRA_NAMESPACE": namespace}

	infraServiceAccountBytes := getContents("infra_serviceaccount.yaml")
	renderedInfraServiceAccount := mustache.Render(string(infraServiceAccountBytes), mustacheRenderTemplate)
	serviceaccount := &corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraServiceAccount)), 500).Decode(&serviceaccount); err != nil {
		panic(err)
	}

	infraRoleBytes := getContents("infra_role.yaml")
	renderedInfraRole := mustache.Render(string(infraRoleBytes), mustacheRenderTemplate)
	role := &rbacv1.Role{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRole)), 500).Decode(&role); err != nil {
		panic(err)
	}

	infraRoleBindingBytes := getContents("infra_rolebinding.yaml")
	renderedInfraRoleBinding := mustache.Render(string(infraRoleBindingBytes), mustacheRenderTemplate)
	rolebinding := &rbacv1.RoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRoleBinding)), 500).Decode(&rolebinding); err != nil {
		panic(err)
	}

	return serviceaccount, role, rolebinding
}

func GenerateTenantNodeServiceAccountResources(namespace string) []crclient.Object {
	mustacheRenderTemplate := map[string]string{"NAMESPACE": namespace}

	infraServiceAccountBytes := getContents("tenant_node_serviceaccount.yaml")
	renderedInfraServiceAccount := mustache.Render(string(infraServiceAccountBytes), mustacheRenderTemplate)
	serviceaccount := &corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraServiceAccount)), 500).Decode(&serviceaccount); err != nil {
		panic(err)
	}

	infraRoleBytes := getContents("tenant_node_clusterrole.yaml")
	renderedInfraRole := mustache.Render(string(infraRoleBytes), mustacheRenderTemplate)
	role := &rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRole)), 500).Decode(&role); err != nil {
		panic(err)
	}

	infraRoleBindingBytes := getContents("tenant_node_clusterrolebinding.yaml")
	renderedInfraRoleBinding := mustache.Render(string(infraRoleBindingBytes), mustacheRenderTemplate)
	rolebinding := &rbacv1.ClusterRoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRoleBinding)), 500).Decode(&rolebinding); err != nil {
		panic(err)
	}

	return []crclient.Object{serviceaccount, role, rolebinding}
}

func GenerateTenantControllerServiceAccountResources(namespace string) []crclient.Object {
	mustacheRenderTemplate := map[string]string{"NAMESPACE": namespace}

	infraServiceAccountBytes := getContents("tenant_controller_serviceaccount.yaml")
	renderedInfraServiceAccount := mustache.Render(string(infraServiceAccountBytes), mustacheRenderTemplate)
	serviceaccount := &corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraServiceAccount)), 500).Decode(&serviceaccount); err != nil {
		panic(err)
	}

	infraRoleBytes := getContents("tenant_controller_clusterrole.yaml")
	renderedInfraRole := mustache.Render(string(infraRoleBytes), mustacheRenderTemplate)
	role := &rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRole)), 500).Decode(&role); err != nil {
		panic(err)
	}

	infraRoleBindingBytes := getContents("tenant_controller_clusterrolebinding.yaml")
	renderedInfraRoleBinding := mustache.Render(string(infraRoleBindingBytes), mustacheRenderTemplate)
	rolebinding := &rbacv1.ClusterRoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedInfraRoleBinding)), 500).Decode(&rolebinding); err != nil {
		panic(err)
	}

	return []crclient.Object{serviceaccount, role, rolebinding}
}

func GenerateKubeconfig(serverUrl string, namespace string, certificateAuthorityData string, userToken string) string {
	mustacheRenderTemplate := map[string]string{
		"SERVER_URL":                 serverUrl,
		"NAMESPACE":                  namespace,
		"CERTIFICATE_AUTHORITY_DATA": certificateAuthorityData,
		"USER_TOKEN":                 userToken,
	}
	kubeconfigBytes := getContents("kubeconfig_template.yaml")

	renderedKubeconfig := mustache.Render(string(kubeconfigBytes), mustacheRenderTemplate)

	return renderedKubeconfig
}

func GenerateKubeconfigSecret(secretName string, secretNamespace string, kubeconfig string) crclient.Object {
	mustacheRenderTemplate := map[string]string{
		"SECRET_NAME":      secretName,
		"SECRET_NAMESPACE": secretName,
		"KUBECONFIG":       kubeconfig,
	}
	secretBytes := getContents("kubeconfig_secret_template.yaml")
	renderedSecret := mustache.Render(string(secretBytes), mustacheRenderTemplate)

	secret := &corev1.Secret{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedSecret)), 500).Decode(&secret); err != nil {
		panic(err)
	}
	return secret
}

func GenerateTenantConfigmap(tenantNamespace string, infraNamespace string) crclient.Object {
	mustacheRenderTemplate := map[string]string{
		"TENANT_NAMESPACE": tenantNamespace,
		"INFRA_NAMESPACE":  infraNamespace,
	}
	configmapBytes := getContents("tenant_configmap.yaml")
	renderedConfigmap := mustache.Render(string(configmapBytes), mustacheRenderTemplate)
	configmap := &corev1.ConfigMap{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedConfigmap)), 500).Decode(&configmap); err != nil {
		panic(err)
	}
	return configmap
}

func GenerateInfraConfigmap(namespace string) crclient.Object {
	mustacheRenderTemplate := map[string]string{
		"NAMESPACE": namespace,
	}
	configmapBytes := getContents("infra_configmap.yaml")
	renderedConfigmap := mustache.Render(string(configmapBytes), mustacheRenderTemplate)
	configmap := &corev1.ConfigMap{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedConfigmap)), 500).Decode(&configmap); err != nil {
		panic(err)
	}
	return configmap
}

func GenerateDaemonset(namespace string) crclient.Object {
	mustacheRenderTemplate := map[string]string{
		"NAMESPACE": namespace,
	}
	daemonsetBytes := getContents("daemonset.yaml")
	renderedDaemonset := mustache.Render(string(daemonsetBytes), mustacheRenderTemplate)
	daemonset := &appsv1.DaemonSet{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedDaemonset)), 500).Decode(&daemonset); err != nil {
		panic(err)
	}
	return daemonset
}

func GenerateController(namespace string) crclient.Object {
	mustacheRenderTemplate := map[string]string{
		"NAMESPACE": namespace,
	}
	controllerBytes := getContents("controller.yaml")
	renderedController := mustache.Render(string(controllerBytes), mustacheRenderTemplate)
	controller := &appsv1.Deployment{}
	if err := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(renderedController)), 500).Decode(&controller); err != nil {
		panic(err)
	}
	return controller
}
