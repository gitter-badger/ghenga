package main

// CleanupErr runs fn and sets err to the returned error if err is nil.
func CleanupErr(err *error, fn func() error) {
	e := fn()
	if *err == nil {
		*err = e
	}
}
