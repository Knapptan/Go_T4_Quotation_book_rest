package models

import "fmt"

// Основная модель цитаты
type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"quote"`
}

// Для запроса на создание циаты
type CreateQuoteRequest struct {
	Author string `json:"author"`
	Text   string `json:"quote"`
}

// Конфигурация
type Config struct {
	Port    string
	Address string
}

func (c *Config) StrAdress() string {
	return fmt.Sprintf("%s:%s", c.Address, c.Port)
}
