package handler

import (
	"fmt"
	"net/http"

	_ "code.google.com/p/log4go"
	"github.com/gorilla/mux"
)

func AllHandler(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func CrawlerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crawler := vars["crawler"]
	fmt.Fprintf(w, "Hello %s!", crawler)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crawler := vars["crawler"]
	fmt.Fprintf(w, "%s status is OK!", crawler)
}
