# Цитатник

REST API для управления цитатами.

## Запуск

```bash
go run cmd/server/main.go
```

Техническое описание:

- Данные хранятся в памяти, но есть интерфейс для добавки других бд.
- Только стандартные библиотеки Go и gorilla/mux

## Примеры запросов:

Добавление новой цитаты (POST /quotes)

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author": "Confucius", "quote": "Life is simple, but we insist on making it complicated."}'
```

Получение всех цитат (GET /quotes)

```bash
curl http://localhost:8080/quotes
```

Получение случайной цитаты (GET /quotes/random)

```bash
curl http://localhost:8080/quotes/random
```

Фильтрация по автору (GET /quotes?author=Confucius)

```bash
curl http://localhost:8080/quotes?author=Confucius
```

Удаление цитаты по ID (DELETE /quotes/{id})

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

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
├── makefile
└── tests                    # Юнит-тесты
    ├── handler
    │   └── quotes_test.go
    └── storage
        └── inmemory_test.go
```
