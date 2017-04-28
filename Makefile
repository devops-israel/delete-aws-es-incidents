clean:
	go clean -i ./...

deps:
	glide install
	glide update --all-dependencies

build_inside_docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o delete-aws-es-incidents .

build_osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o delete-aws-es-incidents .

build:
	docker-compose build

serve:
	docker-compose up
