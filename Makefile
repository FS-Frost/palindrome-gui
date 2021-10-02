build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	mkdir -p docs/web
	cp web/app.wasm docs/web/app.wasm
	go build -o app

run: build
	./app
