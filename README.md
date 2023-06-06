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
In postman, run the url below with POST method:
http://127.0.0.1:8080/api/account/register
In body tab, choose raw radio box and select JSON from dropdown list. Input json content as body as below:
```
{
    "username" : "JanessaTech1",
    "password" : "12345"
}
```
You will get a response like below:
```
{
    "savedAccDto": {
        "id": 0,
        "username": "JanessaTech1",
        "password": ""
    }
}
```
### 2.Login
In postman, run the url below with POST method:
http://127.0.0.1:8080/api/account/login
In body tab, choose raw radio box and select JSON from dropdown list. Input json content as body as below:
```
{
    "username" : "JanessaTech1",
    "password" : "12345"
}
```
You will get a response like below:
```
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2ODYwNTE5NTQsImlkIjowfQ.aozM7M3QSUXMBWl20IGg30Nph8FY9966KehVu_76jW8"
}
```
### 3.Check current user
In postman, run the url below with GET method:
http://127.0.0.1:8080/api/account/me
In Authorization tab, select Bearer Token as Type, input the jwt value returned above as token
You will receive a response like below:
```
{
    "current account": {
        "id": 0,
        "username": "JanessaTech1",
        "password": ""
    }
}
```
You will receive a reponse like below if token is not correct
```
{
    "error": "anthentication is failed"
}
```
### Product operations(CURD)
