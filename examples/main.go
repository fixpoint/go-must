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
