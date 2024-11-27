# 2024_2_GOATS
Backend репозиторий команды GOATS 🐐🐐🐐🐐

## Состав команды

[Ягмуров Игорь](https://github.com/UnicoYal) - *Backend-разработчик*

[Угаров Руслан](https://github.com/Rusy13) - *Backend-разработчик*

[Ашуров Георгий](https://github.com/AshurovG) - *Frontend-разработчик*

[Койбаев Тамерлан](https://github.com/tkoibaev) - *Frontend-разработчик*

## Менторы

[Павловский Андрей](https://github.com/Starlexxx) - *Backend*

[Клонов Александр](https://github.com/Shureks-den) - *Frontend*

Мартынова Галина - *UX*


## Ссылки

[Фронтенд проекта](https://github.com/frontend-park-mail-ru/2024_2_GOATS)

[Деплой](http://83.166.232.3/)

## Тесты

Перед прогоном тестов запустите докер. Тесты постгреса требуют запущенного докера
```
make all или

go test -coverprofile=coverage.out ./...
./filter_coverage.sh coverage.out exclude_from_coverage.txt
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```
