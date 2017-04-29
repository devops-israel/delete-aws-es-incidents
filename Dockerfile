FROM scratch

ADD delete-aws-es-incidents /delete-aws-es-incidents

CMD ["/delete-aws-es-incidents"]