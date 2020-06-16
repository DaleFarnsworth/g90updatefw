FLAGS = -ldflags "-s -w"

default: linux/amd64/g90updatefw

all:darwin/386/g90updatefw darwin/amd64/g90updatefw darwin/arm/g90updatefw darwin/arm64/g90updatefw linux/386/g90updatefw linux/amd64/g90updatefw linux/arm/g90updatefw linux/arm64/g90updatefw linux/mips/g90updatefw linux/mips64/g90updatefw linux/mips64le/g90updatefw linux/mipsle/g90updatefw linux/ppc64/g90updatefw linux/ppc64le/g90updatefw linux/riscv64/g90updatefw linux/s390x/g90updatefw

darwin/386/g90updatefw: main.go
	GOOS=darwin ARCH=386 go build $(FLAGS) -o darwin/386/g90updatefw

darwin/amd64/g90updatefw: main.go
	GOOS=darwin ARCH=amd64 go build $(FLAGS) -o darwin/amd64/g90updatefw

darwin/arm/g90updatefw: main.go
	GOOS=darwin ARCH=arm go build $(FLAGS) -o darwin/arm/g90updatefw

darwin/arm64/g90updatefw: main.go
	GOOS=darwin ARCH=arm64 go build $(FLAGS) -o darwin/arm64/g90updatefw

linux/386/g90updatefw: main.go
	GOOS=linux ARCH=386 go build $(FLAGS) -o linux/386/g90updatefw

linux/amd64/g90updatefw: main.go
	GOOS=linux ARCH=amd64 go build $(FLAGS) -o linux/amd64/g90updatefw
	upx --brute linux/amd64/g90updatefw

linux/arm/g90updatefw: main.go
	GOOS=linux ARCH=arm go build $(FLAGS) -o linux/arm/g90updatefw

linux/arm64/g90updatefw: main.go
	GOOS=linux ARCH=arm64 go build $(FLAGS) -o linux/arm64/g90updatefw

linux/mips/g90updatefw: main.go
	GOOS=linux ARCH=mips go build $(FLAGS) -o linux/mips/g90updatefw

linux/mips64/g90updatefw: main.go
	GOOS=linux ARCH=mips64 go build $(FLAGS) -o linux/mips64/g90updatefw

linux/mips64le/g90updatefw: main.go
	GOOS=linux ARCH=mips64le go build $(FLAGS) -o linux/mips64le/g90updatefw

linux/mipsle/g90updatefw: main.go
	GOOS=linux ARCH=mipsle go build $(FLAGS) -o linux/mipsle/g90updatefw

linux/ppc64/g90updatefw: main.go
	GOOS=linux ARCH=ppc64 go build $(FLAGS) -o linux/ppc64/g90updatefw

linux/ppc64le/g90updatefw: main.go
	GOOS=linux ARCH=ppc64le go build $(FLAGS) -o linux/ppc64le/g90updatefw

linux/riscv64/g90updatefw: main.go
	GOOS=linux ARCH=riscv64 go build $(FLAGS) -o linux/riscv64/g90updatefw

linux/s390x/g90updatefw: main.go
	GOOS=linux ARCH=s390x go build $(FLAGS) -o linux/s390x/g90updatefw
