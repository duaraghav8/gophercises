package apiserver

import "net/http"

func registerAPIRoutes(server APIServer) {
	server.Router.
		HandleFunc("/keys", server.HandlerListSecrets).
		Methods(http.MethodGet)
	server.Router.
		HandleFunc("/secret/{key}", server.HandlerReadSecret).
		Methods(http.MethodGet)
	server.Router.
		HandleFunc("/secret/{key}", server.HandlerWriteSecret).
		Methods(http.MethodPost)
}
