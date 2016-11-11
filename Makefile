all: test

test:
	rot13 < example.txt
	rot13 < example.txt | rot13

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

lint: govet gofmt
