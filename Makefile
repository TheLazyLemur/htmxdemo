generate:
	cd db && sqlc generate
	templ generate

run: generate
	go run ./cmd/htmxdemo

build: generate
	go build -o ./bin/htmxdemo ./cmd/htmxdemo
	upx --lzma ./bin/htmxdemo

clean:
	rm -rf ./bin
