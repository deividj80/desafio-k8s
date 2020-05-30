FROM golang:1.6.1-alpine

ENV APP_NAME main
ENV PORT 8080

COPY ./desafio-go/ /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go build main.go greeting.go

CMD ./${APP_NAME}

EXPOSE ${PORT}
