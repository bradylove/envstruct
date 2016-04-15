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
