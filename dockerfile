FROM golang:latest

LABEL version="1.0"
LABEL maintainer="Tsheri Sherpa <tsherisherpa@gmail.com>"
LABEL description="This docker image is for handling web server using go fiber framework."

WORKDIR /app

RUN apt-get update && \
apt-get install -y build-essential && \
apt-get install -y software-properties-common && \
apt-get install -y curl git vim wget && \
apt-get -y install nodejs npm

RUN npm -y install -g nodemon

COPY go.mod .

COPY go.sum .

RUN go mod download

ADD . /app/

RUN groupadd -r docker && useradd -g docker docker

RUN chown -R docker:docker /app

EXPOSE 5000