Gost
=====

A distributed realtime computation system for Go applications.
Using NSQ to broadcast and load-balance tasks pushed on one of the Gost controllers (http, rpc, ...), Gost could be a basic subsitute to Apache Storm to create scalable Go backend applications.

# Dependency

```
  go get github.com/remeh/go-webserver
  go get github.com/bitly/go-nsq
  go get gopkg.in/yaml.v1
  go get github.com/coreos/go-etcd/etcd
```

# TODO

  * A Logger format.
  * Document an example.
