# Go Course Base

Учебные проекты по языку Go, написанные в процессе прохождения курса.

## Структура

### go-demo — Основы языка
- Первая программа, пакеты, сборка (`go build`, `go run`)
- Переменные (`:=`, `var`, `const`), типы данных
- Ввод/вывод (`fmt.Scan`, `fmt.Printf`)
- Условия (`switch`, `if/else`)
- Калькулятор ИМТ

### go-demo-2 — Массивы, слайсы, циклы
- Массивы и слайсы (`len`, `cap`, `append`, `make`)
- Циклы (`for`, `for range`)
- Указатели (`&`, `*`)
- Счётчик транзакций

### go-demo-3 — Map и структуры данных
- Map (создание, итерация, `delete`)
- Type aliases
- Менеджер закладок

### go-demo-4 — Продвинутые темы
- Структуры (`struct`), методы, конструкторы
- Композиция (embedding), DI (Dependency Injection)
- Обработка ошибок (`errors.New`, `errors.Is`)
- Интерфейсы, `any`, type switch, type assertion
- Generics
- Работа с файлами и JSON
- Шифрование AES-GCM (`crypto/aes`, `crypto/cipher`)
- Переменные окружения (`.env`, `godotenv`)
- Функции как значения, анонимные функции, замыкания
- **Итоговый проект:** Менеджер паролей с зашифрованным хранилищем

### go-demo-5 — HTTP запросы и тесты
- HTTP GET и POST запросы (`net/http`)
- Парсинг JSON ответов (`encoding/json`)
- Работа с URL и query параметрами (`net/url`)
- CLI флаги (`flag.String`, `flag.Int`, `flag.Parse`)
- Unit тесты (`testing`, `t.Error`, `t.Errorf`)
- Табличные тесты (test cases через слайс структур)
- **Итоговый проект:** CLI приложение — погода по городу

`Go` `Structs` `Interfaces` `Generics` `Closures` `JSON` `Error handling` `Encryption` `AES-GCM` `HTTP` `REST API` `Testing` `CLI`
