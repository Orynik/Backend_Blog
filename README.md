# Backend_Blog

## Как запустить?
* * *
#### Примечание
##### Для запуска требуется наличие на устройстве запуска СУБД MySQL (Паролем на root, в конфиг файле 'Q1qqqqq!'), компилятора Golang и настроенного GOPATH
* * *
* Клонировать репозиторий к себе на компьютер по http или по SSH
> git clone https://gitlab.com/Orynik/blogongo-vue.git

* Создать базу данных и импортировать туда таблицы
> 1. `mysql -u root -p `
> 2. `create database user;`
> 3.  `use user;`
> 4. `source <путь-до-папки-с-проектом>/backend/migrations/latest.sql`
> Если все прошло удачно, то при запуске команды show tables будут отображены 4 таблицы:
> articles, сomments, users, category

* Скомпилировать файл запуска сервера и запустить
> 1. `cd backend && make build`
> 2. `./apiserver` (Для Linux) или `apiserver.exe` (для Windows)

* Запустить средство установки зависимостей для фронтенд части
> 1. `cd - && npm install`
> 2. `npm run serve`

* Перейти по ссылке, которая была выведена в конце прошлого пункта (обычно, это localhost:8080)
* RESTful-api располагается на localhost:4444
