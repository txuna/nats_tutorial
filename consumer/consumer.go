package consumer

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Run(url string, index int) {
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

	log.Printf("[%d] create the new stream: %v\n", index, stream)

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "game_stream",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	if err != nil {
		log.Fatal(err)
	}

	consume(index, cons)
}

func consume(index int, cons jetstream.Consumer) {
	cc, err := cons.Consume(func(msg jetstream.Msg) {
		log.Printf("[%d] consume message: %s", index, string(msg.Data()[:]))
		msg.Ack()
	}, jetstream.ConsumeErrHandler(func(consumeCtx jetstream.ConsumeContext, err error) {
		log.Println(err)
	}))

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Stop()

	select {}
}
