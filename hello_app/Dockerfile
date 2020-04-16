FROM golang:1.14
LABEL maintainer="romanitalian.net@gmail.com"
WORKDIR /go/src/app
COPY ./app .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE ${APP_PORT}

CMD ["app"]
