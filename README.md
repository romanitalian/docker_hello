# docker_hello
Practice in docker

### Шаг 1
> docker build -t my-hello .

> docker images

> docker run my-hello

Cписок запущенных контейнеров:
> docker ps

Cписок запущенных контейнеров (даже которые уже остановились):
> docker ps -a

Запустить контейнер и задать ему имя:
> docker run --name=foo my-hello

Запустить контейнер в фоне (`-d`):
> docker run --name=foo -d my-hello

Удалить контейнер (либо по ID, либо по имени):
> docker rm ID

> docker rm NAME

> docker rm foo

Показать список контейнеров и вывести только их ID:
> docker ps -a -q

Удалить все контейреры:
> docker rm $(docker ps -qa)

Остановить контейнер
> docker stop <NAME>

> docker stop my-hello

Запустить контейнер и пробросить порт (и при этом в Dockerfile уже есть EXPOSE 80):
> docker run --rm --name foo -p 8080:80 my-hello

### Шаг 2

> docker build -t my-golang-app .

> docker run -it --rm --name my-running-app my-golang-app

### Шаг 3

> docker build -t my-go-server .

> docker run -it    --rm --name my-running-goserv my-go-server

> docker run -it -d --rm --name my-running-goserv my-go-server


Пробросить наружу порт 8080, при том, что внутри контейнера приложение отдаёт на 80:
> docker build -t mm .

> docker run -it -d --rm --name foo -p 8080:80 mm

> docker rmi mm

Добавить переменную в команде запуска:
> docker run -it -d --rm --name foo -p 8080:80 -e APP_PORT=80 mm

Примонтировать каталог в контейнер:
> docker run -it -d --rm --name foo -p 8080:80 -v /Users/rmn/wks/src/roman.dev/docker_hello:/go/src/app mm

Используем volume (чтобы иметь доступ извне в наш контейнер)

Показать какие volume сейчас есть:
> docker volume ls

Создать volume `moo`:
> docker volume create moo

И теперь `moo` можем использовать для монтирования:
> docker run -it --rm --name foo -p 8080:80 -v moo:/go/src/app mm
