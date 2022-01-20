FROM golang:1.16-alpine
WORKDIR /api

COPY . ./

# add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# prevent the re-installation of vendors at every change in the source code
# COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

# Install Compile Daemon for go. We'll use it to watch changes in go files
# RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

# RUN chmod +x main.go

# RUN CompileDaemon --build="go build -o server main.go" --command=./server
RUN go build -o server main.go

CMD [ "/api/server" ]
# ENTRYPOINT [ "server" ]