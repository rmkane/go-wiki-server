package middleware

import (
	"log"
	"net/http"
	"regexp"
)

var (
	validPath  = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	validTitle = regexp.MustCompile("^[a-zA-Z0-9]+$")
)

// MakeHandler extracts the title from the URL and validates it.
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil || !isValidTitle(m[2]) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Request logging middleware
func logRequest(r *http.Request) {
	log.Printf("[%s] %s %s", r.RemoteAddr, r.Method, r.URL.Path)
}

// Ensure title validity
func isValidTitle(title string) bool {
	return validTitle.MatchString(title)
}
