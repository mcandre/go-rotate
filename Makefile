all: test

test:
	rot13 < example.txt
	rot13 < example.txt | rot13

gofmt:
	gofmt -s -w .

lint: gofmt
