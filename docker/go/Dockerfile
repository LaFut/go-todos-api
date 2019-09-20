FROM golang:latest

RUN mkdir -p /go/src/go-todos-api

WORKDIR /go/src/go-todos-api

RUN go get github.com/go-swagger/go-swagger/cmd/swagger

COPY ./docker/go/entrypoint.sh ./docker/wait-for-it.sh /root/

RUN chmod 755 /root/entrypoint.sh /root/wait-for-it.sh

COPY . /go/src/go-todos-api

RUN swagger generate spec -o doc/swagger.json \
    && go install

ENTRYPOINT ["/root/entrypoint.sh"]
