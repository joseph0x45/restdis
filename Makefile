build:
	@go build -o restdis main.go

install-deps:
	@go mod download

start-tailwind-compilation:
	@npx tailwindcss -i ./assets/app.css -o ./public/output.css --minify --watch

build-css:
	@npx tailwindcss -i ./assets/app.css -o ./public/output.css --minify
