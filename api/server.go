package api

import "net/http"

func CreateNewServer(router *http.ServeMux) http.Server {
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return server
}
