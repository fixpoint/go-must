# go-must

[![Test](https://github.com/fixpoint/go-must/actions/workflows/test.yml/badge.svg)](https://github.com/fixpoint/go-must/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fixpoint/go-must)](https://goreportcard.com/report/github.com/fixpoint/go-must)
[![GoDoc](https://godoc.org/github.com/fixpoint/go-must?status.svg)](https://godoc.org/github.com/fixpoint/go-must)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/fixpoint/go-must/blob/main/LICENSE)

**go-must** is a package to provide `Must[T any](v T, err error) T` for convenience.

## Usage

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/fixpoint/go-must"
)

func main() {
	u := must.Must(url.Parse("https://example.com"))
	fmt.Printf("%v\n", u)
}
```
