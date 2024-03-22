# Микросервис парсинга курса валют ЦБ РФ

Приложение, отдающее курс валюты по ЦБ РФ за определенную дату. Для получения курсов валют используется официальное API ЦБ РФ.

#### Получения курса валюты за определенную дату

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>info</code></summary>

##### Parameters

> | name     | type     | data type | example    | description                 |
> |----------|----------|-----------|------------|-----------------------------|
> | currency | required | string    | USD        | Валюта в стандарте ISO 4217 |
> | date     | optional | string    | 2016-01-06 | Дата в формате YYYY-MM-DD   |

##### Example output

```json 
{
    "data": {
      "USD": "33,4013"
    },
    "service": "currency"
}
```

</details>


## Структура проекта

```bash
.
├── .github         
│   └── workflows   # CI
├── cmd
│   └── main        # Точка входа в приложение
├── docs            # Проектная документация OpenApi
├── internal
│   ├── app         # Настройки приложения
│   ├── config      # Парсинг переменных окружения (стандартный порт)
│   ├── controller
│   │   └── http
│   │       └── v1  # Endpoints 
│   ├── entity      # Сущности
│   └── usecase     # Бизнес-логика приложения
│       └── cbrf    # Обработка данных с ЦБ РФ
└── pkg
    ├── httpserver  # Конфигурации для работы с HTTP сервером
    └── web         # Конфигурации для обработки JSON-ответов
```

## Документация и запуск

Для запуска выполнить сборку приложения

```bash
docker-compose up --build
```