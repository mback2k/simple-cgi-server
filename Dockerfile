FROM golang:alpine as build
RUN apk --no-cache --update upgrade && apk --no-cache add git build-base

ADD . /go/simple-cgi-server
WORKDIR /go/simple-cgi-server

RUN go get
RUN go build -ldflags="-s -w"
RUN chmod +x simple-cgi-server

FROM mback2k/alpine:latest
RUN apk --no-cache --update upgrade && apk --no-cache add ca-certificates

COPY --from=build /go/simple-cgi-server/simple-cgi-server /usr/local/bin/simple-cgi-server

RUN addgroup -g 8080 -S serve
RUN adduser -u 8080 -h /data -S -D -G serve serve

WORKDIR /data
USER serve

CMD [ "/usr/local/bin/simple-cgi-server" ]
