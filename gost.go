package gost

import (
)

// The Gost runtime.
// Has a set of applications to run.
type Gost struct {
    applications    []Application
    broadcaster     Broadcaster     // the broadcaster to use
    exitChannel     chan int
}

func NewGost() {
    return &Gost{applications: make([]Application, 1), exitChannel: make(chan int)}
}

func (g *Gost) Run() {
    g.exitChannel = make(chan int)

    g.startApplications()

    <-g.exitChannel
}

// Cleans everything and stops the runtime.
func (g *Gost) Exit() {
    // Closes the application
    for i := 0; i < len(applications); i++ {
        applications[i].Stop()
    }
    // Exit the main loop.
    g.exitChannel <- 1
}

func (g *Gost) GetBroadcaster() Broadcaster {
    return g.broadcaster
}

func (g *Gost) AddApplication(app Application) {
    applications = append(applications, app)
}

// Starts the application handled by Gost.
// Each application is launched in a go routine.
func (g *Gost) startApplications() {
    for i := 0; i < len(applications); i++ {
        go applications[i].Start(*g)
    }
}
