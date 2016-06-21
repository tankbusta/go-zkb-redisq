package main

import (
	"log"

	"github.com/tankbusta/go-zkb-redisq"
)

func main() {
	poller := zkbq.NewZKBPoller()

	for {
		select {
		case kill := <-poller.Kill:
			log.Println(kill)
		case err := <-poller.Error:
			log.Printf("OH NOES: %v\n", err)
		}
	}
}
