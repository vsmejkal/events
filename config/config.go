package config
import (
	"os"
	"encoding/json"
	"fmt"
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

var Server struct {
	// Port
	Port int
	// Document root
	DocumentRoot string
}

func Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&struct{
		Database interface{}
		Facebook interface{}
		Server interface{}
	}{
		&Database,
		&Facebook,
		&Server,
	})
	if err != nil {
		return fmt.Errorf("Error parsing '%s': %s", path, err)
	}

	return nil
}
