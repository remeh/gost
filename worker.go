package gost

import (
    "errors"
    "fmt"
    "github.com/bitly/go-nsq"
)

// A worker receives tasks from a broadcaster
// and compute them, then the response is written
// in the storage system to be available when
// dest clients come for data.
type Worker interface {
    Init(gost Gost, target string, action string, config nsq.Config) (*nsq.Consumer, error)
    Stop()
    Run(Task) (Task, []byte)    // Executes the supplied task given by the broadcaster
}

// A simple worker is a basic implementation using the NSQ broadcaster.
type SimpleWorker struct {
    gost Gost
}

// TODO nsq config should be abstracted
func (w *SimpleWorker) Init(gost Gost, target string, action string, config nsq.Config) (*nsq.Consumer, error) {
    w.gost = gost

    // Creates the new consumer
    consumer, err := nsq.NewConsumer(target, action, &config)

    if err != nil {
        return nil, err
    }

    return consumer, nil
}

// Connects with NSQ to receive task to do.
func (w *SimpleWorker) Connect(consumer *nsq.Consumer, gost Gost) error {
    config := gost.GetConfig()

    // Use the config to know to which address
    // we want to connect for NSQ
    var err error

    if len(config.Nsqlookupds) != 0 {
        err = consumer.ConnectToNSQLookupds(config.Nsqlookupds)
    } else {
        err = consumer.ConnectToNSQDs(config.Nsqds)
    }

    if err != nil {
        fmt.Println("[WORKER] [writer] ERROR - Unable to connect the PrinterWorker : ")
        fmt.Printf("[WORKER] [writer] ERROR - %s\n", err)
        return errors.New("Unable to start the PrinterWorker")
    }
    return nil
}


func (w *PrinterWorker) HandleMessage(m *nsq.Message) error {
    task := UnserializeSimpleTask(m.Body)
    if task == nil {
        return errors.New(fmt.Sprintf("Unable to unserialize the task : %s", m.Body))
    }

    // Retrieve the result of the task execution
    returnedTask, result := w.Run(task)

    // Store the result
    w.gost.GetStorage().Store(returnedTask.GetId(), result)

    m.Finish()
    return nil
}

