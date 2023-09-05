package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buraksekili/tykstation/k8s"
	"github.com/gorilla/mux"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type Handler struct {
}

var (
	coreV1Types = []string{
		"pods", "services", "secrets", "configmaps", "endpoints",
		"events", "namespaces", "nodes", "pvcs", "pvs", "replicacontrollers", "serviceaccounts",
	}

	appsV1Types = []string{"deploys", "daemonsets", "statefulsets", "replicasets"}
)

// MakeHTTPHandler registers endpoints for k8s GVKs.
func MakeHTTPHandler(ctx context.Context, client *k8s.Client) http.Handler {
	r := mux.NewRouter()

	for _, coreV1Type := range coreV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/corev1/{namespace}/%s", coreV1Type)).
			HandlerFunc(registerListCoreV1Handlers(ctx, client, coreV1Type))
	}

	for _, appsV1Type := range appsV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/appsv1/{namespace}/%s", appsV1Type)).
			HandlerFunc(registerAppsV1Handlers(ctx, client, appsV1Type))
	}

	return r
}

func registerAppsV1Handlers(ctx context.Context, cl *k8s.Client, resource string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		resources, err := cl.ListAppsV1(ctx, namespace, resource, v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(resources)
	}
}

func registerListCoreV1Handlers(ctx context.Context, cl *k8s.Client, resource string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		resources, err := cl.ListCoreV1(ctx, namespace, resource, v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(resources)
	}
}

func errorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	errEnc := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
	if errEnc != nil {
		fmt.Printf("ERROR: Failed to encode error message, err: %v", errEnc)
	}
}

func codeFrom(err error) int {
	switch err {
	//case users.ErrUnauthorized:
	//	return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
