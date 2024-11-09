package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	Token      `yaml:"token"`
	DB         DBConfig `yaml:"db"`
	HTTPServer `yaml:"server"`
}

type Token struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-default:"10m"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-default:"1m"`
}

type DBConfig struct {
	Host             string `yaml:"host" env-default:"localhost"`
	Port             string `yaml:"port" env-default:"5432"`
	User             string `yaml:"user" env-default:"postgres"`
	Password         string `yaml:"password" env-default:"admin"`
	DBName           string `yaml:"name" env-default:"auth_service"`
	SSLMode          string `yaml:"sslmode" env-default:"disable"`
	TimeZone         string `yaml:"timeZone" env-default:"europe/moscow"`
	MigrateTestData  bool   `yaml:"migrate_test_data"`
	UseORMMigrations bool   `yaml:"use_orm_migrations"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	path := "./config/local.yaml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
