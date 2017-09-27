# Delete AWS ES Incidents (indexes)

A small application that deletes indexes in AWS Elasticsearch service.

The purpose of this application to be a workaround of using curator with AWS Elasticsearch service, at the moment of writing there is an [issue deleting indexes using curator](https://www.elastic.co/guide/en/elasticsearch/client/curator/current/faq_aws_iam.html)

## Getting Started

Application is packed as a single binary, just download and run.

### Prerequisites

Nothing.

### Usage

The help section explains everything:

```
Delete ELK incidents on AWS ES 5.1

Usage:
  delete-aws-es-incidents [flags]
  delete-aws-es-incidents [command]

Available Commands:
  help        Help about any command
  version     Display version

Flags:
  -e, --es-url string            Elasticsearch URL, eg. https://path-to-es.aws.com/
  -h, --help                     help for delete-aws-es-incidents
  -d, --older-than-in-days int   delete incidents older then in days (default 14)
  -p, --prefixes string          comma separated list of prefixes for indexs, index date must be in format YYYY.MM.DD. eg. 'logstash-2017.09.28'. default is 'logstash-' (default "logstash-")

Use "delete-aws-es-incidents [command] --help" for more information about a command.
```

## Deployment

Easiest way to deploy is via the [docker image](https://hub.docker.com/r/joshdvir/delete-aws-es-incidents/)

Example usage deploying the docker image:

```
docker run -d joshdvir/delete-aws-es-incidents -e https://path-to-es-server/ -d 30
```

Or you can always download a [release compatible to your OS](https://github.com/devops-israel/delete-aws-es-incidents/releases) and run the application.

## Built With

* [Golang](https://golang.org/)
* [Cobra](https://github.com/spf13/cobra) - CLI framework for Go.
* [Elasticsearch Go Client](https://github.com/olivere/elastic)
* [Cron](https://github.com/robfig/cron) - A cron library for Go

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md)

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/devops-israel/delete-aws-es-incidents/tags).

## Authors

* [**Josh Dvir**](https://github.com/joshdvir)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
