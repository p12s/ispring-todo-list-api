# Postman - тестирование

## Когда Postman подойдет, а когда нет
- ✅ коммуникация с API во время разработки
- ✅ автоматическое тестирование API (в CI/CD)
- ❌ тестирование производительности (лучше использовать wrk, jmeter, ...)
- ❌ тестирование безопасности

## Запуск тестов
На локале:  
```
brew install newman     # установка newman
newman run postman/ispring-todo-list-api.postman_collection.json    # запуск тестов
```
Есть возможность встроить такой запуск в CI/CD, но пока это не делал.