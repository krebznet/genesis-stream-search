package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DBStruct defines fields for databases
type DBStruct struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

// ServerStruct defines fields for main logic
type ServerStruct struct {
	Address string `yaml:"address"`
}

// Config structure for server
type Config struct {
	Server ServerStruct `yaml:"server"`
	DB     DBStruct     `yaml:"db"`
}

// ParseYamlFile the config file
func ParseYamlFile(filename string, c *Config) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	if user := os.Getenv("MONGO_USERNAME"); user != "" {
		c.DB.User = user
	}
	if password := os.Getenv("MONGO_PASSWORD"); password != "" {
		c.DB.Password = password
	}
	if host := os.Getenv("MONGO_HOST"); host != "" {
		c.DB.Host = host
	}
	if port := os.Getenv("MONGO_PORT"); port != "" {
		c.DB.Port = port
	}
	if database := os.Getenv("MONGO_DATABASE"); database != "" {
		c.DB.Name = database
	}

	return nil
}
