package main

import (
	"encoding/json"
	"flag"
	"github.com/FabelVlad/golang_rest_api/internal/app/apiserver"
	"io/ioutil"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "set path to config file")
}

func loadConfig(filename string) (*apiserver.Config, error) {
	var config *apiserver.Config
	configContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configContent, &config)
	return config, err
}

func main() {
	flag.Parse()

	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
