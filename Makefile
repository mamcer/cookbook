run:
	cd ./cmd/api;go run main.go
test:
	go test ./...
build:
	go -C ./cmd/api build -o ../../bin/ main.go;cp ./cmd/api/config.json ./bin;cp ./web/* ./bin
	
