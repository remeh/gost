package gost

// An application is handled by Gost.
// It has 0 or many controllers to answer to client
// It has 0 or many workers to process the task 
// given by the broadcaster.
// It has 0 or many suppliers to return back the task
// to clients.
type Application interface {
    Start(Gost)                     // Init method of the application
    Stop()
    GetControllers() []Controller   // The controllers of the application.
    GetWorkers() []Worker           // The workers of the application
}
