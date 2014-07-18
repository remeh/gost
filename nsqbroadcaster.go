package gooch

import (
    "github.com/bitly/go-nsq"
)

// A broadcaster which uses the NSQ
// message-queue system to broadcast tasks to worker.
type NsqBroadcaster struct {

}

func (b *NsqBroadcaster) Broadcast(task *Task) {
    // TODO 
    message := nsq.NewMessage(task.GetId(), task.GetData());
}

