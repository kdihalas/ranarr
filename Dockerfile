FROM alpine:3.6

MAINTAINER Konstantinos Dichalas <kdihalas@gmail.com>

ADD bin/linux/ranarr /bin/ranarr

RUN apk add --update ca-certificates

ENTRYPOINT [ "/bin/ranarr" ]