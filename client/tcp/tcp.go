package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

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

type actCommon struct {
	Kind   string `json:"kind"`
	Origin string `json:"origin"`
	ID     string `json:"id"`
}

type actRequest struct {
	actCommon
	Time time.Time   `json:"time"`
	Act  interface{} `json:"act"`
}

type actResponse struct {
	actCommon
	Res interface{}
}

// Act on an interface to request and fills in a responce interface
func (r *Requester) Act(req interface{}, res interface{}) error {
	if err := r.openConnection(); err != nil {
		return err
	}

	var id = uuid.New()
	fullReq := actRequest{
		actCommon: actCommon{Kind: "act", Origin: "Go", ID: id},
		Time:      time.Now(),
		Act:       req}
	s, err := json.Marshal(fullReq)
	if err != nil {
		return err
	}

	out, err := r.request(s)
	if err != nil {
		return err
	}

	fullRes := &actResponse{Res: res}
	return json.Unmarshal(out, fullRes)
}
