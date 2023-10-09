package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
