package main

import (
	"context"
	"fmt"
	"github.com/buraksekili/tykstation/api"
	"github.com/buraksekili/tykstation/k8s/client"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()

	cl, err := client.K8sClient(kubeconfig)
	if err != nil {
		panic(err)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	listenAddrs := 8080

	go func() {
		fmt.Printf("Tyk Station started on %v\n", listenAddrs)
		errs <- http.ListenAndServe(fmt.Sprintf("localhost:%v", listenAddrs), api.MakeHTTPHandler(context.Background(), cl))
	}()

	fmt.Printf("ERROR: %v", <-errs)

}

//func getClientset() (*kubernetes.Clientset, error) {
//	config, err := rest.InClusterConfig()
//	if err != nil {
//		kubeconfig :=
//			clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
//		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
//		if err != nil {
//			return nil, err
//		}
//	}
//	return kubernetes.NewForConfig(config)
//}
