.PHONY: dist

build: fmt
	go build -o ./dist/gist ./cmd/gist

fmt: 
	go fmt ./...

test:
	go test ./...

dist:
	gox -osarch="linux/386 linux/amd64 linux/arm darwin/amd64 windows/386 windows/amd64" ./cmd/gist && mv gist_* ./dist/
