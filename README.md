# cronparse

This is a small utility for parsing cron expressions into structured data that can be presented to users in human-readable form

## Running

To run `cronparse`, supply cron expression in quotes on the command line:
```
$ cronparse "*/15 0 1,15 * 1-5 /usr/bin/find"
Cron Expression:
minute		0 15 30 45

hour		0

day of		1 15
month

month		1 2 3 4 5 6 7 8 9 10 11 12

day of week	1 2 3 4 5

command		/usr/bin/find

```

You can use `--version` to check the binary version in use:
```
cronparse version v0.1
```

## Building

### Installing Prerequisites

Building and running tests requires a few prerequisites:

* make
* go 1.12 or higher
* bash 4 or higher
* golangci-lint
* ginkgo & gomega

Once installed, run the supplied configure script to set up the build environment:
```
$ ./configure
```

#### make

On Mac you will need the XCode command line tools installed. To do this, run
```
$ xcode-select --install
```

On Linux, install `make` according to your Linux distribution.


#### bash 4
The default version of bash shipped with MacOS is very old (due to licensing reasons).

A newer version can be installed using Homebrew:
```
$ brew install bash
```


#### golangci-lint
The linter is not strictly required (and can be skipped by passing `--without-linter` to `./configure`) but is recommended for development.

To install `golangci-lint` on Mac, use Homebrew:
```
$ brew install golangci/tap/golangci-lint
```

On Linux, download it directly:
```
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.33.0

```


#### ginkgo & gomega
Ginkgo is a testing framework and associated test runner. The can be installed via `go get`:
```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```


### Building

To just build the binary, run
```
$ make build
```

In order to run full linting, checks and tests:
```
$ make check
```