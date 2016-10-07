package main

import "github.com/bradylove/envstruct"

type HostInfo struct {
	IP       string `env:"HOST_IP,required"`
	Password string `env:"PASSWORD,noreport"`
	Port     int    `env:"HOST_PORT"`
}

func main() {
	hi := HostInfo{Port: 80}

	err := envstruct.Load(&hi)
	if err != nil {
		panic(err)
	}

	envstruct.WriteReport(&hi)
}
