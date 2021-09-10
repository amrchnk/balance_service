# Тестовое задание на позицию стажера-бекендера

## Задача

Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.

[Подробнее...](https://github.com/avito-tech/autumn-2021-intern-assignment)

## Описание API-методов

***Метод начисления/списания средств***

Принимает один из следующих типов операций:
- increase (начисление средств)
- decrease (списание средств)

```
http://localhost:8000/balance/:type [POST]
```

**Начисление средств**

Запрос:

```http://localhost:8000/balance/increase```

Тело запроса:

```
{
    "user_id": 1,
    "balance": 400.45
}
```

Ответ (при успешном создании счета и зачислении средств):
```
{
    "id": 1,
    "message": "Current balance in rubles: 400.45",
    "status": 200
}
```

**Списание средств**

Запрос:

```http://localhost:8000/balance/decrease```

Тело запроса:

```
{
    "user_id": 1,
    "balance": 130
}
```

Ответ (при успешном списании средств со счета):
```
{
    "id": 1,
    "message": "Current balance in rubles: 270.45",
    "status": 200
}
```

***Метод получения текущего баланса пользователя***

```
http://localhost:8000/balance/:id [GET]
```

Запрос:
```
localhost:8000/balance/1
```

Ответ:
```
{
    "user_id": 1,
    "currency": "RUB",
    "balance": 270.45
}
```

При указании доп. параметра currency

Запрос:
```
localhost:8000/balance/1?currency=USD
```
Ответ:
```
{
    "user_id": 1,
    "currency": "USD",
    "balance": 3.7
}
```

***Метод перевода средств от пользователя к пользователю***

Запрос:
```
localhost:8000/balance/transfer [GET]
```
Тело запроса:
```
{
    "sender_id":1,
    "receiver_id":2,
    "sum":70
}
```
Ответ (пользователь с id=2 был создан после получения тела запроса):
```
{
    "sender_id": 1,
    "sender_balance": 200.45,
    "receiver_id": 2,
    "receiver_balance": 70
}
```
***Методы получения списка транзакций***

Получение списка всех транзакций
```
localhost:8000/transactions/ [GET]
```
Запрос:
```
localhost:8000/transactions/?sort=created&page=1&records=10&direction=down
```
Принимает следующие параметры в адресной строке:
- direction – направление сортировки (по возрастанию/убыванию, принимает значения up (по умолчанию)/down)
- page – номер страницы
- records – количество записей на странице
- sort – тип сортировки (amount – по балансу (стоит по умолчанию), created – по дате).

Ответ:
```
[
    {
        "user_id": 1,
        "type": "outgoing transfer",
        "amount": 70,
        "description": "money transfer to user with id=2",
        "created": "2021-09-06T23:29:36.978176Z"
    },
    {
        "user_id": 2,
        "type": "incoming transfer",
        "amount": 70,
        "description": "money transfer from user with id=1",
        "created": "2021-09-06T23:29:36.976598Z"
    },
    {
        "user_id": 1,
        "type": "increase in balance",
        "amount": 70,
        "description": "",
        "created": "2021-09-06T23:28:37.715262Z"
    },
    {
        "user_id": 1,
        "type": "increase in balance",
        "amount": 70,
        "description": "",
        "created": "2021-09-06T23:27:58.153518Z"
    },
    {
        "user_id": 1,
        "type": "outgoing transfer",
        "amount": 70,
        "description": "money transfer to user with id=4",
        "created": "2021-09-06T23:27:39.233541Z"
    }
]
```
Получение списка транзакций одного пользователя

```
localhost:8000/transactions/:id [GET]
```
Запрос:
```
localhost:8000/transactions/1?sort=created&page=1&records=5&direction=down
```

Ответ:
```
[
    {
        "user_id": 1,
        "type": "outgoing transfer",
        "amount": 70,
        "description": "money transfer to user with id=2",
        "created": "2021-09-06T23:29:36.978176Z"
    },
    {
        "user_id": 1,
        "type": "increase in balance",
        "amount": 70,
        "description": "",
        "created": "2021-09-06T23:28:37.715262Z"
    },
    {
        "user_id": 1,
        "type": "increase in balance",
        "amount": 70,
        "description": "",
        "created": "2021-09-06T23:27:58.153518Z"
    },
    {
        "user_id": 1,
        "type": "outgoing transfer",
        "amount": 70,
        "description": "money transfer to user with id=4",
        "created": "2021-09-06T23:27:39.233541Z"
    },
    {
        "user_id": 1,
        "type": "outgoing transfer",
        "amount": 70,
        "description": "money transfer to user with id=3",
        "created": "2021-09-06T23:27:25.889813Z"
    }
]
```
**Примеры ответов при ошибочных запросах**

- Метод списания/начисления 

Ответ (при попытке снятия средств сверх остатка):
```
{
    "message": "pq: новая строка в отношении \"balance\" нарушает ограничение-проверку 
    \"balance_balance_check\"",
    "status": 500
}
```

- Метод получения баланса пользователя при некорректном названии валюты

Запрос:
```
localhost:8000/balance/1?currency=US
```
Ответ:
```
{
    "message": "Incorrect currency",
    "status": 500
}
```

## Запуск приложения

Для запуска приложения необходимо выполнить последовательно следующие команды в корневой папке проекта:
```
docker-compose up --build app
```

```
docker-compose up app
```
Запуск тестов в директории handler (с выводом процента покрытия, в моем случае получилось покрыть 85% модуля):
```
go test -v -cover
```
