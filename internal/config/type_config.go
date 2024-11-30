package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	home_file, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	config_file := home_file + "/.gatorconfig.json"

	file, err := os.Create(config_file)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	home_file, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	config_file := home_file + "/.gatorconfig.json"

	file, err := os.Open(config_file)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
