package gost

import (
    "errors"
    "fmt"
    "github.com/bitly/go-nsq"
)

// The printer worker looks into actions
// to do, print them on the stdout and put them
// in the storage system for client to retrieve the result.
type PrinterWorker struct {
    target      string          // target represented by this worker
    action      string          // the action to listen for
    consumer    nsq.Consumer    // the NSQ consumer listening for work
}

func NewPrinterWorker(target string, action string) *PrinterWorker {
    return &PrinterWorker{target: target, action: action}
}

func (w *PrinterWorker) Start(gost Gost) error {
    // Creates the new consumer
    consumer, err := nsq.NewConsumer(w.target, w.action, nsq.NewConfig())

    if err != nil {
        return err;
    }

    // Stores it in the worker
    w.consumer = *consumer

    fmt.Println("[WORKER] [writer] Created")

    // The worker handle the message reception
    w.consumer.AddHandler(w)

    fmt.Println("[WORKER] [writer] Handler attached.")

    // Connects the worker to NSQ
    return w.Connect(gost)
}

func (w *PrinterWorker) Connect(gost Gost) error {
    config := gost.GetConfig()

    // Use the config to know to which address
    // we want to connect for NSQ
    var err error

    if len(config.Nsqlookupds) != 0 {
        err = w.consumer.ConnectToNSQLookupds(config.Nsqlookupds)
    } else {
        err = w.consumer.ConnectToNSQDs(config.Nsqds)
    }

    if err != nil {
        fmt.Println("[WORKER] [writer] ERROR - Unable to connect the PrinterWorker : ")
        fmt.Printf("[WORKER] [writer] ERROR - %s\n", err)
        return errors.New("Unable to start the PrinterWorker")
    }

    return nil
}

func (w *PrinterWorker) HandleMessage(m *nsq.Message) error {
    fmt.Printf("[WORKER] [writer] %s\n", m)
    return nil
}

func (w *PrinterWorker) Run(Task) []byte {
    // TODO
    return nil
}

func (w *PrinterWorker) Stop() {
}
