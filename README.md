# envstruct
========================

envstruct is a simple library for populating values on structs from environment
variables.

## Documentation

[https://godoc.org/github.com/bradylove/envstruct](https://godoc.org/github.com/bradylove/envstruct)

## Usage

Export some environment variables.

```
$ export HOST_IP="127.0.0.1"
$ export HOST_PORT="443"
```

Write some code. In this example, `Ip` requires that the `HOST_IP` environment variable is set to non empty value and `Port` defaults to `80` if `HOST_PORT` is a non empty value.

```
package main

import (
    "fmt"
    "github.com/bradylove/envstruct"
)

type HostInfo struct {
    Ip   string `env:"host_ip,required"`
    Port int    `env:"host_port"`
}

func main() {
    hi := HostInfo{Port: 80}
    err := envstruct.Load(&hi)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Host: %s, Port: %d\n", hi.Ip, hi.Port)
}
```

Run your code and rejoice!

```
$ go run example/example.go
Host: 127.0.0.1, Port: 443
```

## Supported Types

- [x] string
- [x] bool (`true` and `1` results in true value, anything else results in false value)
- [x] int
- [x] int8
- [x] int16
- [x] int32
- [x] int64
- [x] uint
- [x] uint8
- [x] uint16
- [x] uint32
- [x] uint64
- [ ] float32
- [ ] float64
- [ ] complex64
- [ ] complex128
- [x] []slice (Slices of any other supported type. Environment variable should have coma separated values)
- [ ] time.Duration

## Running Tests

Run tests using ginkgo.

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
$ ginkgo
```

### MIT License

Copyright (c) 2016 Brady Love <love.brady@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
