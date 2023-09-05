package k8s

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	kubeconfig string
	restConfig *rest.Config

	ClientSet *kubernetes.Clientset
}

// K8sClient returns a new Client to interact with Kubernetes, based on the provided kubeconfig.
// If the kubeconfig is an empty string, it uses in-cluter configuration.
func K8sClient(kubeconfig string) (*Client, error) {
	cl := &Client{kubeconfig: kubeconfig}

	var err error

	if cl.kubeconfig == "" {
		fmt.Println("using in-cluster configuration")
		cl.restConfig, err = rest.InClusterConfig()
	} else {
		fmt.Printf("using configuration from '%s'\n", kubeconfig)
		cl.restConfig, err = clientcmd.BuildConfigFromFlags("", cl.kubeconfig)
	}

	if err != nil {
		fmt.Printf("failed to generate config, err: %v", err)
		return nil, err
	}

	cs, err := kubernetes.NewForConfig(cl.restConfig)
	if err != nil {
		fmt.Printf("failed to generate client, err: %v", err)
		return nil, err
	}

	cl.ClientSet = cs

	return cl, nil
}
