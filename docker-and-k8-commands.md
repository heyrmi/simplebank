# Contains all the commands related to migrations, docker and K8

## Migrations:

`brew install golang-migrate` : To install golang migrate CLI

`migrate create -ext sql -dir db/migration -seq init_schema` : To create initial schema

## Docker

`docker ps` : To preview all running containers

`docker images` : To list all available docker images

`docker pull <imageName>:<tag>` : To pull docker image with specified tag

`docker inspect <image_name>` : To get all info on docker image

`docker run --name <container_name> -e <environment_variable> -d <image>:<tag>` : To run the docker container

`docker exec -it <container_name_or_id> <command> [args]` : To run a command inside running container

`docker logs <container_name_or_id>` : To see docker logs

`docker start <container_name_or_id>` : To start any stopped docker container

`docker stop <container_name_or_id>` : To stop any running docker container

`docker rm <container_name_or_id>` : To remove any docker container completely

`docker exec -it <container_name_or_id> /bin/sh` : To access shell incase of alpine image based containers

- Postgres Start Script: `docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine`

- Postgres SQL Enter Command: `docker exec -it postgres16 psql -U root` use `\q` to exit it.
