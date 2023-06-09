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
Run the commands below to install & configure msyql
```
# Install mysql container
sudo docker run -p 3306:3306 --name mysql \
-v /home/janessatech/docker/mysql/data:/data \
-e MYSQL_ROOT_PASSWORD=123456 \
--restart=always \
-itd mysql:5.7.19

# enter the container just created
docker exec -it mysql /bin/bash

# enter mysql
mysql -uroot -p123456

show databases;
# Create a database and username/password pair for that database
CREATE DATABASE gin_micro_template;
CREATE USER 'templateuser'@'%' IDENTIFIED BY 'templatepwd';
SELECT USER, host from mysql.user;
GRANT ALL on *.* TO 'templateuser'@'%';
GRANT super on *.* to 'templateuser'@'%';
GRANT show view on *.* to 'templateuser'@'%';

# Verify if you could access the database gin_micro_template using username and password
mysql -utemplateuser -ptemplatepwd gin_micro_template

# Creat tables
CREATE TABLE accounts
(
`id`              INT NOT NULL AUTO_INCREMENT,
`user_name`       VARCHAR(125) NOT NULL,
`password`        VARCHAR(125) NOT NULL,
`created_at`      TIMESTAMP NULL,
`updated_at`      TIMESTAMP NULL,
PRIMARY KEY (`id`),
UNIQUE INDEX unique_user_name_idx(`user_name`)
)
ENGINE = InnoDB
DEFAULT CHARSET = UTF8MB4;
```
## Starting commands
Run one of the three commands below to start a web server:
```
.\go-microservice-template.exe
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
    "code": 200,
    "message": "success",
    "details": {
        "id": 1,
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
    "code": 200,
    "message": "success",
    "details": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2ODYyODA5NzIsImlkIjoxfQ.gnJ6BkFKgW1mcyw6ypssXMztkAs-M_9XM8maLEGtodM"
}
```
jwt code is contained in details in the response. We will use it as Authorization later on
### 3.Check current user
In postman, run the url below with GET method:
http://127.0.0.1:8080/api/account/me
In Authorization tab, select Bearer Token as Type, input the jwt value returned above as token
You will receive a response like below:
```
{
    "code": 200,
    "message": "success",
    "details": {
        "id": 1,
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
### Product operations(CRUD)
#### Add a new product
In postman, run the url below with POST method:
http://127.0.0.1:8080/api/products
In Authorization tab, select Bearer Token as Type, input the jwt value returned above as token
Input json content as body as below:
```
{
    "name" : "product1"
}
```
You will receive a response like below:
```
{
    "code": 200,
    "message": "success",
    "details": {
        "id": 1,
        "name": "product1",
        "createdAt": "2023-06-09T11:11:39.098+08:00",
        "updatedAt": "2023-06-09T11:11:39.098+08:00"
    }
}
```
You will receive a reponse like below if token is not correct
```
{
    "error": "anthentication is failed"
}
```

#### view all products
In postman, run the url below with GET method:
http://127.0.0.1:8080/api/products
In Authorization tab, select Bearer Token as Type, input the jwt value returned above as token
You will receive a response like below:
```
{
    "code": 200,
    "message": "success",
    "details": [
        {
            "id": 1,
            "name": "product1",
            "createdAt": "2023-06-09T11:11:39+08:00",
            "updatedAt": "2023-06-09T11:11:39+08:00"
        }
    ]
}
```
You will receive a reponse like below if token is not correct
```
{
    "error": "anthentication is failed"
}
```

#### delete a product
In postman, run the url below with DELETE method:
http://127.0.0.1:8080/api/products/1
In Authorization tab, select Bearer Token as Type, input the jwt value returned above as token
You will receive a response like below:
```
{
    "code": 200,
    "message": "success",
    "details": null
}
```
You will receive a reponse like below if token is not correct
```
{
    "error": "anthentication is failed"
}
```

