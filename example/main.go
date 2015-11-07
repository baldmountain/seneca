package main

import (
	"fmt"
	"strings"

	"github.com/baldmountain/seneca/client"
	"github.com/baldmountain/seneca/client/tcp"
	"github.com/baldmountain/seneca/client/web"
	"github.com/bitly/go-simplejson"
)

// Echo sends an echo to a seneca server
func Echo(r client.Acter, s string) (string, error) {
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
	s, err := Echo(&web.Requester{Host: "localhost", Port: 3030}, "hi using web")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("echo:", s)
	}

	r := tcp.Requester{Host: "localhost", Port: 3031}
	defer r.Close()
	s, err = Echo(&r, "hi using tcp")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("echo:", s)
	}
}
