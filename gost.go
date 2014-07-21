package gost

import (
)

// The Gost runtime.
type Gost struct {
    controllers []Controller    // the controllers running
    broadcaster Broadcaster     // the broadcaster to use
    exitChannel chan int
}

func (g *Gost) Run() {
    g.exitChannel = make(chan int)

    g.initBroadcaster()
    g.initControllers()

    <-g.exitChannel
}

// Cleans everything and stops the runtime.
func (g *Gost) Exit() {
    g.exitChannel <- 1
}

func (g *Gost) GetBroadcaster() Broadcaster {
    return g.broadcaster
}

// Inits the controllers
func (g *Gost) initControllers() {
    // TODO configuration etc.
    g.controllers = make([]Controller, 1)
    httpController := &HttpController{gost: *g}
    httpController.Start()
    g.controllers = append(g.controllers, httpController)
}

func (g *Gost) initBroadcaster() {
    // TODO configuration etc.
    g.broadcaster = &NsqBroadcaster{}
}
