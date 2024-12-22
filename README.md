# Go RPN Calculator
Веб-сервер для решения математических выражений, разработанный на языке программирования **Go**.
Для реализации была использована обратная польская нотация.

## Поддерживаемые операции

- Сложение `+`
- Вычитание `-`
- Умножение `*`
- Деление `/`
- Возведение в степень `^`
- Операнды приоритизации `(`, `)`
-

<img alt="Logotype" height="500" src="./.docs/logo.png" width="500"/>

<!--Установка-->
## Установка и первый запуск (Linux)
1. Клонирование репозитория

```bash
$ git clone https://github.com/Anti-Sh/go-rpn-calculator.git
```

2. Переход в директорию репозитория
```bash
$ cd go-rpn-calculator
```

3. Запуск (по умолчанию приложение использует порт 8080)
```bash
$ go run ./cmd/main.go
```

*Если есть необходимость изменить порт Веб-сервера, то запуск осуществлять этой командой:*
```bash
$ export PORT=9000 && go run ./cmd/main.go
```

<!--Конечные точки-->
## Конечные точки

### POST `/api/v1/calculate`
Используется для вычисления выражения. Телом запроса является JSON формата:

```json
{
  "expression": "1235/10"
}
```
Результатом также является объект JSON:
```json
{
  "result": 123.5
}
```

#### Примеры успешных запросов
1. Выражение `1235+500/(2500-100)`

```bash
$ curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "1235+500/(2500-100)"
}'
```
Ответ (Код 200):
```json
{
    "result": 1235.2083333333333
}
```

2. Выражение `10+5*8^(1/5)`

```bash
$ curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "10+5*8^(1/5)"
}'
```

Ответ (Код 200):
```json
{
    "result": 17.57858283255199
}
```

#### Примеры неуспешных запросов
1. Выражение `105-(10` (непарные скобки)

```bash
$ curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "105-(10"
}'
```
Ответ (Код 422):
```json
{
  "error": "Expression is not valid"
}
```

2. Выражение `11623/0` (деление на 0)

```bash
$ curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "11623/0"
}'
```
Ответ (Код 422):
```json
{
  "error": "Expression is not valid"
}
```

3. Выражение `152+10` (GET запрос)

```bash
$ curl --location --request GET  'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "152+10"
}'
```
Ответ (Код 405):
```json
{
  "error": "Invalid request method"
}
```

4. Выражение `152+a` (Неизвестные символы в выражении)

```bash
$ curl --location --request GET  'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "152+a"
}'
```
Ответ (Код 422):
```json
{
  "error": "Expression is not valid"
}
```