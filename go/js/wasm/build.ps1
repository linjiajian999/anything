$env:GOARCH="wasm";
$env:GOOS="js";
go build -o index.wasm main.go;