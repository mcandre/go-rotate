# go-rotate - ROT13 lib and command line utility

# EXAMPLES

```
$ rot13 < example.txt
Cebsrffbe ZpTbantnyy ghearq vagb n png.

$ rot13 < example.txt | rot13
Professor McGonagall turned into a cat.
```

# DOWNLOAD

https://github.com/mcandre/go-rotate/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/go-rotate

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [opennota/check](https://github.com/opennota/check) (e.g. `go get github.com/opennota/check/cmd/...`)
* [megacheck](https://github.com/dominikh/go-tools/tree/master/cmd/megacheck) (e.g. `go get github.com/dominikh/go-tools/cmd/megacheck`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [editorconfig-cli](https://github.com/amyboyd/editorconfig-cli) (e.g. `go get github.com/amyboyd/editorconfig-cli`)
* [flcl](https://github.com/mcandre/flcl) (e.g. `go get github.com/mcandre/flcl/...`)
* [bashate](https://github.com/openstack-dev/bashate)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/go-rotate/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/go-rotate.git $GOPATH/src/github.com/mcandre/go-rotate
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-rotate/cmd/rot13 && go install"
```

# TEST REMOTELY

```
$ go test github.com/mcandre/go-rotate/...
```

# TEST LOCALLY

```
$ go test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
