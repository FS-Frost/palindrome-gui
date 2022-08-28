build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o app

run: build
	./app -mode=server --dev true --autoupdate true

static: build
	rm -rf docs
	mkdir -p docs/web
	cp -r web docs
	./app -mode=static
