package util

// Ignore ignores a single error.
func Ignore(fn func() error) {
	_ = fn()
}
