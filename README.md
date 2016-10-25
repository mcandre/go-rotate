# go-rotate - ROT13 lib and command line utility

# EXAMPLES

```
$ rot13 < example.txt
Cebsrffbe ZpTbantnyy ghearq vagb n png.

$ rot13 < example.txt | rot13
Professor McGonagall turned into a cat.
```

# REQUIREMENTS

* [Go](https://golang.org/) 1.7+

## Optional

* [make](https://www.gnu.org/software/make/)

# INSTALL

```
$ go get github.com/mcandre/go-rotate
$ go install github.com/mcandre/go-rotate/cmd/rot13
```

# COMPILE AND INSTALL LOCALLY

```
$ sh -c 'cd cmd/rot13 && go install'
```

# LINT

Keep the code tidy:

```
$ make lint
```
