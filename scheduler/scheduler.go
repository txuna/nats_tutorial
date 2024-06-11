package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Run(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	_ = cancel
	_ = ctx

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	_ = nc
}
