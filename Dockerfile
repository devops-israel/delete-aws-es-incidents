FROM scratch

ADD delete-aws-es-incidents /delete-aws-es-incidents

EXPOSE 9292

CMD ["/delete-aws-es-incidents"]