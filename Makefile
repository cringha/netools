BIN_FILE=nettools

MAIN=main.go


all: build run

build: clean
	CGO_ENABLED=1  go build -ldflags "-s -w"   -o ${BIN_FILE} ${MAIN}
	@echo build ${BIN_FILE}

test:
	CGO_ENABLED=1 go test ./... -gcflags=-l

run:
	./${BIN_FILE} 


clean:
	rm -rf ${BIN_FILE}

