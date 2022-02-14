package common

import (
	"context"
	"fmt"

	"github.com/openshift/hypershift/api/fixtures"
	"github.com/openshift/hypershift/cmd/log"
	"github.com/openshift/hypershift/cmd/util"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServicesPublishOptions struct {
	UseNodePortPublishStrategy bool
}

func NewServicesPublishOptions(cmd *cobra.Command) *ServicesPublishOptions {

	result := &ServicesPublishOptions{
		UseNodePortPublishStrategy: false,
	}

	cmd.Flags().BoolVar(&result.UseNodePortPublishStrategy, "use-nodeport-publish-strategy", result.UseNodePortPublishStrategy, "Use NodePort publish strategy type for APIServer service expose")

	return result
}

func (o *ServicesPublishOptions) ExampleServicesPublishOpts(ctx context.Context, render bool) (*fixtures.ExampleServicesPublishOptions, error) {
	var err error
	result := &fixtures.ExampleServicesPublishOptions{
		UseNodePortPublishStrategy: o.UseNodePortPublishStrategy,
	}
	if result.UseNodePortPublishStrategy && !render {
		if result.APIServerAddress, err = getAPIServerAddressByNode(ctx); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func getAPIServerAddressByNode(ctx context.Context) (string, error) {
	// Fetch a single node and determine possible DNS or IP entries to use
	// for external node-port communication.
	// Possible values are considered with the following priority based on the address type:
	// - NodeExternalDNS
	// - NodeExternalIP
	// - NodeInternalIP
	apiServerAddress := ""
	kubeClient := kubernetes.NewForConfigOrDie(util.GetConfigOrDie())
	nodes, err := kubeClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{Limit: 1})
	if err != nil {
		return "", fmt.Errorf("unable to fetch node objects: %w", err)
	}
	if len(nodes.Items) < 1 {
		return "", fmt.Errorf("no node objects found: %w", err)
	}
	addresses := map[corev1.NodeAddressType]string{}
	for _, address := range nodes.Items[0].Status.Addresses {
		addresses[address.Type] = address.Address
	}
	for _, addrType := range []corev1.NodeAddressType{corev1.NodeExternalDNS, corev1.NodeExternalIP, corev1.NodeInternalIP} {
		if address, exists := addresses[addrType]; exists {
			apiServerAddress = address
			break
		}
	}
	if apiServerAddress == "" {
		return "", fmt.Errorf("node %q does not expose any IP addresses, this should not be possible", nodes.Items[0].Name)
	}
	log.Log.Info(fmt.Sprintf("detected %q from node %q as external-api-server-address", apiServerAddress, nodes.Items[0].Name))
	return apiServerAddress, nil
}
