package hostedcluster

import (
	"context"
	"fmt"

	hyperv1 "github.com/openshift/hypershift/api/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

type hostedClusterDefaulter struct{}

func (defaulter *hostedClusterDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	_, ok := obj.(*hyperv1.HostedCluster)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a HostedCluster but got a %T", obj))
	}

	fmt.Printf("DEBUG!!! got hc mutate defaulter")
	return nil
}

type nodePoolDefaulter struct{}

func (defaulter *nodePoolDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	_, ok := obj.(*hyperv1.NodePool)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected a NodePool but got a %T", obj))
	}

	fmt.Printf("DEBUG!!! got np mutate defaulter")
	return nil
}

// SetupWebhookWithManager sets up HostedCluster webhooks.
func SetupWebhookWithManager(mgr ctrl.Manager) error {

	err := ctrl.NewWebhookManagedBy(mgr).
		For(&hyperv1.HostedCluster{}).
		WithDefaulter(&hostedClusterDefaulter{}).
		Complete()
	if err != nil {
		return fmt.Errorf("unable to register hostedcluster webhook: %w", err)
	}
	err = ctrl.NewWebhookManagedBy(mgr).
		For(&hyperv1.NodePool{}).
		WithDefaulter(&nodePoolDefaulter{}).
		Complete()
	if err != nil {
		return fmt.Errorf("unable to register nodepool webhook: %w", err)
	}

	err = ctrl.NewWebhookManagedBy(mgr).
		For(&hyperv1.HostedControlPlane{}).
		Complete()
	if err != nil {
		return fmt.Errorf("unable to register hostedcontrolplane webhook: %w", err)
	}

	return nil

}
