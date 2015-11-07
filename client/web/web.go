package web

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

// Requester is a struct to wrap a host and port to send requests to
type Requester struct {
	Host string
	Port int
}

func request(host string, port int, json []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%d/act", host, port)
	res, err := http.Post(url, "text/json", bytes.NewReader(json))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	return b, err
}

// Close does nothing
func (r *Requester) Close() error {
	return nil
}

// Act on an interface to request and fills in a responce interface
func (r *Requester) Act(req *simplejson.Json) (*simplejson.Json, error) {
	jsonStr, err := req.Encode()
	if err != nil {
		return nil, err
	}
	json, err := request(r.Host, r.Port, jsonStr)
	if err != nil {
		return nil, err
	}
	// Unmarshal is smart enough to put lowercase response in uppercase reply
	return simplejson.NewJson(json)
}
