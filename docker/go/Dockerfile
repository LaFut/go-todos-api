FROM golang:latest

WORKDIR /go/src/todos-rest

RUN go get github.com/go-swagger/go-swagger/cmd/swagger

COPY ./docker/go/entrypoint.sh /root/
RUN chmod 755 /root/entrypoint.sh

# Project files
COPY . .
RUN swagger generate spec -o doc/swagger.json \
    && go install

ENTRYPOINT ["/root/entrypoint.sh"]
