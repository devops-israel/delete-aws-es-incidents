FROM alpine:3.6

ADD delete-aws-es-incidents-linux-amd64 /delete-aws-es-incidents

CMD ["-h"]

ENTRYPOINT ["/delete-aws-es-incidents"]