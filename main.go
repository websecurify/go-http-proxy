package main // import "websecurify/go-http-proxy"

// ---
// ---
// ---

import (
	"os"
	"log"
	"path"
	"strings"
	"net/url"
	"net/http"
	"net/http/httputil"
	
	// ---
	
	"github.com/abbot/go-http-auth"
)

// ---
// ---
// ---

func main() {
	funcWrapper := func (f http.HandlerFunc) (http.HandlerFunc) {
		return f
	}
	
	// ---
	
	if os.Getenv("HTPASSWD") != "" {
		secrets := auth.HtpasswdFileProvider(os.Getenv("HTPASSWD"))
		
		// ---
		
		authenticator := auth.NewBasicAuthenticator(os.Getenv("REALM"), secrets)
		
		// ---
		
		funcWrapper = func (f http.HandlerFunc) (http.HandlerFunc) {
			return auth.JustCheck(authenticator, f)
		}
	}
	
	// ---
	
	backend := os.Getenv("BACKEND")
	
	// ---
	
	log.Println("mapping to", backend)
	
	// ---
	
	backendURL, backendURLErr := url.Parse(backend)
	
	if backendURLErr != nil {
		log.Fatal(backendURLErr)
	}
	
	// ---
	
	proxy := &httputil.ReverseProxy{Director: func (r *http.Request) {
		endsWithSlash := strings.HasSuffix(r.URL.Path, "/")
		
		// ---
		
		r.URL.Scheme = backendURL.Scheme
		r.URL.Host = backendURL.Host
		r.URL.Path = path.Join(backendURL.Path, r.URL.Path)
		
		// ---
		
		if endsWithSlash && !strings.HasSuffix(r.URL.Path, "/") {
			r.URL.Path = r.URL.Path + "/"
		}
	}}
	
	// ---
	
	http.HandleFunc("/", funcWrapper(func (w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}))
	
	// ---
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
