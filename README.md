# GoTodo

A simple Todo-list app for everyday tasks made using html, css, javascript for frontend and Golang for backend server.

# Pre-Requisites

Docker
Golang

# Installation

Run the commands given below in a terminal:-

docker pull mysql:latest
docker start mysql
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE GoTodo'
./GoTodo
