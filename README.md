# About this project
This project provides a gin-based micro-service template which enables you to develop a micro-service application in a fast way.
It uses the following tools:
- gin : Http web framework
- fx : A dependency injection system
- gorm : Database access solution
- koanf : a simple, extremely lightweight, extensible, configuration management library
- cobra : A CLI application
- zap :  A fast, structured, leveled logging

# How to run this project
## Prepare MYSQL
I assume you have docker installed. We will use docker to install mysql
Run the commands below to install msyql and configure mysql
```
# Install mysql container
sudo docker run -p 3306:3306 --name mysql \
-v /home/jane/docker/mysql/data:/data \
-e MYSQL_ROOT_PASSWORD=123456 \
--restart=always \
-itd mysql:5.7.19
# enter the container just created
docker exec -it mysql /bin/bash
# enter mysql
mysql -uroot -p123456
show databases;
# Create a database and user/password pair for that database
CREATE DATABASE gin-micro-template;
CREATE USER 'templateuser'@'%' IDENTIFIED BY 'templatepwd';
SELECT USER, host from mysql.user;
GRANT ALL on *.* TO 'gorm'@'%';
GRANT super on *.* to 'gorm'@'%';
GRANT show view on *.* to 'gorm'@'%';
# Verify if you could access the database gin-micro-template using username and password
mysql -utemplateuser -ptemplatepwd gin-micro-template
show tables
// no tables
```
## Start commands
Run one of two commands below to start a web server:
```sh
.\go-microservice-template.exe server -c "./config/properties.json"
.\go-microservice-template.exe -c "./config/properties.json"
```
## Restful APIs
### 1.Register an account
In postman, 
### 2.Login
### 3.Check current user
### Product operations(CURD)
