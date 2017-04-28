FROM scratch

ARG version
ENV VERSION=$version

ADD https://github.com/joshdvir/delete-aws-es-incidents/releases/download/v{$VERSION}/delete-aws-es-incidents-{$VERSION}-linux-amd64 /delete-aws-es-incidents

EXPOSE 9292

CMD ["-h"]
ENTRYPOINT ["/delete-aws-es-incidents"]