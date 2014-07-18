package gooch

// Response State error codes
// Based on HTTP error codes for wider comprehension.
const (
    R_STATE_NOT_COMPUTED = 0
    R_STATE_OK           = 200
    R_STATE_ERROR        = 500
)

// A computed response, delivered in the storage system
// by the workers.
type Response interface {
    GetData() []byte // The actual data of the response
    GetState() int   // The state of the response
}

// Simple container of a response.
type SimpleResponse struct {
    data  []byte
    state int
}

func (r *SimpleResponse) GetData() []byte {
    return r.data
}

func (r *SimpleResponse) GetState() int {
    return r.state
}
