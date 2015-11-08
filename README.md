## seneca

**A client for accessing a nodejs Seneca service from Go.**

### Installation

```bash
go get github.com/baldmountain/seneca
```

### Usage

The Act method takes any interface. You'll need to define a struct that represents
your data for it to be Marshaled and Unmarshaled from JSON correctly.

Start by defining a struct like:

```Go
type echo struct {
	Role string `json:"role"`
	Cmd  string `json:"cmd"`
	Msg  string `json:"msg"`
}
```

You'll need a response struct to capture the reply from the service.

To actually call the service just create either a web.Requester or a
tcp.Requester and call Act on it and pass both the request and response.

```Go
r := web.Requester{Host: "localhost", Port: 3030}
req := echo{Role: "echo", Cmd: "echo", Msg: "Hello, world!"}
res := &echo{}
// call the remote service
err := r.Act(req, res)
if err != nil {
  fmt.Println("error", err)
} else {
  fmt.Println(res.Msg)
}
```

In this case the response is the same as the request so we can pass an instance
of echo. Most likely you'll be creating a different response structure to match
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
