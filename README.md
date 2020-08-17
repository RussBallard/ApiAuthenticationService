# ApiAuthenticationService

По заданию необходимо написать сервис аутентификации.
Сервис должен иметь четыре REST маршрута:

• Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса;

• Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов;

• Третий маршрут удаляет конкретный Refresh токен из базы;

• Четвертый маршрут удаляет все Refresh токены из базы для конкретного пользователя;

Технологии:
Язык программирования Go. База данных MongoDB, топология Replica Set с использованием транзакций. Access токен тип JWT, алгоритм SHA512.
Refresh токен тип произвольный, формат передачи base64, хранится в базе исключительно в виде bcrypt хеша, должен быть защищен от изменения на стороне клиента и попыток повторного использования.
Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.

По завершению задания, проект можно запустить и протестировать выполнив команду docker-compose up, находясь внутри главной папки проекта.
