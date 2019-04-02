package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/module"
	"github.com/qinhan-shu/gp-server/service/auth"
	"github.com/qinhan-shu/gp-server/service/class"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/gate"
	"github.com/qinhan-shu/gp-server/service/problem"
	"github.com/qinhan-shu/gp-server/service/rank"
	"github.com/qinhan-shu/gp-server/service/user"
)

var (
	tag    string
	commit string
	branch string

	version  = flag.Bool("version", false, "show version") // show version
	port     = flag.String("addr", ":8080", "listen address")
	certFile = flag.String("certFile", "", "ssl certficate filename")
	keyFile  = flag.String("keyFile", "", "ssl private key filename")
	logLevel = flag.String("log-level", "error", "log level, optional( debug | info | warn | error | dpanic | panic | fatal), default is error")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(formatFullVersion())
		return
	}

	// init logger
	logger.InitLogger(*logLevel)

	gateService := gate.NewService(*port)
	if *certFile != "" && *keyFile != "" {
		gateService.AddTLSConfig(*certFile, *keyFile)
	}

	// register modules
	registerModule(gateService)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Sugar.Infof("Shutting down gate server...")
		gateService.Stop()
	}()

	logger.Sugar.Infof("Starting gate server on %s", *port)
	gateService.Start()
}

func formatFullVersion() string {
	var parts = []string{"gp_server"}

	if tag != "" {
		parts = append(parts, tag)
	} else {
		parts = append(parts, "(tag: unknown)")
	}

	if branch != "" || commit != "" {
		if branch == "" {
			branch = "unknown_branch"
		}
		if commit == "" {
			commit = "unknown_commit"
		}
	}
	git := fmt.Sprintf("(git: %s %s)", branch, commit)
	parts = append(parts, git)

	return strings.Join(parts, "  ")
}

func registerModule(gate module.Gate) {
	c := config.NewConfig()
	dataStorage, err := c.GetDataStorage()
	if err != nil {
		logger.Sugar.Fatalf("failed to get data storage : %v", err)
	}

	auth.Register(gate, dataStorage)
	problem.Register(gate, dataStorage)
	class.Register(gate, dataStorage)
	user.Register(gate, dataStorage)
	rank.Register(gate, dataStorage)
}
