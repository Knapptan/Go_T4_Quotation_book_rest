package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Knapptan/Go_T1_Name_info_rest/internal/models"
)

func Load() (*models.Config, error) {

	if err := loadEnvFile("ini.env"); err != nil {
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

func loadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}
