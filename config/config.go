package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RPCURL         string
	ContractAddress string
	PrivateKey     string
	OwnerAddress   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RPCURL:         os.Getenv("RPC_URL"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		PrivateKey:     os.Getenv("PRIVATE_KEY"),
		OwnerAddress:   os.Getenv("OWNER_ADDRESS"),
	}
}
