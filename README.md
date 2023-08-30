# Сервис динамического сегментирования пользователей

Данный проект представляет собой сервис для управления сегментами пользователей и их участия в экспериментах на платформе Avito. Сервис предоставляет HTTP API для создания, изменения и управления сегментами пользователей, а также получения актуальных данных о сегментах, в которых пользователь состоит.

## Требования

- Golang
- Docker и Docker Compose
- Реляционная СУБД: MySQL или PostgreSQL

## Установка и запуск

1. Клонирование репозитория:

```bash
git clone https://github.com/Max425/avito-segment.git
cd avito-segment
```

2. Создание и запуск Docker-контейнеров:

```bash
docker-compose up --build
```

Сервис будет доступен по адресу: http://localhost:8000

## Использование API

Документация по API доступна по адресу: [http://localhost:8000/swagger/index.html#/](http://localhost:8000/swagger/index.html#/)

Там же можно проверить все методы.

## Подготовка пользователей для тестирования

По заданию не требуется добавление пользователей, но пользователи нужны для тестирования методов. Чтобы добавить пользователей, выполните следующие шаги:

1. Откройте терминал или командную строку.

2. Введите следующую команду, чтобы подключиться к контейнеру с базой данных:

```bash
docker exec -it avito-segment-db-1 psql -U postgres -d postgres
```

3. Внутри утилиты `psql` выполните SQL-запрос, чтобы добавить 5 пользователей:

```sql
INSERT INTO users (id) VALUES (1), (2), (3), (4), (5);
```

Теперь у вас должно быть 5 новых записей в таблице `users` вашей базы данных. Убедитесь, что контейнер с базой данных запущен и доступен, прежде чем выполнять эти шаги.

## Тестирование

Проект покрыт тестами для обеспечения стабильности и надежности. Вы можете запустить все тесты с помощью следующей команды:

```bash
go test -v ./...
```

Эта команда выполнит все тесты в проекте и выведет результаты в терминале.

## Реализованные дополнительные задания 

1. **История сегментов**: Реализован `GET` метод `/api/users/:user_id/segments/history")`, который позволяет получать историю изменений сегментов пользователя начиная с указанной даты. Этот метод возвращает информацию о действиях (добавление/удаление сегментов) и времени, когда эти действия были совершены.

2. **TTL для сегментов**: В процессе разработки была успешно выполнена дополнительная задача: реализация возможности добавления пользователя в эксперимент на ограниченный срок. Для этого было введено поле TTL (время жизни) при добавлении пользователя в сегмент. Теперь можно указывать, на сколько времени пользователь будет принадлежать к определенному сегменту. Например, можно выдать скидку пользователю на 2 дня.
```json
{
  "segment_slug": "discount_segment",
  "ttl_minutes": 2880
}
```
Данный метод также можно проверить в документации Swagger.

Триггер в PostgreSQL удаляет просроченные записи из таблицы `users_segments` при вставке новых записей. Для проверки доп. задания следует выполнить следующие шаги:

1. Выполнить POST запрос по `/api/users/{user_id}/segments/add_with_ttl` (в swagger), установив поле `ttl_minutes` 1 (минута).

2. Дождитесь, пока пройдет более 1 минуты.

3. Выполните какой-либо запрос на вставку новой записи в таблицу `users_segments`.

4. Проверьте содержимое таблицы `users_segments`. Вы увидите, что запись с истекшим временем `expires_at` будет автоматически удалена триггером.