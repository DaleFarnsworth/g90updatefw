FLAGS = -ldflags "-s -w"

default: linux/amd64/g90updatefw

all: darwin/amd64/g90updatefw linux/386/g90updatefw linux/amd64/g90updatefw linux/arm/g90updatefw linux/arm64/g90updatefw linux/ppc64/g90updatefw linux/ppc64le/g90updatefw linux/riscv64/g90updatefw linux/s390x/g90updatefw windows/386/g90updatefw.exe windows/amd64/g90updatefw.exe

darwin/amd64/g90updatefw: main.go
	GOOS=darwin GOARCH=amd64 go build $(FLAGS) -o darwin/amd64/g90updatefw

darwin/arm64/g90updatefw: main.go
	GOOS=darwin GOARCH=arm64 go build $(FLAGS) -o darwin/arm64/g90updatefw

linux/386/g90updatefw: main.go
	GOOS=linux GOARCH=386 go build $(FLAGS) -o linux/386/g90updatefw

linux/amd64/g90updatefw: main.go
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o linux/amd64/g90updatefw
	upx --brute linux/amd64/g90updatefw

linux/arm/g90updatefw: main.go
	GOOS=linux GOARCH=arm go build $(FLAGS) -o linux/arm/g90updatefw

linux/arm64/g90updatefw: main.go
	GOOS=linux GOARCH=arm64 go build $(FLAGS) -o linux/arm64/g90updatefw

linux/ppc64/g90updatefw: main.go
	GOOS=linux GOARCH=ppc64 go build $(FLAGS) -o linux/ppc64/g90updatefw

linux/ppc64le/g90updatefw: main.go
	GOOS=linux GOARCH=ppc64le go build $(FLAGS) -o linux/ppc64le/g90updatefw

linux/riscv64/g90updatefw: main.go
	GOOS=linux GOARCH=riscv64 go build $(FLAGS) -o linux/riscv64/g90updatefw

linux/s390x/g90updatefw: main.go
	GOOS=linux GOARCH=s390x go build $(FLAGS) -o linux/s390x/g90updatefw

windows/386/g90updatefw.exe: main.go
	GOOS=windows GOARCH=386 go build $(FLAGS) -o windows/386/g90updatefw.exe

windows/amd64/g90updatefw.exe: main.go
	GOOS=windows GOARCH=amd64 go build $(FLAGS) -o windows/amd64/g90updatefw.exe

clobber:
	rm -f */*/g90updatefw*
