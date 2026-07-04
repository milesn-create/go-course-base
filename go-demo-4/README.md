# Go Demo 4 — Продвинутые темы

## Структура проекта
```
go-demo-4/
├── main.go              — точка входа, меню
├── account/
│   ├── account.go       — структура Account, конструктор, генерация пароля
│   ├── vault.go         — Vault, VaultWithDb, интерфейс Db, generics
│   └── output/
│       └── errors.go    — вывод ошибок через any
├── files/
│   └── files.go         — работа с JSON файлом (Read/Write)
└── cloud/
    └── cloud.go         — заглушка облачного хранилища
```

## Темы
- Структуры (`struct`), методы, конструкторы (`NewAccount`)
- Композиция структур (embedding) — `VaultWithDb` встраивает `Vault`
- Struct tags (`json:"login"`) и сериализация JSON (`encoding/json`)
- Обработка ошибок (`errors.New`, `errors.Is`, sentinel errors)
- Интерфейсы (`Db`, `ByteReader`, `ByteWriter`), объединение интерфейсов
- Тип `any`, `PrintError(value any)`
- Generics (`sum[T int | string]`, `List[T any]`)
- Работа с файлами (`os.Create`, `os.ReadFile`, `file.Write`)
- `defer file.Close()`
- Сторонние пакеты (`github.com/fatih/color`, `go.mod`, `go.sum`)
- Map функций вместо `switch/case`
- Анонимные функции как аргументы (`FindAccounts` + checker)
- Вариативные аргументы (`promtData(prompt ...string)`)

## Итоговый проект
Менеджер паролей — консольное приложение с хранением аккаунтов
в JSON файле. Поддерживает:
- создание аккаунта с валидацией URL и логина
- автогенерацию пароля если не введён
- поиск по URL и по логину
- удаление аккаунта по URL
- абстракцию хранилища через интерфейс (JSON файл или облако)