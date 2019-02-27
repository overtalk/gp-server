package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/example"
	"github.com/qinhan-shu/gp-server/service/gate/gin"
)

func main() {
	var (
		debug    bool
		addr     string
		certFile string
		keyFile  string
	)

	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.StringVar(&addr, "addr", ":5353", "listen address")
	flag.StringVar(&certFile, "certFile", "", "ssl certficate filename")
	flag.StringVar(&keyFile, "keyFile", "", "ssl private key filename")
	flag.Parse()

	if debug {
		logger.AddDebugLogger()
	}

	gateService := gate.NewService(addr)
	if certFile != "" && keyFile != "" {
		gateService.AddTLSConfig(certFile, keyFile)
	}

	// 注册具体的模块
	registerModule(gateService)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Sugar.Infof("Shutting down gate server...")
		gateService.Stop()
	}()

	logger.Sugar.Infof("Starting gate server on %s", addr)
	gateService.Start()
}

func registerModule(gate module.Gate) {

	example.Register(gate)
}
