package config

import (
	"fmt"
	"os"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
	"github.com/joho/godotenv"
)

func Load() (*models.Config, error) {

	if err := godotenv.Load("ini.env"); err != nil {
		return nil, fmt.Errorf("error loading ini.env: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("environment variable PORT is not set")
	}

	address := os.Getenv("ADDRESS")
	if address == "" {
		return nil, fmt.Errorf("address variable ADDRESS is not set")
	}

	cfg := &models.Config{
		Port:    port,
		Address: address,
	}

	return cfg, nil
}
