package client

import (
	"fmt"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	kubeconfig string
	restConfig *rest.Config

	clientSet    *kubernetes.Clientset
	crdClientSet *clientset.Clientset
	dynamicClient *dynamic.DynamicClient
}

// K8sClient returns a new Client to interact with Kubernetes, based on the provided kubeconfig.
// If the kubeconfig is an empty string, it uses in-cluster configuration.
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

	crdClientSet, err := clientset.NewForConfig(cl.restConfig)
	if err != nil {
		return nil, err
	}

	dynClient, err := dynamic.NewForConfig(cl.restConfig)
	if err != nil {
		return nil, err
	}

	cl.clientSet = cs
	cl.crdClientSet = crdClientSet
	cl.dynamicClient = dynClient

	return cl, nil
}
