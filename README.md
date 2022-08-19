# go-tenable

[Go](https://golang.org/) client library for [Tenable.io](https://tenable.io).


## Features

* Authentication (API Key)
* Retrieve Repositories, Analysis.

## Requirements

* Go >= 1.18
* Tenable ??

## Installation

It is go gettable

```bash
go get github.com/IBM/go-tenable
```

Usage:

```go
package main

import (
	tenable "gopkg.in/IBM/go-tenable.v1"
)
...
```

(optional) to run unit / example tests:

```bash
cd $GOPATH/src/github.com/IBM/go-tenable
go test -v ./...
```
## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/IBM/go-tenable) for a detailed API description.

The [latest Tenable REST API documentation](https://docs.tenable.com/tenablesc/api/index.htm) was the base document for this package.


