package api

import (
	"context"
	"encoding/json"
	"github.com/buraksekili/tykstation/k8s"
	"github.com/gorilla/mux"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type Handler struct {
}

func MakeHTTPHandler(ctx context.Context, client *k8s.Client) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/deploys").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deploys, err := client.Deploys(ctx, "tyk", v1.ListOptions{})
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(codeFrom(err))
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(deploys)
	})

	return r
}

func codeFrom(err error) int {
	switch err {
	//case users.ErrUnauthorized:
	//	return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
