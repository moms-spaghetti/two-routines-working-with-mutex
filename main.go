package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT)

	store := &[]string{}
	go func() {
		for {
			time.Sleep(3 * time.Second)
			save(store, "SLOW routine item")

		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			save(store, "fast routine item")
		}
	}()

	<-wait
	close(wait)
	log.Print("shutdown ok ")
}

func save(store *[]string, item string) {
	mutex := sync.Mutex{}
	mutex.Lock()
	*store = append(*store, item)
	mutex.Unlock()

	for _, s := range *store {
		fmt.Println(s)
	}
}
