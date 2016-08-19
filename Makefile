.PHONY: dist

build: fmt
	go build -o ./dist/gist ./cmd/gist

fmt: 
	go fmt ./...

test:
	go test ./...

dist:
	gox -osarch="linux/amd64 darwin/amd64" ./cmd/gist && mv gist_darwin_amd64 ./dist/gist_darwin_amd64 && mv gist_linux_amd64 ./dist/gist_linux_amd64
