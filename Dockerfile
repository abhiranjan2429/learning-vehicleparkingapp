####################################################################
# ParkingLot Dockerfile for linux runtime package
#
# Author: Abhiranjan Sigh
# email: abhiranjan2429@gmail.com
# 
####################################################################
# # Get base image 
FROM debian:stretch-slim as build-apps
#Get Updates and install ==> curl, git and go  
RUN apt-get update
RUN apt install -y curl
RUN apt install -y git
RUN curl -O https://dl.google.com/go/go1.14.linux-amd64.tar.gz
RUN tar xvf go1.14.linux-amd64.tar.gz
RUN chown -R root:root ./go
RUN mv go /usr/local
# Setting the env for go 
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOPATH/bin:$GOROOT/bin
#ENV GOPROXY="https://proxy.golang.org,direct"

#Create new Directory ("vehicleparkingapp") and copy Filestructure
RUN mkdir -p /vehicleparkingapp
ADD . /vehicleparkingapp
WORKDIR /vehicleparkingapp

#Building the apps => servicestarter vehicleentry and vehicleexit
RUN go build -o "./bin/" "./cmd/servicestarter/"
RUN go build -o "./bin/" "./cmd/vehicleentry/"
RUN go build -o "./bin/" "./cmd/vehicleexit/"
RUN go build -o "./bin" "./pkg/webApis/webbillingandtime"
RUN go build -o "./bin" "./pkg/webApis/webexit"
RUN go build -o "./bin" "./pkg/webApis/webentry"
RUN chmod +x -R "./bin/"
# RUN chmod +x "./bin/vehicleentry"
# RUN chmod +x "./bin/vehicleexit"
# Stage 2 of buildcreation to keep the minimal size of build Image
FROM debian:stretch-slim
RUN apt-get update
RUN apt install -y iputils-ping
RUN mkdir -p /vehicleparkingapp
COPY --from=build-apps ./vehicleparkingapp/bin/ ./vehicleparkingapp/bin
WORKDIR "./vehicleparkingapp/bin"
#CMD ["./vehicleparkingapp/bin/servicestarter"]