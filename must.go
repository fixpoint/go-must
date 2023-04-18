package must

// Must returns `v` as-is when `err` is not nil. Otherwise it panics with `err`.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
