package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// App config
type App struct {
	Address string
	Static  string
	Log     string
	Version string
}

// Database config
type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

// Configuration all config struct
type Configuration struct {
	App App
	Db  Database
}

var config *Configuration
var once sync.Once

// LoadConfig loads server config
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}

		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})

	return config
}
