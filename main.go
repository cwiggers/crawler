package main

import (
	"flag"
	"os"

	"github.com/cwiggers/crawler/spider"
	"github.com/cwiggers/crawler/tools"

	log "github.com/alecthomas/log4go"
)

func main() {
	logConfigFile := flag.String("l", "runtime/log4go.xml", "Log config file")
	configFile := flag.String("c", "runtime/conf.json", "Config file")

	flag.Parse()

	log.LoadConfiguration(*logConfigFile)

	if err := tools.LoadConf(*configFile); err != nil {
		log.Error("failed to load configure, Err:[%s]", err)
		os.Exit(1)
	}

	spider.NewSpider().Run(tools.Conf.Addr)
}
