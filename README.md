# sanic.exe

Dumb meme for kid.

```bash
go mod tidy

go build -o bin/sanic sanic.go

GOOS=windows GOARCH=amd64 go build -o bin/sanic.exe sanic.go

# ?
# GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" go build -o bin/sanic.exe sanic.go
```
