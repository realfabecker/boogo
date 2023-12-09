.PHONY: build
build: clean build-linux

.PHONY: clean
clean:
	$(info INFO: cleaning previous assets)
	rm -rf ./out

.PHONY: build-linux
build-linux:
	$(info INFO: compiling linux binary)
	GOOS=linux GOARCH=amd64 go build -o ./out/bogo ./cmd/bogo/bogo.go
