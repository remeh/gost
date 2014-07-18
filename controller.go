package gooch

// A controller in gooch is not the same controller as in MVC.
// The role of a controller here is to broadcast tasks supplied
// by clients to workers which should be interested.
//
// A controller can return data after a call, but it shouldn't be
// considered or mandatory at all, that's the main process of
// Gooch.
type Controller interface {
    Start() // Start method for the controller. Called when Gooch is initializated.
    Close() // TODO
}
