package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/buraksekili/tykstation/k8s"
	"github.com/gorilla/mux"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type Handler struct {
}

var coreV1Types = []string{
	"pods", "services", "secrets", "configmaps", "endpoints",
	"events", "namespaces", "nodes", "pvcs", "pvs", "replicacontrollers", "serviceaccounts",
}

func MakeHTTPHandler(ctx context.Context, client *k8s.Client) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/deploys").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deploys, err := client.Deploys(ctx, "tyk", v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(deploys)
	})

	for _, coreV1Type := range coreV1Types {
		r.Methods("GET").
			Path(fmt.Sprintf("/%s", coreV1Type)).
			HandlerFunc(registerListCoreV1Handlers(ctx, client, coreV1Type))
	}

	return r
}

func registerListCoreV1Handlers(ctx context.Context, cl *k8s.Client, resource string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pods, err := cl.ListCoreV1(ctx, "tyk", resource, v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(pods)
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
