package kubevirtexternalinfra

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sync"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	cr "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hyperv1 "github.com/openshift/hypershift/api/v1beta1"
)

type KubevirtInfraClientMap interface {
	DiscoverKubevirtClusterClient(context.Context, client.Client, string, *hyperv1.KubevirtPlatformCredentials, string, string) (*KubevirtInfraClient, error)
	GetClient(key string) *KubevirtInfraClient
	Delete(string)
}

func NewKubevirtInfraClientMap() KubevirtInfraClientMap {
	return &kubevirtInfraClientMapImp{
		theMap: sync.Map{},
	}
}

type kubevirtInfraClientMapImp struct {
	theMap sync.Map
}

type KubevirtInfraClient struct {
	Client          client.Client
	DiscoveryClient *discovery.DiscoveryClient

	Namespace string
}

func (k *kubevirtInfraClientMapImp) DiscoverKubevirtClusterClient(ctx context.Context, cl client.Client, key string, credentials *hyperv1.KubevirtPlatformCredentials, localInfraNamespace string, secretNS string) (*KubevirtInfraClient, error) {
	if k == nil {
		return nil, nil
	}

	if credentials == nil || credentials.InfraKubeConfigSecret == nil {
		cfg, err := cr.GetConfig()
		if err != nil {
			return nil, err
		}

		discoveryClient, err := discovery.NewDiscoveryClientForConfig(cfg)
		if err != nil {
			return nil, err
		}

		return &KubevirtInfraClient{
			Client:          cl,
			DiscoveryClient: discoveryClient,
			Namespace:       localInfraNamespace,
		}, nil
	}
	loaded, ok := k.theMap.Load(key)
	if ok {
		return loaded.(*KubevirtInfraClient), nil
	}
	targetClient, targetDiscoveryClient, err := generateKubevirtInfraClusterClient(ctx, cl, credentials, secretNS)
	if err != nil {
		return nil, err
	}

	cluster := &KubevirtInfraClient{
		Client:          targetClient,
		DiscoveryClient: targetDiscoveryClient,
		Namespace:       credentials.InfraNamespace,
	}

	k.theMap.LoadOrStore(key, cluster)
	return cluster, nil
}

func (k *kubevirtInfraClientMapImp) GetClient(key string) *KubevirtInfraClient {
	if k == nil {
		return nil
	}
	if cl, ok := k.theMap.Load(key); ok {
		if clnt, ok := cl.(*KubevirtInfraClient); ok {
			return clnt
		}
	}
	return nil
}

func (k *kubevirtInfraClientMapImp) Delete(key string) {
	if k != nil {
		k.theMap.Delete(key)
	}
}

func generateKubevirtInfraClusterClient(ctx context.Context, cpClient client.Client, credentials *hyperv1.KubevirtPlatformCredentials, secretNamespace string) (client.Client, *discovery.DiscoveryClient, error) {
	infraKubeconfigSecret := &corev1.Secret{}

	infraKubeconfigSecretKey := client.ObjectKey{Namespace: secretNamespace, Name: credentials.InfraKubeConfigSecret.Name}
	if err := cpClient.Get(ctx, infraKubeconfigSecretKey, infraKubeconfigSecret); err != nil {
		return nil, nil, fmt.Errorf("failed to fetch infra kubeconfig secret %s/%s: %w", secretNamespace, credentials.InfraKubeConfigSecret.Name, err)
	}

	kubeConfig, ok := infraKubeconfigSecret.Data["kubeconfig"]
	if !ok {
		return nil, nil, errors.New("failed to retrieve infra kubeconfig from secret: 'kubeconfig' key is missing")
	}

	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create K8s-API client config: %w", err)
	}

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create REST config: %w", err)
	}
	var infraClusterClient client.Client

	infraClusterClient, err = client.New(restConfig, client.Options{Scheme: cpClient.Scheme()})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create infra cluster client: %w", err)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create infra cluster discovery client: %w", err)
	}

	return infraClusterClient, discoveryClient, nil
}

func (k *KubevirtInfraClient) ValidateInfraVersioning() error {

	type info struct {
		GitVersion   string `json:"gitVersion"`
		GitCommit    string `json:"gitCommit"`
		GitTreeState string `json:"gitTreeState"`
		BuildDate    string `json:"buildDate"`
		GoVersion    string `json:"goVersion"`
		Compiler     string `json:"compiler"`
		Platform     string `json:"platform"`
	}

	restClient := k.DiscoveryClient.RESTClient()

	var group metav1.APIGroup
	// First, find out which version to query
	uri := "/apis/subresources.kubevirt.io"
	result := restClient.Get().AbsPath(uri).Do(context.Background())
	if data, err := result.Raw(); err != nil {
		connErr, isConnectionErr := err.(*url.Error)

		if isConnectionErr {
			return connErr.Err
		}

		return err
	} else if err = json.Unmarshal(data, &group); err != nil {
		return err
	}

	// Now, query the preferred version
	uri = fmt.Sprintf("/apis/%s/version", group.PreferredVersion.GroupVersion)
	var serverInfo info

	result = restClient.Get().AbsPath(uri).Do(context.Background())
	if data, err := result.Raw(); err != nil {
		connErr, isConnectionErr := err.(*url.Error)

		if isConnectionErr {
			return connErr.Err
		}

		return err
	} else if err = json.Unmarshal(data, &serverInfo); err != nil {
		return err
	}

	fmt.Printf("CNV SERVER INFO: %v\n", serverInfo)

	// K8S VERSION
	k8sVersion, err := k.DiscoveryClient.ServerVersion()
	if err != nil {
		return err
	}

	fmt.Printf("K8S SERVER INFO: %v\n", k8sVersion)

	return nil

}
