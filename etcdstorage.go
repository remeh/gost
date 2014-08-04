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

func (s *EtcdStorage) Read(id string) []byte {
    response, err := s.client.Get(id, false, false)

    if err != nil || response == nil {
        return []byte("")
    }

    return []byte(response.Node.Value)
}

func (s *EtcdStorage) Store(id string, data []byte) {
    _, err := s.client.Set(id, string(data), 0)

    // TODO err return
    if err != nil {
        fmt.Printf("[STORAGE] [ETCD] ERROR - Unable to storage id["+id+"], cause : %s\n", err)
    }
}
