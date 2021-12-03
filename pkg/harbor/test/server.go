package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/goharbor/go-client/pkg/sdk/assist/client/version"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

const (
	secret = "Basic Yjph"
	result = "yes"
	v      = "1.2"
)

func NewServer() *httptest.Server {
	mux := http.NewServeMux()

	// v2 server
	mux.HandleFunc("/api/v2.0/health", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != secret {
			fmt.Println(auth)
			http.Error(w, "Auth header was incorrect", http.StatusUnauthorized)
			return
		}
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Add("Content-Type", "application/json")

		h := &models.OverallHealthStatus{
			Status: result,
		}

		b, _ := json.Marshal(h)
		if _, err := w.Write(b); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	})

	// legacy server
	mux.HandleFunc("/api/v2.0/email/ping", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != secret {
			http.Error(w, "Auth header was incorrect", http.StatusUnauthorized)
			return
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	// assist server
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != secret {
			http.Error(w, "Auth header was incorrect", http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Add("Content-Type", "application/json")

		h := &version.GetVersionOKBody{
			Version: v,
		}

		b, _ := json.Marshal(h)
		if _, err := w.Write(b); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(mux)
}
