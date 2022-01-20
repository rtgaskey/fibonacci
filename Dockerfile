FROM golang:1.16-alpine
WORKDIR /api

COPY . ./

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

RUN go mod download && go mod verify

EXPOSE 8080

RUN go build -o server main.go

CMD [ "/api/server" ]