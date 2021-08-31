build :
	@go build -o bin/tserver

run: build
	@bin/tserver

image:
	@docker build -t tserver .

docker-run: image
	@docker run --publish 8080:8080 -d tserver

fmt:
	@gofmt -s -w .
