#/usr/bin/bash

curl --silent --location https://rpm.nodesource.com/setup_8.x | sudo bash -
yum update -y
yum install nodejs -y
yum install -y docker docker-compose

docker pull mysql:8.0
docker pull golang:1.10.4-alpine3.7
