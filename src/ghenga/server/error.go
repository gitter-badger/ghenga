package server

// Error bundles an HTTP status code and an error.
type Error interface {
	error
	Status() int
}

// statusError bundles an HTTP status code with an error.
type StatusError struct {
	error
	status int
}

// Status returns the HTTP status for this error
func (err StatusError) Status() int {
	return err.status
}
