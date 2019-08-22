package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/olivere/env"

	"github.com/olivere/chrono/server"
)

var (
	addr     = flag.String("addr", env.String(":8080", "ADDR"), "HTTP address to bind to")
	loglevel = flag.String("loglevel", env.String("info", "LOGLEVEL"), "Log level: error, warn, info, or debug")
)

func main() {
	flag.Parse()
	logger := createLogger()
	if err := runMain(logger); err != nil {
		logger.Log("err", err)
	}
}

func runMain(logger log.Logger) error {
	srv := server.New(
		server.WithLogger(logger),
	)
	httpSrv := &http.Server{
		Addr:    *addr,
		Handler: srv,
	}
	logger.Log("msg", "Starting server", "addr", *addr)
	return httpSrv.ListenAndServe()
}

func createLogger() log.Logger {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "@time", log.DefaultTimestampUTC)
	switch *loglevel {
	default:
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "info":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "warn":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "error":
		logger = level.NewFilter(logger, level.AllowDebug())
	}
	return logger
}
