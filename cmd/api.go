package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"telegram-notification-api/internal/app"
	"telegram-notification-api/internal/clients"
	projectConfig "telegram-notification-api/internal/config"
	"telegram-notification-api/internal/dao"
	"telegram-notification-api/internal/storage"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func main() {
	config, err := projectConfig.NewConfig(projectConfig.LocalEnv)
	if err != nil {
		logger.Error("can't create config: ", err)
		return
	}

	s, err := storage.NewStorage(config.MustGetDatabaseConnectionString())
	if err != nil {
		logger.Error("can't create storage: ", err)
		return
	}

	c, err := clients.NewClients(config)
	if err != nil {
		logger.Error("can't create clients: ", err)
		return
	}

	d := dao.NewDAO(s)
	a := app.New(logger, d, c, config.MustGetServerHost(), config.MustGetServerPort())
	go func() {
		if err = a.Run(); err != nil {
			logger.Error("can't run app: ", err)
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	a.Stop()
	logger.Info("Gracefully stopped")
}
