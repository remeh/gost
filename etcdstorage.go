package gost

import (
    "errors"
    "fmt"

    "github.com/coreos/go-etcd/etcd"
)

type EtcdStorage struct {
    client *etcd.Client
}

func (s *EtcdStorage) Init(config Config) error {
    if len(config.Etcds) == 0 {
        errors.New("No Etcds host provided.")
    }
    // Connect
    s.client = etcd.NewClient(config.Etcds)

    fmt.Printf("[STORAGE] [ETCD] INFO - Etcd client created with hosts : %s\n", config.Etcds)

    return nil
}

// Reads in Etcd with the given ID and returns
// what have be found.
func (s *EtcdStorage) Read(id string) []byte {
    response, err := s.client.Get(id, false, false)

    if err != nil || response == nil {
        return []byte("")
    }

    return []byte(response.Node.Value)
}

// Stores the given data into the given ID in Etcd.
func (s *EtcdStorage) Store(id string, data []byte) error {
    _, err := s.client.Set(id, string(data), 0)

    if err != nil {
        return err
    }
}
