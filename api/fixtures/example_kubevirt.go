package fixtures

import (
	"fmt"

	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
)

func ExampleKubeVirtTemplate(o *ExampleKubevirtOptions) *hyperv1.KubevirtNodePoolPlatform {
	var storageClassName *string
	memory := apiresource.MustParse(o.Memory)
	volumeSize := apiresource.MustParse(fmt.Sprintf("%vGi", o.RootVolumeSize))

	if o.RootVolumeStorageClass != "" {
		storageClassName = &o.RootVolumeStorageClass
	}

	exampleTemplate := &hyperv1.KubevirtNodePoolPlatform{
		RootVolume: &hyperv1.KubevirtRootVolume{
			KubevirtVolumeTypes: hyperv1.KubevirtVolumeTypes{
				Persistent: &hyperv1.KubevirtPersistentVolume{
					Size:         &volumeSize,
					StorageClass: storageClassName,
				},
			},
		},
		Compute: &hyperv1.KubevirtCompute{
			Memory: &memory,
			Cores:  &o.Cores,
		},
	}

	if o.Image != "" {
		if exampleTemplate.RootVolume == nil {
			exampleTemplate.RootVolume = &hyperv1.KubevirtRootVolume{}
		}
		exampleTemplate.RootVolume.Image = &hyperv1.KubevirtDiskImage{
			ContainerDiskImage: &o.Image,
		}
	}

	return exampleTemplate
}
