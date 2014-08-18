package gost

import (
    "log"

    "github.com/bitly/go-nsq"
)

// The printer worker looks into actions
// to do, print them on the stdout and put them
// in the storage system for client to retrieve the result.
type PrinterWorker struct {
    SimpleWorker
    target      string          // target represented by this worker
    action      string          // the action to listen for
    consumer    *nsq.Consumer   // the NSQ consumer listening for work
    gost        Gost            // instance of the Gost service
}

func NewPrinterWorker(target string, action string) *PrinterWorker {
    return &PrinterWorker{target: target, action: action}
}

func (w *PrinterWorker) Start(gost Gost) error {
    w.gost = gost

    consumer, err := w.Init(gost, w.target, w.action, *nsq.NewConfig())

    if err != nil {
        return err
    }

    // Stores it in the worker
    w.consumer = consumer

    log.Println("[WORKER] [logger] Created")

    // The worker handle the message reception
    w.consumer.AddHandler(w)

    log.Println("[WORKER] [logger] Handler attached.")

    // Connects the consumer to the broadcaster
    w.Connect(consumer, gost)

    // Connects the worker to NSQ
    return nil
}

func (w *PrinterWorker) Run(task Task) (Task, []byte) {
    log.Printf("[WORKER] [logger] [Target: %s] [Action: %s] %s\n", task.GetTarget(), task.GetAction(), task.GetData())
    return task, task.GetData()
}

func (w *PrinterWorker) Stop() {
}
