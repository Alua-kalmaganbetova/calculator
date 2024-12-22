Описание
Этот проект для вычисления арифметических выражений. Сервис принимает выражение в виде строки, рассчитывает его и возвращает результат.

Установка и запуск
Клонируйте репозиторий:

bash
Копировать код
git clone https://github.com/your-username/arithmetic-api.git
Перейдите в папку с проектом:

bash
Копировать код
cd arithmetic-api
Запустите сервис:

bash
Копировать код
go run ./cmd/calc_service/...
Сервис будет доступен по адресу: http://localhost:8080.

Пример использования
Запрос с арифметическим выражением
bash
Копировать код
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:

json
Копировать код
{
  "result": "6.000000"
}
Ошибка при неверном выражении
bash
Копировать код
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2a"
}'
Ответ:

json
Копировать код
{
  "error": "Expression is not valid"
}
Ошибки
422 — Expression is not valid: Неверное выражение (например, символы, кроме цифр и разрешённых операций).
500 — Internal server error: Внутренняя ошибка сервера (например, деление на ноль).
