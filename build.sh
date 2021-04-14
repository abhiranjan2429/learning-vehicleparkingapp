#!/bin/sh
sudo docker build -t parkinglot:latest .
sudo docker-compose up -d
#sudo docker run -d --rm --network container:broker parkinglot ./servicestarter
#apt install iputils-ping