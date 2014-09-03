package gost

import (
    "log"
)

// The printer worker looks into actions
// to do, print them on the stdout and put them
// in the storage system for client to retrieve the result.
type LoggerWorker struct {
    NSQWorker
    target      string          // target represented by this worker
    action      string          // the action to listen for
    gost        Gost            // instance of the Gost service
}

func NewLoggerWorker(target string, action string) *LoggerWorker {
    return &LoggerWorker{target: target, action: action}
}

func (w *LoggerWorker) Start(gost Gost) error {
    w.gost = gost

    err := w.Init(gost, w.target, w.action)

    if err != nil {
        return err
    }

    log.Println("[WORKER] [logger] Created")

    for {
        // Wait for a task to execute and execute it
        task := <-w.taskChannel
        task, data := w.Run(task)

        // Stores the data.
        w.Store(task, data)
    }

    return nil
}

func (w *LoggerWorker) Run(task Task) (Task, []byte) {
    log.Printf("[WORKER] [logger] [Target: %s] [Action: %s] %s\n", task.GetTarget(), task.GetAction(), task.GetData())
    return task, task.GetData()
}

func (w *LoggerWorker) Stop() {
    // Nothing to do
}
