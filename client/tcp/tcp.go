package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/pborman/uuid"
)

// Requester is a struct to wrap a host and port to send requests to
type Requester struct {
	Host   string
	Port   int
	Conn   net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func (r *Requester) request(json []byte) ([]byte, error) {
	_, err := r.Writer.Write(json)
	if err != nil {
		return nil, err
	}
	err = r.Writer.WriteByte('\n')
	if err != nil {
		return nil, err
	}
	err = r.Writer.Flush()
	return r.Reader.ReadBytes('\n')
}

// Close closes the open tcp connection if it is open
func (r *Requester) Close() error {
	var err error
	if r.Conn != nil {
		err = r.Conn.Close()
		r.Conn = nil
		r.Reader = nil
		r.Writer = nil
	}
	return err
}

func (r *Requester) openConnection() error {
	var err error
	if r.Conn == nil {
		seconds := 5
		r.Conn, err = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", r.Host, r.Port), time.Duration(seconds)*time.Second)
		if err != nil {
			return err
		}
		r.Reader = bufio.NewReader(r.Conn)
		r.Writer = bufio.NewWriter(r.Conn)
	}
	return err
}

// Act on an interface to request and fills in a responce interface
func (r *Requester) Act(req *simplejson.Json) (*simplejson.Json, error) {
	if err := r.openConnection(); err != nil {
		return nil, err
	}

	json, err := simplejson.NewFromReader(strings.NewReader(`{"kind":"act","origin":"Go"}`))
	if err != nil {
		return nil, err
	}

	var id = uuid.New()
	json.Set("id", id)
	json.Set("time", time.Now().String())
	json.Set("act", req)
	s, err := json.Encode()
	if err != nil {
		return nil, err
	}

	out, err := r.request(s)
	if err != nil {
		return nil, err
	}

	json, err = simplejson.NewJson(out)
	if err != nil {
		return nil, err
	}

	if json.Get("id").MustString("") != id {
		return nil, errors.New("unexpected result")
	}
	return json.Get("res"), nil
}
