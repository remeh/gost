package gost

import (
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
    consumer, err := nsq.NewConsumer(w.target, w.action, nsq.NewConfig())

    if err != nil {
        return err;
    }

    w.consumer = *consumer

    fmt.Println("[printer worker] Started")

    handler := nsq.HandlerFunc(func(m *nsq.Message) error {
        fmt.Printf("[writer] : %s\n", m)
        return nil
    })

    consumer.AddHandler(handler)

    // Connects the worker to NSQ
    return w.Connect(gost)
}

func (w *PrinterWorker) Connect(gost Gost) error {
    config := gost.GetConfig()

    // Use the config to know to which address
    // we want to connect for NSQ

    // TODO use the list of address provided.
    addr := ""

    if len(config.Nsqds) != 0 {
        addr = config.Nsqds[0]
    } else {
        addr = config.Nsqlookupds[0]
    }

    if len(addr) == 0 {
       fmt.Println("[gost] [ERROR] : can't connect the printer worker, no connect point supplied.")
    }

    // Finally, connect.
    err := w.consumer.ConnectToNSQLookupd(addr)

    if err != nil {
        fmt.Println("[error] Unable to connect the PrinterWorker")
        return err
    }

    return nil
}

func (w *PrinterWorker) Run(Task) []byte {
    return nil
}

func (w *PrinterWorker) Stop() {
}
