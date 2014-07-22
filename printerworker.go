package gost

import (
    "github.com/bitly/go-nsq"
)

// The printer worker looks into actions
// to do, print them on the stdout and put them
// in the storage system for client to retrieve the result.
type PrinterWorker struct {
    target  string // target represented by this worker
    action  string // the action to listen for
}

func (w *PrinterWorker) Start() {
    // TODO consumer
}

func (w *PrinterWorker) Close() {
}
