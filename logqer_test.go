package logqer

import (
	"fmt"
	"gopkg.in/nowk/assert.v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOutputsLogOfRequest(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Authorization", "asecret")
		w.WriteHeader(404)
		w.Write([]byte("Hello World!"))
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	log := httptest.NewRecorder()
	Handler(h, func(r Responsed, req *http.Request) {
		log.Write([]byte(fmt.Sprintf("%s [%d] %s Authorization: %s", req.Method,
			r.Status(), req.URL, r.Header().Get("Authorization"))))
	}).ServeHTTP(w, req)

	assert.Equal(t, "GET [404] /foo Authorization: asecret", log.Body.String())
}

func TestNoExplicitWriteHeader(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	log := httptest.NewRecorder()
	Handler(h, func(r Responsed, req *http.Request) {
		log.Write([]byte(fmt.Sprintf("%s [%d] %s", req.Method, r.Status(),
			req.URL)))
	}).ServeHTTP(w, req)

	assert.Equal(t, "GET [200] /foo", log.Body.String())
}
