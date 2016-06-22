FROM alpine:3.4
MAINTAINER Daisuke Fujita <dtanshi45@gmail.com> (@dtan4)

ENV GOPATH /go
COPY . /go/src/github.com/dtan4/paus-watcher
RUN apk --no-cache --update add curl git go make mercurial \
    && cd /go/src/github.com/dtan4/paus-watcher \
    && make deps \
    && make \
    && mkdir /app \
    && cp bin/paus-watcher /app/paus-watcher \
    && cd /app \
    && rm -rf /go \
    && apk del --purge curl git go make mercurial

WORKDIR /app

CMD ["./paus-frontend"]
