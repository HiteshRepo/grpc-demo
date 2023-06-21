package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hiteshrepo/grpc-demo/internal/apps/restapp"
)

func main() {
	app := restapp.NewRestApp()
	app.Start()

	<-interrupt()

	app.Stop()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
