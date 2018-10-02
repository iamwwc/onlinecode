#/usr/bin/bash

# build env
curl --silent --location https://rpm.nodesource.com/setup_8.x | sudo bash -
yum update -y
yum install nodejs -y
yum install -y docker docker-compose

# needed images
docker pull mysql:8.0
docker pull golang:1.10.4-alpine3.7
docker pull python:3.6-alpine3.6
docker pull java:8u111-jdk-alpine


cd ../client-side
npm install && npm run build
cd ..
cp client-side/dist/* . -R
#rm -Rf client-side
docker build --tag 0.0.1-alpine3.7 .