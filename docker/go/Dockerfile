FROM golang:latest

WORKDIR /go/src/todos-rest

COPY ./docker/go/entrypoint.sh /root/
RUN chmod 755 /root/entrypoint.sh

# Project files
COPY . .
RUN swagger generate spec -o doc/swagger.json \
    && go install

ENTRYPOINT ["/root/entrypoint.sh"]
