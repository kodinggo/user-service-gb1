package pubsub

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func InitJetstream() (jetstream.JetStream, error) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	_, err = js.CreateStream(nats.Context(ctx), jetstream.StreamConfig{
		Name:     "USER",
		Subjects: []string{"USER.*"},
	})
	if err != nil {
		return nil, err
	}

	return js, nil
}
