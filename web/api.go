package web

import (
	"net/http"
	"time"
	"user-service/global"
)

type api struct {
	mux *http.ServeMux
	server *http.Server
	config *global.Config
}

func (a *api) handlers() {
	a.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (a api) ListenAndServe() error {
	return a.server.ListenAndServe()
}

func New(config *global.Config) *api {
	timeout := 10 * time.Second
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           mux,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
	}

	api := &api{
		mux:    mux,
		server: server,
	}
	api.handlers()
	return api
}

