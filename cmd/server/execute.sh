cd ../wasm
GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm
cd ../server
go run .