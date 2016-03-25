package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Database struct {
	// PostgreSQL database name
	Name string
}

var Facebook struct {
	// Access token (app_id|app_secret)
	Token string
	// Graph API URL
	GraphURL string
}

var Web struct {
	// Port
	Port int
	// Document root
	DocumentRoot string
}

var Admin struct {
	// Port
	Port int
	// Document root
	DocumentRoot string
	// Admin name
	User string
	// Admin password
	Password string
}

func Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&struct {
		Database interface{}
		Facebook interface{}
		Web interface{}
		Admin interface{}
	}{
		&Database,
		&Facebook,
		&Web,
		&Admin,
	})
	if err != nil {
		return fmt.Errorf("Error parsing '%s': %s", path, err)
	}

	return nil
}
