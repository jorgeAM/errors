package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration -> configuration model
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfig -> get configuration to db's access
func GetConfig() (c Configuration) {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)
	return
}
