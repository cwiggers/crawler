package spider

import (
	"fmt"
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
	vars := mux.Vars(r)
	crawler := vars["crawler"]
	fmt.Fprintf(w, crawler)
	return
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crawler := vars["crawler"]
	fmt.Fprintf(w, "%s is OK", crawler)
	return
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Spider) Run(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/{crawler}", CrawerHandler).Methods("POST")
	r.HandleFunc("/{crawler}/status", StatusHandler).Methods("GET")
	r.HandleFunc("/", HomeHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	log.Error(srv.ListenAndServe())
}
