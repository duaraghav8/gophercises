package apiserver

import "net/http"

func (server *APIServer) HandlerReadSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
