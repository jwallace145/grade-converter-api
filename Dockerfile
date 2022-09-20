# syntax=docker/dockerfile:1

FROM golang:1.19-buster

WORKDIR /app/grade-converter-api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./out/grade-converter-api

EXPOSE 8080

CMD [ "./out/grade-converter-api" ]