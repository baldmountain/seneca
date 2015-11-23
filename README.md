## seneca

**A client for accessing a nodejs Seneca service from Go.**

### Installation

```bash
go get github.com/pborman/uuid
go get github.com/baldmountain/seneca
```
Ignore the message about no buildable Go source files. We are just getting the repository.

### Usage

To use the Seneca client you just create a Requester and call Act on it. There are
Requesters for the tcp and web transport. All the Requesters implement the Acter
interface.

```Go
type Acter interface {
	Act(req interface{}, res interface{}) ([]byte, error)
	io.Closer
}
```

The `io.Closer` interface specifies a Close method for the Requester that
is used to close any open network connections the Requester may have opened.

The Act method takes two parameters. The first is a struct specifying the request.

It should look something like:

```Go
type echo struct {
	Role string `json:"role"`
	Cmd  string `json:"cmd"`
	Msg  string `json:"msg"`
}
```

where the fields specify the pattern of the seneca request and any parameters
sent to the Seneca action. Note the `json:"role"` for the Role field. This allows
us to specify a lowercase json name since the Go JSON library only encodes exported
names. (The ones that start with a capital letter.)

The second parameter to Act is a response struct to capture the reply from the
service. It should include any fields that you want to decode from the response.
Any extra fields in the response, that are not specified in the response struct,
are discarded.

The Act method returns the JSON byte array for the response struct and an error if
there was one.

To actually call the service just create either a web.Requester or a
tcp.Requester and call Act on it. Pass both the request and response structs.

```Go
r := web.Requester{Host: "localhost", Port: 3030}
req := echo{Role: "echo", Cmd: "echo", Msg: "Hello, world!"}
res := &echo{}
// call the remote service
_, err := r.Act(req, res)
if err != nil {
  fmt.Println("error", err)
} else {
  fmt.Println(res.Msg)
}
```

In this case the response is the same as the request so we can pass an instance
of echo. Most likely you'll create a different response structure to match
the response from the service.

See the example folder for an example Go program that calls a Nodejs Seneca echo
service. In order to run the service you'll need a recent version of Nodejs.

*NOTE: This library probably isn't thread safe. If you make requests from multiple
Goroutines using the same requester you will most likely get unexpected results.*

### License

The MIT License (MIT)

Copyright (c) 2015 Geoffrey Clements

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
