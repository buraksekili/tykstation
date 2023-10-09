package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/buraksekili/tykstation/k8s/client"
	"github.com/gorilla/mux"
	"helm.sh/helm/v3/pkg/time"
	"io"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

type Handler struct {
}

// MakeHTTPHandler registers endpoints for k8s GVKs.
func MakeHTTPHandler(ctx context.Context, client *client.Client) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").
		Path("/logs/{namespace}/{name}").
		HandlerFunc(logsHandler(ctx, client))

	// Core V1 types
	for _, coreV1Type := range coreV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/corev1/{namespace}/%s", coreV1Type)).
			HandlerFunc(registerListCoreV1Handlers(ctx, client, coreV1Type))
	}
	for _, coreV1Type := range coreV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/corev1/{namespace}/%s/{name}", coreV1Type)).
			HandlerFunc(registerGetCoreV1Handlers(ctx, client, coreV1Type))
	}

	// Apps V1 types
	for _, appsV1Type := range appsV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/appsv1/{namespace}/%s", appsV1Type)).
			HandlerFunc(registerListAppsV1Handlers(ctx, client, appsV1Type))
	}

	for _, appsV1Type := range appsV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/appsv1/{namespace}/%s/{name}", appsV1Type)).
			HandlerFunc(registerGetAppsV1Handlers(ctx, client, appsV1Type))
	}

	return r
}

func logsHandler(ctx context.Context, client *client.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		defer func() {
			fmt.Println("finished totally\n")
		}()

		namespace, namespaceExists := vars["namespace"]
		if !namespaceExists {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		name, nameExists := vars["name"]
		if !nameExists {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		now := time.Now()

		yesterday := now.AddDate(0, 0, -1)
		fmt.Println(yesterday.String())
		u := now.Unix()

		logStreamReq := client.ClientSet.CoreV1().Pods(namespace).GetLogs(name, &corev1.PodLogOptions{Follow: true, SinceSeconds: &u})
		rc, _ := logStreamReq.Stream(ctx)
		//scanner := bufio.NewScanner(rc)

		io.Copy(w, rc)

		//for scanner.Scan() {
		//	line := scanner.Text() + "\n"
		//	w.Write([]byte(line))
		//}
	}
}
