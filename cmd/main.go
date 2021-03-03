package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheIncredibleMulk/microservice_gokit"
)

func main() {
	var (
		httpAddr = flag.String("http", "8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	//the microservice_gokit service
	srv := microservice_gokit.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := microservice_gokit.Endpoints{
		GetEndpoint:      microservice_gokit.MakeGetEndpoint(srv),
		StatusEndpoint:   microservice_gokit.MakeStatusEndpoint(srv),
		ValidateEndpoint: microservice_gokit.MakeValidateEndpoint(srv),
	}

	// HTTP Trasport
	go func() {
		log.Println("microservice_gokit is listening on port:", *httpAddr)
		handler := microservice_gokit.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
