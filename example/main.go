package main

import (
	"fmt"

	"github.com/baldmountain/seneca/client"
	"github.com/baldmountain/seneca/client/tcp"
	"github.com/baldmountain/seneca/client/web"
)

type echo struct {
	Role string `json:"role"`
	Cmd  string `json:"cmd"`
	Msg  string `json:"msg"`
}

// Echo sends an echo to a seneca server
func Echo(r client.Acter, s string) (string, error) {
	req := struct {
		Role string `json:"role"`
		Cmd  string `json:"cmd"`
		Msg  string `json:"msg"`
	}{"echo", "echo", s}
	var res = &echo{}
	// actually call the remote service
	err := r.Act(req, res)
	if err != nil {
		return "", err
	}
	return res.Msg, nil
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
