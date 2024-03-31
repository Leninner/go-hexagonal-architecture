package core

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/leninner/go-feature-flags/pkg/core"
)

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pr, pw := io.Pipe()
		tee := io.TeeReader(r.Body, pw)
		r.Body = pr

		go func() {
			body, _ := io.ReadAll(tee)
			log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, string(body))
		}()

		next.ServeHTTP(w, r)
	})
}

func ApplicationHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func HandleExceptionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(core.NewResponseMessage(r.(string), nil))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
