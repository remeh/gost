package gost

import (
     "fmt"
     "github.com/bitly/go-nsq"
)

// A broadcaster which uses the NSQ
// message-queue system to broadcast tasks to worker.
type NsqBroadcaster struct {
    producer    nsq.Producer
}

func (b *NsqBroadcaster) Broadcast(task Task) error {
    err := b.producer.Publish(task.GetTarget(), task.GetData())
    return err
}

func (b *NsqBroadcaster) Init(config Config) error {
    // Use the config to know to which address
    // we want to connect for NSQ
    addr := ""

    // TODO use the list of address provided.

    if len(config.Nsqds) != 0 {
        addr = config.Nsqds[0]
    }

    if len(addr) == 0 {
       fmt.Println("[BROADCASTER] [nsq] ERROR - Can't init the nsqbroadcaster : no connect point supplied.")
    }

    // Creates the producer
    producer, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
    b.producer = *producer
    return err
}

func (b *NsqBroadcaster) Close() {
    b.producer.Stop()
}

