package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/hypersleep/easyssh.v0"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Servers  []string `yaml:"Servers"`
	User     string   `yaml:"User"`
	Port     string   `yaml:"Port"`
	Key      string   `yaml:"Key"`
	Commands []string `yaml:"Commands"`
}

func main() {
	err := filepath.Walk("./configs", readConfing)
	if err != nil {
		log.Println(err)
	}
}

func readConfing(path string, info os.FileInfo, err error) error {
	if info != nil && !info.IsDir() {
		log.Println("Run command:", path)
		conf, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println("Read error:", err)
		}
		parseAndRun(conf)
	}
	return nil
}

func parseAndRun(conf []byte) {
	//Init with default values
	config := Config{
		Port: "22",
	}

	err := yaml.Unmarshal(conf, &config)
	if err != nil {
		log.Println("Config error:", err)
	}

	for _, server := range config.Servers {
		ssh := &easyssh.MakeConfig{
			User:   config.User,
			Server: server,
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
}
