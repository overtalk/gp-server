package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/QHasaki/Server/logger"
	"github.com/QHasaki/Server/model/v1"
	"github.com/QHasaki/Server/service/config/v1"
	"github.com/QHasaki/Server/service/config/v1/configSource/github"
	"github.com/QHasaki/Server/service/gate/v1"
)

func main() {
	var (
		debug    bool
		addr     string
		httpDir  string
		certFile string
		keyFile  string
	)

	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&addr, "addr", ":5353", "listen address")
	flag.StringVar(&httpDir, "httpDir", "/tmp", "root dir for http fileserver")
	flag.StringVar(&certFile, "certFile", "", "ssl certficate filename")
	flag.StringVar(&keyFile, "keyFile", "", "ssl private key filename")
	flag.Parse()

	conf := getConfig()
	dataStorage, err := conf.GetDataStorage()
	if err != nil {
		logger.Sugar.Fatalf("failed to get dataStorage : %v", err)
	}

	gateService := gate.NewService(addr, httpDir, dataStorage)
	if certFile != "" && keyFile != "" {
		gateService.AddTLSConfig(certFile, keyFile)
	}

	if debug {
		l, err := zap.NewDevelopment(
			zap.AddStacktrace(zapcore.PanicLevel),
		)
		if err != nil {
			log.Fatal("failed to initialize zap logger")
		}
		logger.Sugar = l.Sugar()
		log.Printf("Debug is on\n")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down gate server...")
		gateService.Stop()
	}()

	log.Printf("Starting gate server on %s\n", addr)
	gateService.Start()
}

func getConfig() model.Config {
	confSource := configSource.NewGithub()
	conf := config.NewConfig(confSource)

	return conf
}
