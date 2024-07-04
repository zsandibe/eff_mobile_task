# Effective Mobile Test Task


## Clone the project

```
$ git clone https://github.com/zsandibe/eff_mobile_task
$ cd eff_mobile_task
```

## Launch a project

```
$ make run
```

## Execute migrations

```
$ make migrate-up
$ make migrate-down
```

OR

```
$ make start
$ make stop
```

## SwaggerUI

```

localhost:8888/swagger/index.html

```


## API server provides the following endpoints:
* `GET /api/v1/users` - returns a users list by filter
* `POST /api/v1/users` - creates a user  by body(passport_serie,passport_number)
* `GET /api/v1/users/{id}` - returns a user by id from query path
* `PUT /api/v1/users/{id}` - updates a user personal data by id from path and body(any fields of structure User)
* `DELETE /api/v1/users/{id}` - deletes a user by id from path
* `POST /api/v1/tasks` - creates a task by body(user_id,name,description)
* `PUT /api/v1/tasks/{id}` - updates task`s data by task_id from path and body(user_id)


# .env file
## API configuration

```
API_URL=YOUR_TESTING_URL
API_KEY=YOUR_TESTING_KEY
```

## Server configuration

```
SERVER_HOST=localhost
SERVER_PORT=8888
```

## Postgres configuration

```
DRIVER=
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

# ATTENTION 
## Примечание по тестовому заданию
### Из-за того,что в тз не было url и key апишки,по сваггеру сделал мок респонс.Для проверки прошу добавить url и key апишки в env file.