FROM scratch

ARG version=0.0.0

ADD https://github.com/joshdvir/delete-aws-es-incidents/releases/download/v{$VERSION}/delete-aws-es-incidents-{$VERSION}-linux-amd64 /delete-aws-es-incidents

CMD ["-h"]

ENTRYPOINT ["/delete-aws-es-incidents"]