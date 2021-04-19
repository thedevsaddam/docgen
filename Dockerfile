FROM golang:1.16.3-alpine3.13 as build

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN apk update && \
    apk upgrade && \
    apk add git

RUN git clone https://github.com/thedevsaddam/docgen.git .

RUN go generate
RUN go build -o main .
RUN chmod +x main 

FROM alpine:3.13

COPY --from=build /build/main /usr/local/bin/docgen

RUN mkdir /export
VOLUME "/export"
WORKDIR "/export"

ENTRYPOINT [ "docgen", "build", "-i", "postman_collection.json" ]