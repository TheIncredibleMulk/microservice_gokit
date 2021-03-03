# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="TheIncredibileMulk <andrew.mulkey@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
# WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
# RUN go mod tidy

RUN go mod download 
# RUN go get -u github.com/olivere/elastic github.com/olivere/elastic/v7 github.com/sirupsen/logrus gopkg.in/sohlich/elogrus.v7 

# Copy the source from the current directory to the working Directory inside the container 

COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a main .

EXPOSE 8080