package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func ProcessWorker(in chan int, worker int) {
	for x := range in {
		t := time.Duration(rand.Intn(10) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker ", worker, ": ", int(x))
	}
}

func main() {
	log.Info("Running")
	concurrency := 8
	in := make(chan int)
	done := make(chan []byte)
	go func() {
		i := 0
		for {
			in <- 1
			i++
		}
	}()
	for x := 0; x < concurrency; x++ {
		go ProcessWorker(in, x)
	}
	<- done
}