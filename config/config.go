// Package config represents database servers and credentials
// Parsing read & parse
package config

import (
	"log"

	"courses/BurntSushi/toml"
)

// Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
