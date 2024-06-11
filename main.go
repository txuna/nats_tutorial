package main

import (
	"jet/consumer"
	"jet/producer"
	"jet/scheduler"
	"sync"
)

const (
	NATS_ADDR = "127.0.0.1:4222"
)

func main() {
	wq := &sync.WaitGroup{}

	wq.Add(4)

	go func() {
		defer wq.Done()
		producer.Run(NATS_ADDR)
	}()

	go func() {
		defer wq.Done()
		consumer.Run(NATS_ADDR, 1)
	}()

	go func() {
		defer wq.Done()
		consumer.Run(NATS_ADDR, 2)
	}()

	go func() {
		defer wq.Done()
		scheduler.Run(NATS_ADDR)
	}()

	wq.Wait()
}
