package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"server/api/rest"
	"server/config"
	"server/pkg/log"
	"syscall"
)

func main() {
	c := config.GetConfig()
	log.InitLog(c.LogPath, c.LogLevel)

	go rest.InitRouter(c.Port)

	done := make(chan os.Signal)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(done)
	}()
	<-done

	logrus.Info("Server closed")
}
