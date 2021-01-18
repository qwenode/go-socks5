CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o \
 release/s5 cmd/main.go && upx -9 release/s5

