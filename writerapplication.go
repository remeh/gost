package gost

type WriterApplication struct {
    gost        Gost            // the gost runtime instance
    controllers []Controller    // the controllers running
    workers     []Worker        // the instanciated worker
}

func (a *WriterApplication) Start(gost Gost) {
    a.gost = gost
    a.initControllers()
    a.initWorkers()
}

func (a *WriterApplication) Stop() {
}

// Inits the controllers
func (a *WriterApplication) initControllers() {
    // TODO configuration etc.
    a.controllers = make([]Controller, 1)
    httpController := &HttpController{gost: *a.gost}
    httpController.Start()
    a.controllers = append(a.controllers, httpController)
}

// Inits the workers
func (a *WriterApplication) initWorkers() {

}
