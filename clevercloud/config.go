package clevercloud

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

var path = "/.config/clever-cloud"

// Config The struct matching config file
type Config struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

// GetConfigFromUser Read clever-cloud user config file
func GetConfigFromUser() *Config {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(usr.HomeDir + path)
	if err != nil {
		panic(err)
	}

	config := &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		panic(err)
	}

	return config
}

// GetConfigFromEnv Read environment for clever-cloud variables
func GetConfigFromEnv() *Config {
	return &Config{
		Token:  os.Getenv("CLEVER_TOKEN"),
		Secret: os.Getenv("CLEVER_SECRET"),
	}
}
