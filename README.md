![License](https://img.shields.io/github/license/p12s/ispring-todo-list-api)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/ispring-todo-list-api?style=plastic)
[![Coverage Status](https://codecov.io/gh/p12s/ispring-todo-list-api/branch/master/graph/badge.svg?token=sTWAW1J7hW)](https://codecov.io/gh/p12s/ispring-todo-list-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/ispring-todo-list-api)](https://goreportcard.com/report/github.com/p12s/ispring-todo-list-api)
<img src="https://github.com/p12s/ispring-todo-list-api/workflows/lint-build/badge.svg?branch=master">

**Внимание:** *Тестовое задание найдено на просторах github-а. Для обучения и тренировки, попробовал решить ее в меру своего понимания. На ревью не отправлял, за оптимальность не ручаюсь.*

# Сервис Todo List

## Задача
Создать REST-API сервис, который позволяет вести список задач, и провести с ними простейшие действия.     
Подробнее [здесь](task.md)

## Решение  

```
git clone https://github.com/p12s/ispring-todo-list-api.git
cd ispring-todo-list-api 

docker-compose up
```
Что происходит:
- поднимается контейнер с PostgreSQL
- поднимается контейнер с сервисом [migrate](https://github.com/golang-migrate/migrate) для инициализации/миграции БД
- БД пиводится к рабочему виду, контейнер с migrate удалаяется
- поднимается контейнер с кодом приложения: [документация методов](http://localhost/swagger/index.html)

## Требования
- ✅ Язык программирования: PHP 7.x или Go   
  **Go**
- ✅ СУБД: MySQL или postgresql  
  **PostgreSQL**
- ✅ Исходный код разместить в репозитории на GitHub
- ✅ Использование фреймворков и библиотек по желанию  
  **Gin**

## Дополнительные задания
- ✅ Метод API, позволяющий получить список архивных задач (которые были помечены выполненными)
  **{{host}}:{{port}}/api/items/completed**
- ✅ набор API тестов, позволяющих проверить работу приложения в автоматическом режиме. Реализовать можно с помощью Postman (https://www.postman.com/use-cases/api-testing-automation/)  
  **[Описание](postman/README.md) запуска из командной строки**
- ✅ Dockerfile, который содержит все необходимые настройки и зависимости для быстрого разворачивания окружения.  
  **docker-compose**

## Добавил от себя  
- swagger - [документация](http://localhost:80/swagger/index.html)  
- бейдж с %-ом покрытия тестами
- бейдж с результатом запуска тестов и пробной сборки в Github Actions для master-ветки


## Что можно улучшить  
- Покрыть код unit-тестами (пока только интеграционные - с помощью postman)
- Настроить запуск сервиса с помощью docker-compose, чтобы поднимался сервис и БД (сейчас контейнер с postgres поднимаю руками + накатываю миграции)
- Использовать [FastHTTP](https://github.com/valyala/fasthttp) вместо Gin  
- Добавить индексы в БД
- Забенчмаркать с помощью wrk для проверки производительности
- Отпрофилировать с помощью pprof и по-возможности, зарефакторить. А после снова забенчмаркать  

