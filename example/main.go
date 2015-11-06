package main

import (
	"fmt"
	"strings"

	"github.com/baldmountain/seneca"
	"github.com/baldmountain/seneca/tcp"
	"github.com/baldmountain/seneca/web"
	"github.com/bitly/go-simplejson"
)

// Echo sends an echo to a seneca server
func Echo(r seneca.Acter, s string) (string, error) {
	req, _ := simplejson.NewFromReader(strings.NewReader(`{"role":"echo","cmd":"echo"}`))
	req.Set("msg", s)
	// actually call the remote service
	res, err := r.Act(req)
	if err != nil {
		return "", err
	}
	return res.Get("msg").String()
}

func main() {
	s, _ := Echo(&web.Requester{Host: "localhost", Port: 3030}, "hi using web")
	fmt.Println("echo:", s)

	r := tcp.Requester{Host: "localhost", Port: 3031}
	defer r.Close()
	s, _ = Echo(&r, "hi using tcp")
	fmt.Println("echo:", s)
}
