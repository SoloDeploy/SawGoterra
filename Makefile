BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME}

test:
	go test -v ./...
