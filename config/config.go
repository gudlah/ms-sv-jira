package config

import (
	"ms-sv-jira/models/dto"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var Config dto.ServerConfig

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)
	return err
}
