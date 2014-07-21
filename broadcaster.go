package gost

// A broadcaster is the part dealing with receiving
// task and sending them to all the intersted workers.
type Broadcaster interface {
    Init()
    Close()
    Broadcast(Task) error                  // Do your job, broadcaster.
    // TODO Register(Channel, Topic) 
}
