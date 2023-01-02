package config

import (
	_ "embed"
	"encoding/json"
)

type Env struct {
	DBHost       string `json:"DB_HOST"`
	DBPort       string `json:"DB_PORT"`
	DBUsername   string `json:"DB_USERNAME"`
	DBPassword   string `json:"DB_PASSWORD"`
	DatabaseName string `json:"DATABASE_NAME"`
	RedisHost string `json:"REDIS_HOST"`
	RedisPort string `json:"REDIS_PORT"`
}

var (
	env Env
	//go:embed env/development.json
	s string
)

func init() {
	loadEnv()
}

func loadEnv() {
	json.Unmarshal([]byte(s), &env)
}

func GetEnv() Env {
	return env
}
