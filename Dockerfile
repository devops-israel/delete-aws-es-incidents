FROM scratch

ARG version=0.0.0

ADD https://github.com/joshdvir/delete-aws-es-incidents/releases/download/v{$version}/delete-aws-es-incidents-{$version}-linux-amd64 /delete-aws-es-incidents

CMD ["-h"]

ENTRYPOINT ["/delete-aws-es-incidents"]