all: test

test:
	rot13 < example.txt
	rot13 < example.txt | rot13

govet:
	go vet -v

gofmt:
	gofmt -s -w .

lint: govet gofmt
