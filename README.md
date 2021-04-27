# GoTodo

A simple Todo-list app for everyday tasks made using html, css, javascript for frontend and Golang for backend server.

# Pre-Requisites

Docker<br/>
Golang

# Installation

Run the commands given below in a terminal:-

docker pull mysql:latest<br/>
docker start mysql<br/>
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql<br/>
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE GoTodo'<br/>
./GoTodo<br/>

# Screenshot

![Screenshot (83)](https://user-images.githubusercontent.com/32940477/116273006-04f90780-a79f-11eb-87ff-74edf8e6f87f.png)
