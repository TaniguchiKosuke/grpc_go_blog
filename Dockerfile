FROM golang:1.17-alpine as dev
ENV TZ=Asia/Tokyo

ADD . /go/src/app
WORKDIR /go/src/app

RUN apk add --no-cache alpine-sdk build-base
RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

# CMD ["go", "run", "cmd/main.go"]