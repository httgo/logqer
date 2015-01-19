package logqer

import (
	"net/http"
)

// Responsed provides an interface to retrieve status and headers from a
// ResponseWriter
type Responsed interface {
	Status() int

	// This is already built into ResponseWriter
	Header() http.Header
}

// responseWriter wraps a ResponseWriter to grab some of it's particulars
// Implements ResponseWriter and satisfies the Responsed interface
type responseWriter struct {
	http.ResponseWriter

	status int
}

// WriteHeader calls WriteHeader but saves the status to our responseWriter
func (r *responseWriter) WriteHeader(status int) {
	r.status = status

	r.ResponseWriter.WriteHeader(r.status)
}

// Status returns the status of the response.
// Because we are not intercepting the write function which calls WriteHeader
// when no explicit call to WriteHeader is given before writing. We just return
// status 0 as 200.
func (r responseWriter) Status() int {
	if r.status == 0 {
		return 200
	}

	return r.status
}

type LogFunc func(Responsed, *http.Request)

// Handler wraps a Handler and calls the LogFunc once the Handler has been
// Handled
func Handler(h http.Handler, logfn LogFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
		}

		h.ServeHTTP(rw, req)
		logfn(rw, req)
	})
}
