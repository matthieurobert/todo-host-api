FROM golang:alpine3.13 AS builder
WORKDIR /go/src/github.com/todo-host/todo-host-api
COPY . /go/src/github.com/todo-host/todo-host-api/
RUN set -xe; \
    apk add --no-cache git && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-binary

#       FINAL IMAGE
FROM alpine:3.13
#   LABELLING
LABEL maintainer="Matthieu ROBERT <matthieurobert82@gmail.com>"
LABEL version="alpha-0.1"
LABEL description="A docker container to run the todo-host API"
#   SETUP
WORKDIR /app
COPY --from=builder /go/src/github.com/todo-host/todo-host-api .
EXPOSE 8000
#   RUN
CMD ./api-binary