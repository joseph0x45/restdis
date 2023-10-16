build:
	@go build -o restdis main.go

install-deps:
	@go mod download

launch:
	@./restdis launch
