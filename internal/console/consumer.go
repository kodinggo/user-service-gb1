package console

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kodinggo/user-service-gb1/internal/model"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consumerCmd)
}

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the consumer server",
	Run:   consumerServer,
}

func consumerServer(cmd *cobra.Command, args []string) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		logrus.Fatal(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx := context.Background()
	stream, err := js.Stream(ctx, "USER")
	if err != nil {
		logrus.Fatal(err)
	}

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       "USER",
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "USER.*",
	})
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = cons.Consume(func(msg jetstream.Msg) {
		err := msg.Ack()
		if err != nil {
			logrus.Fatal(err)
		}

		var user model.User
		err = json.Unmarshal(msg.Data(), &user)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println("data consumed!")
		fmt.Println("name: ", user.Name)
		fmt.Println("email: ", user.Email)

		// ORDER.create
		// create payment
	}, jetstream.ConsumeErrHandler(nil),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("consumer server started...")
	select {}
}
