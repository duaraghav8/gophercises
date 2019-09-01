package apiserver

import "net/http"

func (server *APIServer) HandlerWriteSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
