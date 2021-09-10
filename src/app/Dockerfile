# FROM golang:1.17 as base
FROM golang:1.17-alpine as base

FROM base as dev

RUN wget --no-check-certificate -q -O - https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/web

EXPOSE 8080

CMD ["air"]