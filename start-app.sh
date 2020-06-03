#!/bin/bash

docker build -t appointment .
docker-compose up -d --remove-orphans
