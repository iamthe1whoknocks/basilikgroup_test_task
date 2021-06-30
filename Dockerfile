FROM golang:1.15-alpine3.12 AS builder

COPY . /bazilik-app
WORKDIR /bazilik-app

RUN go mod download
RUN GOOS=linux go build -o ./cmd/app/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /bazilik-app/

COPY --from=0 /bazilik-app/cmd/app/app .
COPY --from=0 /bazilik-app/cmd/app/config.yaml .

EXPOSE 8086

CMD ["./app"]