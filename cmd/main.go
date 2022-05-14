package main

import (
	"flag"
	log "github.com/hashicorp/go-hclog"
	"github.com/huster-hbl/jaeger-bq/storage"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc"
	"github.com/jaegertracing/jaeger/plugin/storage/grpc/shared"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "", "The absolute path the BigQuery Plugin's configuration file")
	flag.Parse()

	logger := log.New(&log.LoggerOptions{
		Name:       "jaeger-bigquery",
		Level:      log.Trace,
		JSONFormat: true,
	})

	cfgFile, err := ioutil.ReadFile(filepath.Clean(cfgPath))
	if err != nil {
		logger.Error("could not read configuration file", "cfg", cfgPath, "err", err)
		os.Exit(1)
	}

	var cfg storage.Configuration
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		logger.Error("yaml Unmarshal file err", "err", err)
	}

	var pluginServices shared.PluginServices
	store, err := storage.NewStore(logger, cfg)
	if err != nil {
		logger.Error("init store err", "err", err)
		os.Exit(1)
	}

	pluginServices.Store = store
	pluginServices.ArchiveStore = store

	grpc.Serve(&pluginServices)
	if err = store.Close(); err != nil {
		logger.Error("store close err", "err", err)
		os.Exit(1)
	}
}
