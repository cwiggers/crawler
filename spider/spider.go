package spider

import (
	"net/http"

	log "github.com/alecthomas/log4go"
	"github.com/gorilla/mux"
)

type Spider struct {
}

func NewSpider() *Spider {
	return &Spider{}
}

func CrawerHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Spider) Run(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/{crawler}", CrawerHandler).Methods("POST")
	r.HandleFunc("/{crawler}/status", StatusHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	log.Error(srv.ListenAndServe())
}
