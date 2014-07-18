package gooch

// A broadcaster is the part dealing with receiving
// task and sending them to all the intersted workers.
type Broadcaster interface {
    Broadcast(*Task) error      // Do your job, broadcaster.
    // TODO Register(Channel, Topic) 
}
