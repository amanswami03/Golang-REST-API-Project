package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configpath string

	configpath = os.Getenv("CONFIG_PATH")

	if configpath == "" {
		flags := flag.String("config", "", "path to the confihuration file")
		flag.Parse()
		configpath = *flags

		if configpath == "" {
			log.Fatal("Config file not found")
		}
	}
	var cfg Config
	// if _, err := os.Stat(configPath); os.IsNotExist(err) {
	// 	log.Fatalf("config file does not exist: %s", configPath)
	// }
	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
