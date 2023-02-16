package main

import (
	"context"
	"food_delivery/pubsub"
	localpb "food_delivery/pubsub/local_pubsub"
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var localPubsub pubsub.Pubsub = localpb.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"
	sub1, close1 := localPubsub.Subscribe(context.Background(), topic)
	sub2, close2 := localPubsub.Subscribe(context.Background(), topic)

	localPubsub.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPubsub.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			log.Println("Sub 1: ", (<-sub1).Data())

		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			log.Println("Sub 2: ", (<-sub2).Data())

		}
	}()

	time.Sleep(time.Second * 5)
	close1()
	close2()
	localPubsub.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 5)
}
