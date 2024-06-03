package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/app"
	"github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/VladislavLisovenko/hw-vladl-prof/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)
	logg := logger.New(config.Logger.Level)

	var storage app.Storage
	switch strings.ToLower(config.Storage.Type) {
	case "postgres":
		var err error
		storage, err = sqlstorage.New(config.Storage.Postgres.DataSourceName)
		if err != nil {
			logg.Error(err.Error())
		}
	case "inmemory":
		storage = memorystorage.New()
	default:
		logg.Error("Unknown storage type")
		os.Exit(1)
	}

	calendar := app.New(logg, storage)
	server := internalhttp.NewServer(logg, calendar, config.HTTPServer.Host, config.HTTPServer.Port)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
