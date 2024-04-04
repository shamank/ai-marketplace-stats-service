# Statistic Service (AI Marketplace)

***тестовое задание***

Микросервис, хранящий и обновляющий в базе данных информацию о количестве вызовов сервисов

### Использованные библиотеки:
- **pgx** - для работы с PostgreSQL
- **migrate** - для работы с миграциями
- **cleanenv** - для работы с конфигом
- **grpc** - для работы с gRPC (неожиданно)

### Структура базы данных

**Таблица *users***

| Column   | Type              | Nullable | Default           |
|----------|-------------------|----------|-------------------|
| uid      | uuid              | not null | gen_random_uuid() |
| username | character varying | not null |                   |


**Таблица *aiservices***

| Column      | Type              | Nullable | Default           |
|-------------|-------------------|----------|-------------------|
| uid         | uuid              | not null | gen_random_uuid() |
| title       | character varying | not null |                   |
| description | text              |          |                   |
| price       | double precision  | not null |                   |

**Таблица *statistics***

| Column        | Type                        | Nullable | Default           |
|---------------|-----------------------------|----------|-------------------|
| uid           | uuid                        | not null | gen_random_uuid() |
| user_uid      | uuid                        | not null |                   |
| aiservice_uid | uuid                        | not null |                   |
| amount        | double precision            | not null |                   | 
| created_at    | timestamp without time zone | not null | now()             |

*Примечание:*
В таблицу *statistics* при каждом новом вызове добавляется новая запись (вместо использования обычного поля-счетчика).
Поступил так из расчета того, что цена за использование сервиса может меняться.


### Запуск микросервиса

**Для старта микросервиса**:

```sh
make run
```
или
```sh
go run ./cmd/app/main.go --cfg=./configs/prod.yaml
```


**Для запуска миграций**:
```sh
make migrate
```
или
```sh
go run ./cmd/migrate/main.go --cfg=./configs/prod.yaml --migrations=./migrations/
```

### Структура проекта

```

├───cmd
│   ├───app // запуск основного приложения
│   └───migrate // запуск миграция
│
├───configs // здесь лежат различные конфиги
│
├───internal
│   ├───app // структура приложения
│   ├───config // парсинг конфига
│   ├───delivery
│   │   └───grpc // реализация grpc-абстракций 
│   │       └───statistic 
│   ├───domain
│   │   ├───errors
│   │   └───models
│   ├───logger
│   ├───repository 
│   │   └───postgres // реализация repository'я для postgresql
│   └───service
│       └───statistic // сервис для работы со статистикой
│
└───migrations // директория с .sql-файлами для мигнрации
```

