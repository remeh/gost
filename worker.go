package gost

// A worker receives tasks from a broadcaster
// and compute them, then the response is written
// in the storage system to be available when
// dest clients come for data.
type Worker interface {
    Run(*Task) []byte    // Executes the supplied task given by the broadcaster
}
