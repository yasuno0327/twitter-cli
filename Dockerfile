FROM golang:1.13-alpine

ENV GO111MODULE=on
EXPOSE 3000

RUN mkdir -p /github.com/yasuno0327/twitter-cli
WORKDIR /github.com/yasuno0327/twitter-cli
ADD . .

CMD [ "go", "run", "main.go"]