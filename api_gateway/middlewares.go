package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var (
	passwordCensorRe *regexp.Regexp
)

func init() {
	passwordCensorRe = regexp.MustCompile(`("?password"?:)(".*?")`)
}

func SetMiddlewares(mux http.Handler) http.Handler {
	mux = TimeRequestCompletion(mux)
	mux = LogRequest(mux)
	return mux
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Printf("%v %v %v", r.Method, r.URL.Path, r.URL.RawQuery)
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("%v failed to read body %v", r.URL.Path, err)
			}

			var bodyCompact bytes.Buffer
			if err := json.Compact(&bodyCompact, body); err != nil {
				log.Printf("%v failed to compact body %v", r.URL.Path, err)
			}

			censoredBody := passwordCensorRe.ReplaceAllString(bodyCompact.String(), `$1"****"`)

			log.Printf("%v %v %v", r.Method, r.URL.Path, censoredBody)

			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		next.ServeHTTP(w, r)
	})
}

func TimeRequestCompletion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%v %v completed in %.2fms", r.Method, r.URL.Path, float64(time.Now().Sub(start).Nanoseconds())/float64(time.Millisecond))
	})
}
