FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody


EXPOSE 8080

ADD build/_output/bin/gorest /usr/local/bin/gorest


