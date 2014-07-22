package gost

import (
    "fmt"
//  "github.com/bitly/go-nsq"
)

type WriterApplication struct {
//  gost        Gost            // the gost runtime instance
//  controllers []Controller    // the controllers running
    workers     []Worker        // the instanciated worker
}

func (a *WriterApplication) Start(gost Gost) {
//  a.gost = gost
//  a.initControllers()
    a.initWorkers()
}

func (a *WriterApplication) Stop() {
    for i := 0; i < len(a.workers); i++ {
        a.workers[i].Stop()
    }
}

// Inits the workers
func (a *WriterApplication) initWorkers() {
    a.workers = make([]Worker, 1)
    printer := NewPrinterWorker("printer_application", "writer")
    printer.Start()
    fmt.Println("[writer] Writer application started.")
    a.workers = append(a.workers, printer)
}
