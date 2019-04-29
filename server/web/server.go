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
	"github.com/qinhan-shu/gp-server/service/analysis"
	"github.com/qinhan-shu/gp-server/service/announcement"
	"github.com/qinhan-shu/gp-server/service/auth"
	"github.com/qinhan-shu/gp-server/service/class"
	"github.com/qinhan-shu/gp-server/service/config"
	"github.com/qinhan-shu/gp-server/service/gate"
	"github.com/qinhan-shu/gp-server/service/judge"
	"github.com/qinhan-shu/gp-server/service/match"
	"github.com/qinhan-shu/gp-server/service/problem"
	"github.com/qinhan-shu/gp-server/service/rank"
	"github.com/qinhan-shu/gp-server/service/user"
	"github.com/qinhan-shu/gp-server/utils/parse"
)

var (
	tag    string
	commit string
	branch string

	version = flag.Bool("version", false, "show version") // show version
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(formatFullVersion())
		return
	}

	// get config
	c := config.NewConfig()
	dataStorage, err := c.GetDataStorage()
	if err != nil {
		logger.Sugar.Fatalf("failed to get data storage : %v", err)
	}

	// init logger
	logLevel, isExist := dataStorage.Configs.Load("LOG_LEVEL")
	if !isExist {
		logLevel = "info"
	}
	logger.InitLogger(parse.String(logLevel))

	// create gate
	port, isExist := dataStorage.Configs.Load("WEB_PORT")
	if !isExist {
		port = ":8081"
	}
	gateService := gate.NewService(parse.String(port))
	certFile, isCertFileExist := dataStorage.Configs.Load("CERTFILE")
	keyFile, isKeyFileExist := dataStorage.Configs.Load("KEYFILE")
	if isCertFileExist && isKeyFileExist {
		c := parse.String(certFile)
		k := parse.String(keyFile)
		if c != "" && k != "" {
			logger.Sugar.Infof("TSL : certFile[%s], keyFile[%s]", c, k)
			gateService.AddTLSConfig(c, k)
		}
	}

	// register modules
	registerModule(gateService, dataStorage)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// start service
	go func() {
		<-sigChan
		logger.Sugar.Infof("Shutting down gate server...")
		gateService.Stop()
	}()

	logger.Sugar.Infof("Starting gate server on %s", parse.String(port))
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

func registerModule(gate module.Gate, dataStorage *module.DataStorage) {
	auth.Register(gate, dataStorage)
	problem.Register(gate, dataStorage)
	class.Register(gate, dataStorage)
	user.Register(gate, dataStorage)
	rank.Register(gate, dataStorage)
	announcement.Register(gate, dataStorage)
	match.Register(gate, dataStorage)
	judge.Register(gate, dataStorage)
	analysis.Register(gate, dataStorage)
}
