# Цитатник

REST API для управления цитатами.

## Запуск
```bash
go run cmd/server/main.go
```

Примеры запросов:
Добавить цитату

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author": "Confucius", "quote": "Life is simple, but we insist on making it complicated."}'
```
Получить все цитаты

```bash 
curl http://localhost:8080/quotes
```

Добавление новой цитаты (POST /quotes) 
Получение всех цитат (GET /quotes)  
Получение случайной цитаты (GET /quotes/random)  
Фильтрация по автору (GET /quotes?author=Confucius)  
Удаление цитаты по ID (DELETE /quotes/{id})

```
quotes-service/
├── cmd/
│   └── server/
│       └── main.go          # Точка входа
├── internal/
│   ├── handler/             # HTTP-обработчики
│   │   └── quotes.go
│   ├── storage/             # Логика хранения данных
│   │   └── inmemory.go
│   └── models/              # Модели данных
│       └── quote.go
├── go.mod
├── go.sum
├── README.md                # Инструкции
└── tests/        # Юнит-тесты
```