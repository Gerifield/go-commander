package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/hypersleep/easyssh.v0"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   string   `yaml:"Server"`
	User     string   `yaml:"User"`
	Port     string   `yaml:"Port"`
	Key      string   `yaml:"Key"`
	Commands []string `yaml:"Commands"`
}

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println("Read error:", err)
	}

	//Init with default values
	config := Config{
		Port: "22",
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Println("Config error:", err)
	}

	ssh := &easyssh.MakeConfig{
		User:   config.User,
		Server: config.Server,
		Key:    config.Key,
		Port:   config.Port,
	}

	for _, cmd := range config.Commands {
		resp, err := ssh.Run(cmd)
		if err != nil {
			log.Println("Command error", cmd, err)
		}
		log.Println(resp)
	}
}
