# Go Demo 5 — HTTP запросы и тесты

## Структура проекта
```
go-demo-5/
├── main.go              — точка входа, CLI флаги
├── geo/
│   ├── geo.go           — определение города по IP и валидация города
│   └── geo_test.go      — тесты для geo пакета
└── weather/
    ├── weather.go       — получение погоды через wttr.in
    └── weather_test.go  — тесты для weather пакета
```

## Темы

### HTTP запросы
- GET запрос (`http.Get`) — определение города по IP через ip-api.com
- POST запрос (`http.Post`) — проверка существования города через countriesnow.space
- Чтение тела ответа (`io.ReadAll`)
- Закрытие тела ответа (`defer resp.Body.Close()`)
- Проверка статус кода (`resp.StatusCode`)

### Работа с URL
- Парсинг URL (`url.Parse`)
- Query параметры (`url.Values`, `params.Add`, `baseUrl.RawQuery`)

### CLI флаги
- `flag.String` — флаг строкового типа
- `flag.Int` — флаг числового типа
- `flag.Parse` — парсинг флагов
- Запуск: `go run . -city=London -format=3`

### Тесты
- Unit тесты (`testing`, `t.Error`, `t.Errorf`)
- Табличные тесты (test cases через слайс структур)
- Запуск тестов: `go test ./...`

## Итоговый проект
CLI приложение — погода по городу.

Возможности:
- определяет город автоматически по IP если не указан
- валидирует введённый город через внешний API
- получает погоду через wttr.in в одном из 4 форматов
- выводит результат в терминал

## Запуск
```bash
# погода по текущему местоположению
go run .

# погода по конкретному городу
go run . -city=London

# погода в определённом формате (1-4)
go run . -city=London -format=3
```

## Тесты
```bash
go test ./...
```
