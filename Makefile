air:
	~/.air

templ-generate:
	templ generate

tailwind-watch:
	~/tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

tailwind-build:
	~/tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

run:
	make templ-generate
	go build -o ./tmp/$(APP_NAME) ./cmd/main.go
	make air

build:
	make templ-generate
	go build -o ./bin/${APP_NAME} ./cmd/main.go
