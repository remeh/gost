package gost

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
    GetId()         string // ID of this Task, provided by the client asking to run this task.
    GetTarget()     string // Targets of this ask (workers which understand this task)
    GetAction()     string // Which action should be executed to run this task.
    GetData()       []byte // Actual data of the task.
    // XXX Don't we need a "type" or "subtype" here ?
}

// Simple implementation of a Task
// The whole data is stored in memory in a byte array.
type SimpleTask struct {
    id          string
    target      string
    action      string
    data        []byte
}

// Constructs a SimpleTask
func NewSimpleTask(id string, target string, data []byte) *SimpleTask {
    return &SimpleTask{id: id, target: target, data: data}
}

func (t *SimpleTask) GetId() string {
    return t.id
}

func (t *SimpleTask) GetTarget() string {
    return t.target
}

func (t *SimpleTask) GetAction() string {
    return t.action
}

func (t *SimpleTask) GetData() []byte {
    return t.data
}
