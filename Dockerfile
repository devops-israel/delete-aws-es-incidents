FROM alpine:3.6

RUN apk add --update --no-cache ca-certificates

ADD delete-aws-es-incidents-linux-amd64 /delete-aws-es-incidents

CMD ["-h"]

ENTRYPOINT ["/delete-aws-es-incidents"]