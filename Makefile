air:
	~/.air

templ-generate:
	templ generate

run:
	make templ-generate
	go build -o ./tmp/$(APP_NAME) ./cmd/main.go
	make air

build:
	make templ-generate
	go build -o ./bin/${APP_NAME} ./cmd/main.go
