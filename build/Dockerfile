FROM golang:1.18.3-alpine3.16 AS builder

RUN apk update && apk add bash && apk --no-cache add tzdata

RUN go env -w GOPRIVATE=github.com/phamdinhha

WORKDIR /build

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN apk --update add git less openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

RUN ls -la

COPY [".", "./"]

RUN go mod tidy



RUN go build -v -o go-authorizer ./cmd/go-authorizer/


#RUN LAYER
FROM alpine

WORKDIR /

RUN apk update && apk add bash && apk --no-cache add tzdata

COPY --from=builder ["/build/go-authorizer", "/build/config/config.yml", "/build/model.conf", "/build/policy.csv", "/"]

CMD ["/go-authorizer"]
