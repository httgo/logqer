# logqer

[![Build Status](https://travis-ci.org/httgo/logqer.svg?branch=master)](https://travis-ci.org/httgo/logqer)
[![GoDoc](https://godoc.org/gopkg.in/httgo/logqer.v0?status.svg)](http://godoc.org/gopkg.in/httgo/logqer.v0)

A simple HTTP request logger.

## Install

    go get gopkg.in/httgo/logqer.v0

## Example

    mux := http.NewServerMux()
    ...

    var reqlog = func(r logqer.Responsed, req *http.Request) {
      log.Printf("%-7s [%d] %s\n", req.Method, r.Status(), req.URL)
    }

    h := logqer.Handler(mux)
    ...

Will output something similar to:

    2015/01/18 21:47:21 GET     [200] /v1/timers/9dfce24?location=America%2FNew_York&:id=9dfce24
    2015/01/18 21:47:37 OPTIONS [200] /v1/timers
    2015/01/18 21:47:37 POST    [200] /v1/timers
    2015/01/18 21:47:37 GET     [200] /v1/timers/fda6ae8?location=America%2FNew_York&:id=fda6ae8
    2015/01/18 21:47:40 GET     [400] /v1/timers/fda6ae8a?location=America%2FNew_York&:id=fda6ae8a

## License

MIT
