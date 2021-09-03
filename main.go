package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/liuwqiang/tykd/config"
	"github.com/liuwqiang/tykd/service"

	log "github.com/sirupsen/logrus"
)

var (
	confPaths = []string{
		"tykd.conf",
		// TODO: add ~/.config/tyk/tyk.conf here?
		"/etc/tykd/tykd.conf",
	}
)

func main() {
	svc, err := service.Spawn()
	if err != nil {
		log.Errorf("Failed to start service: err=(%s)", err)
		os.Exit(1)
	}

	globalConf := config.Config{}
	if err := config.Load(confPaths, &globalConf); err != nil {
		log.Errorf("Failed to start service: err=(%s)", err)
		os.Exit(1)
	}

	// Spawn OS signal listener to ensure graceful stop.
	osSigCh := make(chan os.Signal, 1)
	signal.Notify(osSigCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	// Wait for a quit signal and terminate the service when it is received.
	<-osSigCh
	svc.Stop()
}
