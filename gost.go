package gost

import (
)

// The Gost runtime.
// Has a set of applications to run.
type Gost struct {
    applications    []Application
    controllers     []Controller    // the activated controller
    storage         Storage         // the storage to use
    broadcaster     Broadcaster     // the broadcaster to use
    config          Config
    exitChannel     chan int
}

func NewGost() *Gost {
    return &Gost{applications: make([]Application, 0), exitChannel: make(chan int)}
}

func (g *Gost) Run() {
    g.exitChannel = make(chan int)

    g.config = *ReadConfig("config.yaml")

    // init the main broadcaster
    g.initBroadcaster()

    // init the storage system
    g.initStorage()

    // init all the activated controllers
    g.initControllers()

    // start all the subscribed applications
    g.startApplications()

    <-g.exitChannel
}

// Cleans everything and stops the runtime.
func (g *Gost) Exit() {
    // Closes the application
    for i := 0; i < len(g.applications); i++ {
        g.applications[i].Stop()
    }
    // Exit the main loop.
    g.exitChannel <- 1
}

func (g *Gost) GetBroadcaster() Broadcaster {
    return g.broadcaster
}

func (g *Gost) GetStorage() Storage {
    return g.storage
}

func (g *Gost) GetConfig() Config {
    return g.config
}

func (g *Gost) AddApplication(app Application) {
    g.applications = append(g.applications, app)
}

// Starts the application handled by Gost.
// Each application is launched in a go routine.
func (g *Gost) startApplications() {
    for i := 0; i < len(g.applications); i++ {
        go g.applications[i].Start(*g)
    }
}

// Inits the broadcaster
func (g *Gost) initBroadcaster() {
    if (g.config.Broadcaster == "nsq") {
        g.broadcaster = &NsqBroadcaster{}
    }
    g.broadcaster.Init(g.config)
}

// Inits the storage
func (g *Gost) initStorage() {
    if g.config.Storage == "etcd" {
        g.storage = &EtcdStorage{}
    }
    g.storage.Init(g.config)
}

// Inits the controllers
func (g *Gost) initControllers() {
    g.controllers = make([]Controller, len(g.config.Controllers))
    for _, c := range g.config.Controllers {
        switch {
            case c == "http": {
                // HTTP Controller
                httpController := &HttpController{gost: *g}
                httpController.Start()
                g.controllers = append(g.controllers, httpController)
            }
        }
    }
}
