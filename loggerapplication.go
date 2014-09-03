package gost

import (
    "fmt"
)

type LoggerApplication struct {
    workers     []Worker        // the instanciated worker
}

func NewLoggerApplication() *LoggerApplication {
    return &LoggerApplication{}
}

func (a *LoggerApplication) Start(gost Gost) {
    a.initWorkers(gost)
}

func (a *LoggerApplication) Stop() {
    for i := 0; i < len(a.workers); i++ {
        a.workers[i].Stop()
    }
}

// Inits the workers
func (a *LoggerApplication) initWorkers(gost Gost) {
    a.workers = make([]Worker, 1)
    printer := NewLoggerWorker("printer_application", "logger")
    err := printer.Start(gost)
    if err != nil {
        fmt.Println("[APPLICATION] [logger] ERROR - while starting the workers of LoggerApplication : ")
        fmt.Printf("[APPLICATION] [logger] ERROR - %s\n", err)
    } else {
        fmt.Println("[APPLICATION] [logger] Logger application started.")
        a.workers = append(a.workers, printer)
    }
}
