# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /service .

##
## Deploy
##
FROM jrottenberg/ffmpeg:4.3-alpine

WORKDIR /

COPY --from=build /service ./server
RUN mkdir -p /media/images
RUN mkdir -p /media/videos
RUN mkdir -p /media/videos/raw
RUN mkdir -p /media/videos/dash

ENV PORT=5000
EXPOSE 5000

ENTRYPOINT ["/server"]
