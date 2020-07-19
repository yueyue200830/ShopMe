FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/backend-go
COPY . $GOPATH/src/backend-go
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./backend-go"]