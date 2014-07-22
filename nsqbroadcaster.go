package gost

import (
        "fmt"
    "github.com/bitly/go-nsq"
)

// A broadcaster which uses the NSQ
// message-queue system to broadcast tasks to worker.
type NsqBroadcaster struct {
    producer nsq.Producer
}

func (b *NsqBroadcaster) Broadcast(task Task) error {
    err := b.producer.Publish(task.GetTarget(), task.GetData())
    return err
}

func (b *NsqBroadcaster) Init() error {
    config := nsq.NewConfig()
    producer, err := nsq.NewProducer("localhost:4160", config)
    fmt.Println(b.producer)
    b.producer = *producer
    return err
}

func (b *NsqBroadcaster) Close() {
}

