package client

import "io"

// Acter is a wrapper around the Act method
/* params:

req - An outgoing request in the form of a struct that will be encoded to json.
res - A response structure that will be decoded from the returned stream.

returns:

The returned byte array from the seneca service. Since it is JSON it can
be passed as a response to a web client.

An error if there was one.

*/
type Acter interface {
	Act(req interface{}, res interface{}) ([]byte, error)
	io.Closer
}
