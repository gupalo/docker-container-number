FROM golang:1.13

RUN go get -t github.com/tidwall/gjson

WORKDIR /code
COPY ./server.go .

RUN go get -d -v ./...
RUN go build server.go

EXPOSE 80

CMD ["/code/server", "-port", "80"]