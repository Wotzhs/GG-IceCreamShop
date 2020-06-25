package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"google.golang.org/grpc/metadata"
)

var (
	passwordCensorRe *regexp.Regexp
)

func init() {
	passwordCensorRe = regexp.MustCompile(`("?password"?:)(".*?")`)
}

func SetMiddlewares(mux http.Handler) http.Handler {
	mux = InjectAuthTokenPlaceholderCtx(mux)
	mux = TimeRequestCompletion(mux)
	mux = SetAuthMetadataIfAny(mux)
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

func InjectAuthTokenPlaceholderCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var placeholder string
		ctx := context.WithValue(r.Context(), "authtoken", &placeholder)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func SetAuthMetadataIfAny(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("authtoken")
		if err != nil {
			log.Printf("get authtoken from cookie err: %v", err)
		}

		if token != nil {
			jwt, err := jws.ParseJWT([]byte(token.Value))
			if err != nil {
				log.Printf("jwt.ParseJWT() err: %v", err)
			}

			if err := jwt.Validate([]byte("zalora"), crypto.SigningMethodHS256); err != nil {
				log.Printf("jwt.Validate() err: %v", err)
			}

			email, ok := jwt.Claims().Get("email").(string)
			if !ok || email == "" {
				log.Printf("email is not found in jwt")
			}

			ctx := metadata.AppendToOutgoingContext(r.Context(), "email", email)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
