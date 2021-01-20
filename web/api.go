package web

import (
	"encoding/json"
	"net/http"
	"time"
	"user-service/global"
	"user-service/users"
)

type api struct {
	mux *http.ServeMux
	server *http.Server
	config *global.Config
	userSvc users.Service
}

func (a *api) handlers() {
	a.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	a.mux.HandleFunc("/getUserSummary", a.GetUserSummary)
}

func (a api) ListenAndServe() error {
	return a.server.ListenAndServe()
}

func New(config *global.Config, userSvc users.Service) *api {
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
		config: config,
		userSvc: userSvc,
	}
	api.handlers()
	return api
}

func (a api) GetUserSummary(w http.ResponseWriter, r *http.Request) {
	req := new(UserSummaryRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No user ids were supplied."))
		return
	}
	result := a.userSvc.GetUserSummary(req.Ids)
	resp, _ := json.Marshal(result)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}