package gooch

// A task is delivered by a client to be
// executed on a worker.
//
// The format/type of the task should be known by either
// the src client, the worker and the dest client.
//
// It is identifier by an ID.
// For speed purpose, it is to the client to ensure the
// unicity in the system of this ID.
// If this ID isn't unique, the system should detect it,
// give the task another ID and provide to the client
// a way to know the transformation which has occured
// on its task.
type Task interface {
	GetData() []byte // Actual data of the task.
	GetType() string // identifier type of this task for clients.
}
