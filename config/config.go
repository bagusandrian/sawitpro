package config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func New(ctx context.Context) (*Config, error) {
	cfg := new(Config)
	err := initLocalConfig(ctx, cfg)
	if err != nil {
		return nil, err

	}
	return cfg, nil
}

func initLocalConfig(ctx context.Context, cfg *Config) error {

	localFile := getConfigFile()

	f, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}
func getConfigFile() string {
	var (
		env      = os.Getenv("APPS_ENV")
		filename = "config.yaml"
	)
	log.Printf("%s", env)
	if env != "" {
		return filepath.Join("/etc/sawitpro", filename)
	}

	return filepath.Join("../files/etc/development", filename)

}
