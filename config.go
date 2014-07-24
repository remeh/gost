package gost

import (
    "fmt"
    "io/ioutil"

    "gopkg.in/yaml.v1"
)

const (
    BROADCASTER_NSQ = "nsq"

    CONTROLLER_HTTP = "http"
    CONTROLLER_TCP  = "tcp"
)

type Config struct {
    Broadcaster string   // which broadcaster should we use.
    Controllers []string // which controllers should be started.

    Nsqlookupds []string // when we want to connect to nslookupds
    Nsqds []string       // when we want to connect directly to nsqds
}

func ReadConfig(filename string) *Config {
    // Reads the file 
    content, err := ioutil.ReadFile(filename)

    if err != nil {
        return ConfigError(filename, err)
    }

    // Decodes the YAML
    var c Config
    yaml.Unmarshal(content, &c)

    fmt.Println("[gost] Read configuration : ")
    fmt.Println(c)

    return &c
}

func DefaultConfig() *Config {
    // TODO return a default config
    return nil
}

func ConfigError(filename string, err error) *Config {
    fmt.Println("[gost] Error reading the config file : " + filename);
    fmt.Println(err)
    fmt.Println("[gost] Will use default configuration.")
    return DefaultConfig()
}
