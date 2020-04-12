# docker_hello
Practice in docker


1.
> docker build -t my-hello .
> docker images
> docker run my-hello

Cписок запущенных контейнеров:
> docker ps

Cписок запущенных контейнеров (даже которые уже остановились):
> docker ps -a

Запустить контейнер и задать ему имя:
> docker run --name=foo my-hello

Удалить контейнер (либо по ID, либо по имени):
> docker rm <ID>
> docker rm <NAME>
> docker rm foo

Показать список контейнеров и вывести только их ID:
> docker ps -a -q

Удалить все контейреры:
> docker rm $(docker ps -qa)

2.
docker build -t my-golang-app .
docker run -it --rm --name my-running-app my-golang-app