run:
	cd ./cmd/api;go run main.go
test:
	go test ./...
build:
	@ printf "building app.. "
	@ go -C ./cmd/api build -o ../../bin/ main.go;cp ./configs/config.json ./bin;cp ./web/* ./bin
	@ echo "[done]"	
