Pet проект на Golang.

Сервис авторизации по JWT-токену. Авторизация в разных приложениях происходит по различным токенам, генерируемым для каждого приложения.

Для запуска приложения перейдите в корень проекта и выполните команду docker-compose up. При первом запуске приложение может запуститься раньше БД и упасть с ошибкой, при повторных запусках ошибка не воспроизводилась. В дальнейшем планируется решить эту проблему и использовать Swagger для документирования. После сборки и запуска приложение будет доступно по адресу http://localhost:8080.

Краткое описание сервиса.

Регистрация пользователя
POST localhost:8080/register
Ожидает JSON вида
{
    "email": "test@test.test",
    "password": "password",
    "password-again": "password"
}

Авторизация пользователя
POST localhost:8080/login
Ожидает JSON вида
{
    "email": "test@test.test",
    "password": "password"
}
Возвращает API ключ для некоторых операций с токенами

Отзыв API ключа
DELETE localhost:8080/revoke-api-key
{
    "api-key": "0c288aded099000fc1669f82c504c4d07fbd34cb165500a75cc0c3d2c80d39b5"
}

Генерация access token
Токены разделены по приложениям, поэтому для генерации требуется код приложения, для которого токен создается.
POST localhost:8080/generate-access-token
Ожидает JSON вида
{
    "app-code": "test"
}
Требует в заголовках запроса Bearer с API ключом с шага авторизации пользователя
Возвращает сгенерированный access token.

Генерация refresh token
Токены разделены по приложениям, поэтому для генерации требуется код приложения, для которого токен создается.
POST localhost:8080/generate-refresh-token
Ожидает JSON вида
{
    "app-code": "test"
}
Требует в заголовках запроса Bearer с API ключом с шага авторизации пользователя
Возвращает сгенерированный refresh token.

Отзыв refresh token
DELETE localhost:8080/revoke-refresh-token
Ожидает JSON вида
{
    "refresh-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOiJ0ZXN0IiwiZXhwIjoxNzMwMTMxNTc3LCJ1aWQiOjU0fQ.Hco-iUvyR5Zh1EMaHIQKUxlw99NsguvB7UdlsMMoBlY"
}
Требует в заголовках запроса Bearer с API ключом с шага авторизации пользователя

Получение нового access token по переданному refresh token
Get localhost:8080/refresh-access-token
Требует в заголовках запроса Bearer с refresh token
