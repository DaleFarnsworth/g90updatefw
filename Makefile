default: all

all: linux

linux: main.go
	GOOS=linux ARCH=amd64 go build -ldflags="-s -w"
	upx --brute g90updatefw
