package client

import "io"

// Acter is a wrapper around the Act method
type Acter interface {
	Act(req interface{}, res interface{}) error
	io.Closer
}
