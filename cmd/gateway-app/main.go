package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hiteshrepo/grpc-demo/internal/apps/gatewayapp"
)

func main() {
	a := gatewayapp.NewGatewayApp()
	a.Start()
	<-interrupt()
	a.Shutdown()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
