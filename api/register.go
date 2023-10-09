package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buraksekili/tykstation/k8s/client"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

// Custom Resource Definition Handlers
func registerGetCRHandler(ctx context.Context, c *client.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		group, ok := vars["group"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		version, ok := vars["version"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		resource, ok := vars["resource"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		ns, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		name, ok := vars["name"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
			return
		}

		crs, err := c.GetCR(ctx, ns, name, group, version, resource)
		if err != nil {
			errorHandler(w, err)
		}

		json.NewEncoder(w).Encode(crs)
	}
}

func registerCRDsHandler(ctx context.Context, c *client.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		crds, err := c.GetCRDs(ctx)
		if err != nil {
			errorHandler(w, err)
		}

		json.NewEncoder(w).Encode(crds)
	}
}

// Apps V1 Handlers
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// You can customize this function to check the origin of the request
		return true
	},
}

// TODO: handle connection terminations.
func registerWatchAppsV1Handlers(ctx context.Context, cl *client.Client, appsV1Type string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("registering watch")

		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		fmt.Println("namespace", namespace)

		watchInterface, err := cl.WatchAppsV1(ctx, namespace, appsV1Type, v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			errorHandler(w, err)
			fmt.Println("failed to upgrade", err)
			return
		}
		defer conn.Close()

		for update := range watchInterface.ResultChan() {
			fmt.Println("writing update")
			err := conn.WriteJSON(update)
			if err != nil {
				fmt.Println("failed to write update")
				errorHandler(w, err)
				return
			}
			fmt.Println("!WORKED")

			fmt.Printf(
				"Watch Event: %s %s\n",
				update.Type, update.Object.GetObjectKind().GroupVersionKind().Kind,
			)
		}
		fmt.Println("finished")
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

// Core V1 handlers.
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

// TODO: handle connection terminations.
func registerWatchCoreV1Handlers(ctx context.Context, cl *client.Client, appsV1Type string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("registering watch")

		vars := mux.Vars(r)
		namespace, ok := vars["namespace"]
		if !ok {
			errorHandler(w, errors.New("invalid request path"))
		}

		fmt.Println("namespace", namespace)

		watchInterface, err := cl.WatchCoreV1(ctx, namespace, appsV1Type, v1.ListOptions{})
		if err != nil {
			errorHandler(w, err)
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			errorHandler(w, err)
			fmt.Println("failed to upgrade", err)
			return
		}
		defer conn.Close()

		for update := range watchInterface.ResultChan() {
			fmt.Println("writing update")
			err := conn.WriteJSON(update)
			if err != nil {
				fmt.Println("failed to write update")
				errorHandler(w, err)
				return
			}
			fmt.Println("!WORKED")

			fmt.Printf(
				"Watch Event: %s %s\n",
				update.Type, update.Object.GetObjectKind().GroupVersionKind().Kind,
			)
		}
		fmt.Println("finished")
	}

}
