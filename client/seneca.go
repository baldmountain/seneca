package client

import "github.com/bitly/go-simplejson"

// Acter is a wrapper around the Act method
type Acter interface {
	Act(req *simplejson.Json) (*simplejson.Json, error)
	Close() error
}
