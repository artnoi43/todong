package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/artnoi43/todong/config"
	"github.com/artnoi43/todong/lib/handler"
	"github.com/artnoi43/todong/lib/httpserver"
	"github.com/artnoi43/todong/lib/store"
)

var (
	conf        *config.Config
	dataGateway store.DataGateway
	server      httpserver.Server
)

func init() {
	var err error
	conf, err = config.LoadConfig("config")
	if err != nil {
		log.Fatalln("error: failed to load config:", err.Error())
	}
	dataGateway = store.Init(conf)
	if dataGateway == nil {
		log.Fatalln("nil dataGateway")
	}
	// Init server and middleware
	handlerAdaptor := handler.NewAdaptor(conf.Server, dataGateway, &conf.Middleware)
	server = httpserver.New(conf.Server)
	server.SetUpRoutes(&conf.Middleware, handlerAdaptor)
}

func main() {
	// sigChan is for receiving os.Signal from the host OS.
	// Graceful shutdowns are tested on macOS and Arch Linux
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	// Wrap server.Serve() in goroutine so that we can have graceful shutdown
	// and server concurrently listening.
	go func() {
		log.Printf("Server started on %s", conf.Address)
		log.Fatal(server.Serve(conf.Address))
	}()

	// main() will block here, waiting for value to be received from sigChan
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		log.Println("Shutting down server and data store")
		// TODO: httpserver.Shutdown does not work yet
		server.Shutdown(context.Background())
		dataGateway.Shutdown()
		log.Println("Server and data store shutdown gracefully")
	}()
	wg.Wait()
	os.Exit(0)
}
