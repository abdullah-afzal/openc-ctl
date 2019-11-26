FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/openc-ctl_linux_386/openc-ctl /usr/local/bin/

ENTRYPOINT ["openc-ctl"]
CMD [ "--help" ]