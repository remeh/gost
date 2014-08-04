package gost

// The storage is where the data is written
// to be retrieved by the ResponseProviders.
type Storage interface {
    Init(config Config) error
    Read(id string) []byte  // Returns the content of the task which has the given ID
    Store(id string, data []byte) // Stores the result of a task by its ID
}
