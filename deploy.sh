#! /bin/bash

docker-compose down

docker rmi gp-server_web gp-server_file

git pull origin

docker-compose up --build -d