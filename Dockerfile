FROM golang:1.19-alpine

WORKDIR /app

ADD ../.. /app

RUN go mod download

RUN go env -w GO111MODULE=on

RUN go build -o /RepositoryAPI

EXPOSE 8080

CMD [ "/RepositoryAPI" ]