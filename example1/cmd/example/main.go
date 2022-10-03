package main

import (
	"context"
	"example1"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "example1"

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

	d := example1.NewExample1()
	svr := api.NewExample1ToldataServer(bus, d)
	wait, err := svr.SubscribeExample1()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	go func() {
		numLoops := 100
		timePerIteration := time.Duration(3) * time.Second
		ticker := time.NewTicker(timePerIteration)
		for i := 0; i < numLoops; i++ {
			// your code
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			d.Say(ctx, &example1.Hello{Name: "Carmack"})
			<-ticker.C
		}
		ticker.Stop()
	}()

	go func() {
		_ = <-sigs
		done <- true
	}()

	log.Println("Example1 started")
	<-done
	cancel()
	log.Println("Example1 stopped")
	<-wait

}
