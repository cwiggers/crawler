package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	log "code.google.com/p/log4go"
	"github.com/cwiggers/crawler/handler"
	"github.com/cwiggers/crawler/misc"
	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("fuck crash: (%v)\n", r)
			debug.PrintStack()
		}
	}()

	logConfigFile := flag.String("l", "./runtime/log4go.xml", "Log config file")
	configFile := flag.String("c", "./runtime/conf.json", "Config file")

	flag.Parse()

	log.LoadConfiguration(*logConfigFile)

	if err := misc.LoadConf(*configFile); err != nil {
		fmt.Printf("failed to load conf [%s]: (%s)", *configFile, err)
		os.Exit(1)
	}

	misc.InitBackend()
	misc.InitQueue()

	r := mux.NewRouter()
	r.HandleFunc("/", handler.AllHandler).Methods("GET")
	r.HandleFunc("/{crawler}", handler.CrawlerHandler).Methods("POST")
	r.HandleFunc("/{crawler}/status", handler.StatusHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         misc.Conf.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Error(srv.ListenAndServe())

}
