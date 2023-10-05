package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"mckp/helper/datastore"
	internalServer "mckp/helper/http"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)

	switch log_level := os.Getenv("LOG_LEVEL"); log_level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	datastore.DatastoreInstance.Connect()
	datastore.DatastoreInstance.EnsureMigration()
}

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	router, err := internalServer.New()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	srv.ListenAndServe()

	<-gracefulShutdown

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	handleGracefulShutdown(cancel, srv)
}

func handleGracefulShutdown(cancel context.CancelFunc, server *http.Server) {
	log.Trace("We are being requested to shutdown.")
	datastore.DatastoreInstance.Disconnect()
	server.Close()
}
