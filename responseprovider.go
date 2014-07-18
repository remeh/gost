package gooch

// A Response Provider is the part which knows
// where and how to look for responses.
//
// The responses could have been stored into DB, files, 
// RAM or wherever the storage system is configured
// to write/read.
type ResponseProvider interface {
    GetResponse(id []byte) (*Response, error) // Returns the found response with the given ID.
}
