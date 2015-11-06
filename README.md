## seneca

** A client for accessing a nodejs Seneca service from Go. **

### Installation

```bash
go get github.com/baldmountain/seneca
```

### Usage

The client relies heavily on github.com/bitly/go-simplejson. If you aren't
familiar with that you may want to look there first. simplejson is used to
build requests that will be sent to the server. Seneca uses json for all
it's communications so the client needs to generate json.

The easiest way to build a request is like:

```Go
req, _ := simplejson.NewFromReader(strings.NewReader(`{"role":"echo","cmd":"echo"}`))
req.Set("msg", s)
```

Then to actually call the service just create either a web.Requester or a
tcp.Requester and call Act on it.

```Go
r := web.Requester{Host: "localhost", Port: 3030}
// call the remote service
res, err := r.Act(req)
if err != nil {
  return "", err
}
res.Get("msg").String()
```

The returned value from Act is a simplejson object that can be queried for values.
See the main.go file in the example directory for a more complete example. There
is also a simple echo server in the example directory that can be used to try out
the example.

### License

Copyright (c) 2015 Geoffrey Clements. - All rights reserved.
