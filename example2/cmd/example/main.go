package main

import (
	"context"
	"example2"
	"log"
	"os"
	"os/signal"
	"syscall"

	api "example2"

	"github.com/citradigital/toldata"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_ = godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	bus, err := toldata.NewBus(ctx, toldata.ServiceConfiguration{URL: "localhost:4222"})
	if err != nil {
		log.Fatalln(err)
	}
	d := example2.NewExample2()
	svr := api.NewExample2ToldataServer(bus, d)
	wait, err := svr.SubscribeExample2()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	go func() {
		_ = <-sigs
		done <- true
	}()

	log.Println("Example2 started")
	<-done
	cancel()
	log.Println("Example2 stopped")
	<-wait

}
