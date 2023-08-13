run-dev:
	make build
	make run

build:
	@go build .

run:
	./restdis

install-deps:
	@go mod tidy
