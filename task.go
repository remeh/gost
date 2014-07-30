package gost

import (
    "strings"
    "encoding/binary"
)

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
    Serialize()     []byte
    //Unserialize()   []byte
}

// Simple implementation of a Task
// The whole data is stored in memory in a byte array.
type SimpleTask struct {
    id          string
    target      string
    action      string
    data        []byte
}

// Constructs a SimpleTask, takes an UUID as identifier.
// A simple task is serialized in binary for quicker serialization/deserialization
// and a tiny weight.
func NewSimpleTask(uuid string, target string, action string, data []byte) *SimpleTask {
    withoutHyphen := strings.Replace(uuid, "-", "", -1)

    if len(withoutHyphen) > 32 {
        println(withoutHyphen)
        println(len(withoutHyphen))
        return nil
    }
    if len(target) > 32 {
        return nil
    }
    if len(action) > 32 {
        return nil
    }

    return &SimpleTask{id: withoutHyphen, target: target, action: action, data: data}
}

func UnserializeSimpleTask(data []byte) *SimpleTask {
    // Magic number 01
    if data[0] != 0 || data[1] != 1 {
        return nil
    }

    // TODO

    return nil
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

// Format of a simple task:
// 2 bytes      : magic number
// 32 bytes     : task ID
// 32 bytes     : target
// 32 bytes     : action
// 8 bytes      : data length : n
// n bytes      : data
func (t *SimpleTask) Serialize() []byte {
    serialized := make([]byte, 2 + 32 + 32 + 32 + 8 + len(t.data))

    // Magic number
    copy(serialized[0:], []byte("01"))

    // Task ID
    copy(serialized[2:], []byte(t.id));

    // Target
    copy(serialized[34:], []byte(t.target))

    // Action
    copy(serialized[66:], []byte(t.action))

    // Len
    binary.PutUvarint(serialized[98:], uint64(len(t.data)))

    // Data
    copy(serialized[106:], t.data)

    return serialized
}
