package gost

import (
    "fmt"

    "github.com/bitly/go-nsq"
)

// A broadcaster which uses the NSQ
// message-queue system to broadcast tasks to worker.
type NsqBroadcaster struct {
}

func (b *NsqBroadcaster) Broadcast(task Task) error {
    // TODO 
    //message := nsq.NewMessage(task.GetId(), task.GetData())
    command := nsq.Publish(task.GetTopic(), task.GetData())
    fmt.Println(command)
    return nil
}

func (b *NsqBroadcaster) Init() {
}

func (b *NsqBroadcaster) Close() {
}

