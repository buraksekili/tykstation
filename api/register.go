package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/buraksekili/tykstation/k8s/client"
	"github.com/gorilla/mux"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func registerGetAppsV1Handlers(ctx context.Context, cl *client.Client, appsV1Type string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		objectName, ok := vars["name"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		appsV1Resource, err := cl.GetAppsV1(ctx, namespace, objectName, appsV1Type, v1.GetOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		json.NewEncoder(w).Encode(appsV1Resource)
	}
}

func registerListAppsV1Handlers(ctx context.Context, cl *client.Client, resource string) func(w http.ResponseWriter, r *http.Request) {
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

func registerGetCoreV1Handlers(ctx context.Context, cl *client.Client, coreV1Type string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		objectName, ok := vars["name"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		appsV1Resource, err := cl.GetCoreV1(ctx, namespace, objectName, coreV1Type, v1.GetOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		json.NewEncoder(w).Encode(appsV1Resource)
	}
}

func registerListCoreV1Handlers(ctx context.Context, cl *client.Client, resource string) func(w http.ResponseWriter, r *http.Request) {
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
