package gost

import (
    "fmt"
    "github.com/bitly/go-nsq"
)

// The printer worker looks into actions
// to do, print them on the stdout and put them
// in the storage system for client to retrieve the result.
type PrinterWorker struct {
    target  string // target represented by this worker
    action  string // the action to listen for
}

func NewPrinterWorker(target string, action string) *PrinterWorker {
    return &PrinterWorker{target: target, action: action}
}

func (w *PrinterWorker) Start() error {
    consumer, err := nsq.NewConsumer(w.target, w.action, nsq.NewConfig())

    // XXX
    consumer.ConnectToNSQLookupd("localhost:4160")

    if err != nil {
        return err;
    }

    fmt.Println("[printer worker] Started")

    handler := nsq.HandlerFunc(func(m *nsq.Message) error {
        fmt.Printf("[writer] : %s\n", m)
        return nil
    })

    consumer.AddHandler(handler)

    return nil
}

func (w *PrinterWorker) Run(Task) []byte {
    return nil
}

func (w *PrinterWorker) Stop() {
}
