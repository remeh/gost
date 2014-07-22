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

    // Creates the topic/channel
    /*
    conn := nsq.NewConn("localhost:4150", nsq.NewConfig(), nil)

    _, err := conn.Connect()

    if err != nil {
        panic("Unable to connect") // XXX
    }

    // Identify
    json := make(map[string]interface{})
    json["client_id"]   = "me"
    json["hostname"]    = "localhost"

    command, err := nsq.Identify(json)
    conn.WriteCommand(command)

    command = nsq.Register("printer_application", "writer")
    fmt.Println("???")
    */

    a.workers = make([]Worker, 1)
    printer := NewPrinterWorker("printer_application", "writer")
    printer.Start()
    fmt.Println("Writer application started.")
    a.workers = append(a.workers, printer)
}
