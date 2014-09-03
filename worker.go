package gost

import (
    "errors"
    "fmt"
    "log"
    "github.com/bitly/go-nsq"
)

// A worker receives tasks from a broadcaster
// and compute them, then the response is written
// in the storage system to be available when
// dest clients come for data.
type Worker interface {
    Init(gost Gost, target string, action string) error
    Stop()
    Run(Task) (Task, []byte)    // Executes the supplied task given by the broadcaster
    Store(Task, []byte) error
}

// A simple worker is a basic implementation using the NSQ broadcaster.
type NSQWorker struct {
    gost Gost
    consumer    *nsq.Consumer   // the NSQ consumer listening for work
    taskChannel chan Task
}

// TODO nsq config should be abstracted and given in a parameter
func (w *NSQWorker) Init(gost Gost, target string, action string) error {
    w.gost = gost

    // Creates the new consumer
    consumer, err := nsq.NewConsumer(target, action, nsq.NewConfig())
    w.consumer = consumer

    if err != nil {
        return err
    }

    // The worker handle the message reception
    w.consumer.AddHandler(w)

    log.Println("[WORKER] [logger] Handler attached.")

    // Connects the consumer to the broadcaster
    w.Connect(consumer, gost)

    // Inits the communication chan
    w.taskChannel = make(chan Task)

    return nil
}

// Connects with NSQ to receive task to do.
func (w *NSQWorker) Connect(consumer *nsq.Consumer, gost Gost) error {
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
        fmt.Println("[WORKER] [writer] ERROR - Unable to connect the LoggerWorker : ")
        fmt.Printf("[WORKER] [writer] ERROR - %s\n", err)
        return errors.New("Unable to start the LoggerWorker")
    }
    return nil
}

func (w *NSQWorker) Store(task Task, data []byte) error {
    // Store the result
    err := w.gost.GetStorage().Store(task.GetId(), data)
    return err
}

// Handles the message coming from NSQ, unserialize
// it to create a task, execute it in the worker and
// finally call the gost storage to store the result.
func (w *NSQWorker) HandleMessage(m *nsq.Message) error {
    task := UnserializeSimpleTask(m.Body)
    if task == nil {
        return errors.New(fmt.Sprintf("Unable to unserialize the task : %s", m.Body))
    }

    // Distribute it to the implementation
    w.taskChannel <- task

    m.Finish()
    return nil
}
