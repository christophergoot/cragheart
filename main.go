package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/customerio/homework/datastore"
	"github.com/customerio/homework/serve"
	"github.com/customerio/homework/stream"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()

	file, err := os.Open("data/initial_data.json")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := stream.Process(ctx, file)
	if err != nil {
		log.Fatal(err)
	}

	ds, err := datastore.NewDatastoreFromChannel(ch)

	if err := ctx.Err(); err != nil {
		log.Fatal(err)
	}

	if ds == nil {
		log.Fatal("you need to implement the serve.Datastore interface to run the server")
	}

	if err := serve.ListenAndServe(":1323", ds); err != nil {
		log.Fatal(err)
	}
}
