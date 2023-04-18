package must_test

import (
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/fixpoint/go-must"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("Must returns `v` when `err` is `nil`", func(t *testing.T) {
		assert := assert.New(t)
		fn := func() (string, error) {
			return "success", nil
		}
		r := must.Must(fn())
		assert.Equal("success", r)
	})
	t.Run("Must panics with `err` when `err` is not `nil`", func(t *testing.T) {
		assert := assert.New(t)
		fn := func() (string, error) {
			return "hello", errors.New("fail")
		}
		assert.Panics(func() {
			_ = must.Must(fn())
		})
	})
}

func ExampleMust() {
	// url.Parse() must not be failed in this case
	u := must.Must(url.Parse("https://example.com"))
	fmt.Printf("%s\n", u)

	// Output:
	// https://example.com
}
