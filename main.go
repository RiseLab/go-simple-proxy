package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
)

func main() {
	host, port := "0.0.0.0", "8080"

	http.HandleFunc("/", proxyHandler)

	log.Printf("Server running at %s:%s\n", host, port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil); err != nil {
		log.Fatal(err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("New request: \"%s\"\n", r.URL)

	target := r.URL.Query().Get("t")

	if ok, err := isValidURI(target); ok {
		log.Printf("Proxy target is \"%s\"\n", target)

		res, err := http.Get(target)
		if err != nil {
			log.Printf("[ERROR] %s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Printf("[ERROR] %s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
		w.WriteHeader(res.StatusCode)
		fmt.Fprint(w, string(body))
	} else {
		log.Printf("[ERROR] %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func isValidURI(path string) (bool, error) {
	u, err := url.ParseRequestURI(path)

	if err != nil {
		return false, err
	}

	switch u.Scheme {
	case "http":
	case "https":
	default:
		return false, errors.New("invalid scheme")
	}

	_, err = net.LookupHost(u.Host)
	if err != nil {
		return false, err
	}

	return true, nil
}
