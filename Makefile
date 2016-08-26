build: fmt
	go build ./cmd/gist

fmt: 
	go fmt ./...

test:
	go test ./...

bin:
	gox -osarch="linux/386 linux/amd64 linux/arm darwin/amd64 windows/386 windows/amd64" ./cmd/gist
