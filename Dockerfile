FROM scratch

ADD delete-aws-es-incidents /delete-aws-es-incidents

CMD ["-h"]

ENTRYPOINT ["/delete-aws-es-incidents"]