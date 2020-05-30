#!/bin/bash

docker-compose down && docker-compose up -d

export DBUSER=root
export DBNAME=appointment
export DBPASSWORD=prueba12345
export DBHOST=localhost
export DBPORT=3306

cd http/cmd/a-p-pointment-backend-server
go run main.go
