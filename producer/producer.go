package producer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Run(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	_ = cancel

	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "game_stream",
		Subjects: []string{"game.*"},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("create the new stream: %v\n", stream)

	produce(ctx, nc, js)
}

func produce(ctx context.Context, nc *nats.Conn, js jetstream.JetStream) {
	var i int
	for {
		time.Sleep(1 * time.Second)
		if nc.Status() != nats.CONNECTED {
			log.Println("nats server closed")
			break
		}
		data := []byte(fmt.Sprintf("msg %d", i))
		if _, err := js.Publish(ctx, "game.log", data); err != nil {
			log.Println(err)
		}

		log.Printf("produce message: %s\n", string(data[:]))
		i += 1
	}
}
