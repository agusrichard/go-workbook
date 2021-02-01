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

	"learn_gokit"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our learn_gokit service
	srv := learn_gokit.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := learn_gokit.Endpoints{
		GetEndpoint:      learn_gokit.MakeGetEndpoint(srv),
		StatusEndpoint:   learn_gokit.MakeStatusEndpoint(srv),
		ValidateEndpoint: learn_gokit.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("learn_gokit is listening on port:", *httpAddr)
		handler := learn_gokit.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
