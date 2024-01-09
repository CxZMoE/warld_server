go env -w GOOS=windows
go env -w GOARCH=amd64
go env -w CGO_ENABLED=0

go build -o warld_server.exe
