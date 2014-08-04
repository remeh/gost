package gost

import (
    "bytes"
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
    Serialize()     []byte // Serialize the task to send it on the broadcaster
}

// Simple implementation of a Task
// The whole data is stored in memory in a byte array.
type SimpleTask struct {
    id          string
    target      string
    action      string
    data        []byte
}

// Constructs a SimpleTask.
// A simple task is serialized in binary for quicker serialization/deserialization
// and a tiny weight.
func NewSimpleTask(id string, target string, action string, data []byte) *SimpleTask {
    if len(id) == 0 {
        return nil
    }
    if len(target) == 0 {
        return nil
    }
    if len(action) == 0 {
        return nil
    }

    return &SimpleTask{id: id, target: target, action: action, data: data}
}

// Format of a simple task:
// 2 bytes      : magic number
// 8 bytes      : id length : l
// 8 bytes      : data length : n
// 32 bytes     : target
// 32 bytes     : action
// l bytes      : id
// n bytes      : data
func UnserializeSimpleTask(data []byte) *SimpleTask {

    // Magic number 01
    if data[0] != 0 || data[1] != 1 {
        return nil
    }

    // ID Length
    idLength, err := binary.ReadUvarint(bytes.NewBuffer(data[2:10]))
    if err != nil {
        return nil
    }

    // Data Length
    length, err := binary.ReadUvarint(bytes.NewBuffer(data[10:18]))
    if err != nil {
        return nil
    }

    // Target
    target := string(data[18:50])

    // Action
    action := string(data[50:82])

    // ID
    readId := string(data[82:82+idLength])

    // Data
    readData := data[82+idLength:82+idLength+length]

    return NewSimpleTask(readId, target, action, readData)
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
// 8 bytes      : id length : l
// 8 bytes      : data length : n
// 32 bytes     : target
// 32 bytes     : action
// l bytes      : id
// n bytes      : data
func (t *SimpleTask) Serialize() []byte {
    serialized := make([]byte, 2 + 8 + 8 + 32 + 32 + len(t.id) + len(t.data))

    // Magic number
    serialized[0] = 0
    serialized[1] = 1

    // Task ID length
    binary.PutUvarint(serialized[2:], uint64(len(t.id)))

    // Task data length
    binary.PutUvarint(serialized[10:], uint64(len(t.data)))

    // Target
    copy(serialized[18:], []byte(t.target))

    // Action
    copy(serialized[50:], []byte(t.action))

    // Data
    copy(serialized[82:], t.id)

    // Data
    copy(serialized[82+len(t.id):], t.data)

    return serialized
}
