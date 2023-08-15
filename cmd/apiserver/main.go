package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"http-rest-api21/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		panic(err)
	}
}
