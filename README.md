# go-rotate - ROT13 lib and command line utility

# EXAMPLES

```console
$ rot13 < example.txt
Cebsrffbe ZpTbantnyy ghearq vagb n png.

$ rot13 < example.txt | rot13
Professor McGonagall turned into a cat.
```

# DOWNLOAD

https://github.com/mcandre/go-rotate/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/go-rotate

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/go-rotate/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/go-rotate.git $GOPATH/src/github.com/mcandre/go-rotate
$ cd "$GOPATH/src/github.com/mcandre/go-rotate"
$ git submodule update --init --recursive
$ go install ./...
```

# TEST

```console
$ mage test
```

# PORT

```console
$ mage port
```

# LINT

Keep the code tidy:

```console
$ mage lint
```
