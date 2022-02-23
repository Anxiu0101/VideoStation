FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/Anxiu0101/VideoStation
COPY . $GOPATH/src/github.com/Anxiu0101/VideoStation
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./VideoStation"]