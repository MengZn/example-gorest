FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody


EXPOSE 8080

ADD _output/bin/rest /usr/local/bin/rest


