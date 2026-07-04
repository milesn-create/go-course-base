# Go Demo 4 — Продвинутые темы

## Структура проекта
```
go-demo-4/
├── main.go                  — точка входа, меню
├── .env                     — переменные окружения (в .gitignore, не пушится!)
├── data.vault               — зашифрованное хранилище аккаунтов
├── account/
│   ├── account.go           — структура Account, конструктор, генерация пароля
│   ├── vault.go             — Vault, VaultWithDb, интерфейс Db, generics
│   └── output/
│       └── errors.go        — вывод ошибок через any
├── encrypter/
│   └── encrypter.go         — шифрование/дешифрование AES-GCM
├── files/
│   └── files.go             — работа с файлом (Read/Write)
└── cloud/
    └── cloud.go             — заглушка облачного хранилища
```

## Темы

### Структуры и ООП
- Структуры (`struct`), методы, конструкторы (`NewAccount`, `NewVault`)
- Композиция структур (embedding) — `VaultWithDb` встраивает `Vault`
- Struct tags (`json:"login"`) и сериализация JSON (`encoding/json`)

### Обработка ошибок
- Sentinel errors (`errors.New`, `errors.Is`)
- `panic()` при критических ошибках

### Интерфейсы и абстракция
- Интерфейсы (`Db`, `ByteReader`, `ByteWriter`)
- Объединение интерфейсов
- DI (Dependency Injection) — `NewVault(db Db, enc Encrypter)` принимает любую реализацию хранилища
- Тип `any`, `PrintError(value any)`

### Generics
- Generic функции (`sum[T int | string]`)
- Generic структуры (`List[T any]`)
- Хак `any(a).(type)` для type switch внутри generic

### Работа с файлами и ОС
- Чтение/запись файлов (`os.ReadFile`, `os.Create`)
- `defer file.Close()`
- Переменные окружения (`os.Getenv`, пакет `godotenv`)
- Файл `.env` для хранения секретов (добавлен в `.gitignore`)

### Шифрование
- Симметричное шифрование AES-GCM (`crypto/aes`, `crypto/cipher`)
- Генерация случайного nonce (`crypto/rand`)
- Шифрование (`Encrypt`) и дешифрование (`Decrypt`) данных перед записью в файл

### Функции как значения
- Map функций вместо `switch/case`
- Анонимные функции как аргументы (`FindAccounts` + checker)
- Вариативные аргументы (`promtData(prompt ...string)`)

### Внешние зависимости
- `github.com/fatih/color` — цветной вывод в терминал
- `github.com/joho/godotenv` — загрузка `.env` файла
- `go get`, `go.mod`, `go.sum`, `go mod tidy`

## Итоговый проект
Менеджер паролей — консольное приложение с **зашифрованным** хранилищем аккаунтов.

Возможности:
- создание аккаунта с валидацией URL и логина
- автогенерация пароля если не введён
- поиск по URL и по логину
- удаление аккаунта по URL
- данные хранятся в зашифрованном файле `data.vault` (AES-GCM)
- ключ шифрования передаётся через переменную окружения `KEY`
- абстракция хранилища через интерфейс `Db` (JSON файл или облако — взаимозаменяемы)

## Запуск
1. Создай файл `.env` в корне проекта:
```
KEY=твой_ключ_32_символа
```
2. Запусти:
```bash
go run .
```