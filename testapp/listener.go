package main

import (
	"log"

	"github.com/tankbusta/go-zkb-redisq"
)

func main() {
	poller, err := zkbq.NewZKBPoller()
	if err != nil {
		log.Fatalf("Error creating ZMQ RedisQ Poller: %v\n", err)
	}

	for {
		select {
		case kill := <-poller.Kill:
			log.Println(kill)
		case err := <-poller.Error:
			log.Printf("OH NOES: %v\n", err)
		}
	}
}
