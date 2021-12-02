package harbor_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/goharbor/go-client/pkg/harbor"
	assistclient "github.com/goharbor/go-client/pkg/sdk/assist/client"
	"github.com/goharbor/go-client/pkg/sdk/assist/client/version"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	legacyclient "github.com/goharbor/go-client/pkg/sdk/v2.0/legacy/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/legacy/client/products"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

var (
	httpsSchema = "https"
	httpSchema  = "http"
)

type testParam struct {
	url               string
	transport         http.RoundTripper
	authInfo          runtime.ClientAuthInfoWriter
	expectedHost      string
	expectedScheme    string
	expectedBasePath  string
	expectedTransport http.RoundTripper
	expectedAuthInfo  runtime.ClientAuthInfoWriter
}

var v2ClientParams = []testParam{
	testParam{
		url:               "//10.0.0.1:443",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1:443",
		expectedBasePath:  v2client.DefaultBasePath,
		expectedScheme:    httpsSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
	testParam{
		url:               "http://10.0.0.1",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1",
		expectedBasePath:  v2client.DefaultBasePath,
		expectedScheme:    httpSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
}

func TestToV2Config(t *testing.T) {
	for _, test := range v2ClientParams {
		u, err := url.Parse(test.url)
		if err != nil {
			t.Errorf("Failed to pase the url %s", test.url)
		}

		c := &harbor.Config{
			URL:       u,
			Transport: test.transport,
		}

		v2c := c.ToV2Config()

		if v2c.URL.Host != test.expectedHost {
			t.Errorf("Output Host %q not match expected URL %q", v2c.URL.Host, test.expectedHost)
		}
		if v2c.URL.Scheme != test.expectedScheme {
			t.Errorf("Output Scheme %q not match expected Scheme %q", v2c.URL.Scheme, test.expectedScheme)
		}
		if v2c.URL.Path != test.expectedBasePath {
			t.Errorf("Output BasePath %q not match expected BasePath %q", v2c.URL.Scheme, test.expectedBasePath)
		}
		if v2c.Transport != test.expectedTransport {
			t.Errorf("Output Transport %q not match expected Transport %q", v2c.Transport, test.expectedTransport)
		}
		if v2c.AuthInfo != test.expectedAuthInfo {
			t.Errorf("Output AuthInfo %q not match expected AuthInfo %q", v2c.AuthInfo, test.expectedAuthInfo)
		}
	}
}

var assistClientParams = []testParam{
	testParam{
		url:               "//10.0.0.1:443",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1:443",
		expectedBasePath:  assistclient.DefaultBasePath,
		expectedScheme:    httpsSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
	testParam{
		url:               "http://10.0.0.1",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1",
		expectedBasePath:  assistclient.DefaultBasePath,
		expectedScheme:    httpSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
}

func TestToAssistConfig(t *testing.T) {
	for _, test := range assistClientParams {
		u, err := url.Parse(test.url)
		if err != nil {
			t.Errorf("Failed to pase the url %s", test.url)
		}

		c := &harbor.Config{
			URL:       u,
			Transport: test.transport,
		}

		v2c := c.ToAssistConfig()

		if v2c.URL.Host != test.expectedHost {
			t.Errorf("Output Host %q not match expected URL %q", v2c.URL.Host, test.expectedHost)
		}
		if v2c.URL.Scheme != test.expectedScheme {
			t.Errorf("Output Scheme %q not match expected Scheme %q", v2c.URL.Scheme, test.expectedScheme)
		}
		if v2c.URL.Path != test.expectedBasePath {
			t.Errorf("Output BasePath %q not match expected BasePath %q", v2c.URL.Scheme, test.expectedBasePath)
		}
		if v2c.Transport != test.expectedTransport {
			t.Errorf("Output Transport %q not match expected Transport %q", v2c.Transport, test.expectedTransport)
		}
		if v2c.AuthInfo != test.expectedAuthInfo {
			t.Errorf("Output AuthInfo %q not match expected AuthInfo %q", v2c.AuthInfo, test.expectedAuthInfo)
		}
	}
}

var legacyClientParams = []testParam{
	testParam{
		url:               "//10.0.0.1:443",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1:443",
		expectedBasePath:  legacyclient.DefaultBasePath,
		expectedScheme:    httpsSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
	testParam{
		url:               "http://10.0.0.1",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1",
		expectedBasePath:  legacyclient.DefaultBasePath,
		expectedScheme:    httpSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
	},
}

func TestToLegacyConfig(t *testing.T) {
	for _, test := range legacyClientParams {
		u, err := url.Parse(test.url)
		if err != nil {
			t.Errorf("Failed to pase the url %s", test.url)
		}

		c := &harbor.Config{
			URL:       u,
			Transport: test.transport,
		}

		v2c := c.ToLegacyConfig()

		if v2c.URL.Host != test.expectedHost {
			t.Errorf("Output Host %q not match expected URL %q", v2c.URL.Host, test.expectedHost)
		}
		if v2c.URL.Scheme != test.expectedScheme {
			t.Errorf("Output Scheme %q not match expected Scheme %q", v2c.URL.Scheme, test.expectedScheme)
		}
		if v2c.URL.Path != test.expectedBasePath {
			t.Errorf("Output BasePath %q not match expected BasePath %q", v2c.URL.Scheme, test.expectedBasePath)
		}
		if v2c.Transport != test.expectedTransport {
			t.Errorf("Output Transport %q not match expected Transport %q", v2c.Transport, test.expectedTransport)
		}
		if v2c.AuthInfo != test.expectedAuthInfo {
			t.Errorf("Output AuthInfo %q not match expected AuthInfo %q", v2c.AuthInfo, test.expectedAuthInfo)
		}
	}
}

var clientSetParams = []testParam{
	testParam{
		url:       "//10.0.0.1:443",
		transport: harbor.InsecureTransport,
		authInfo:  nil,
	},
}

func TestClientSet(t *testing.T) {
	for _, test := range clientSetParams {

		c := &harbor.ClientSetConfig{
			URL:      test.url,
			Password: "a",
			Username: "b",
			Insecure: true,
		}

		cs, err := harbor.NewClientSet(c)
		if cs == nil {
			t.Errorf("ClientSet shouldn't be empty")
		}
		if err != nil {
			t.Errorf("err when creating the client %v", err)
		}

		assistClient := cs.Assist()
		if assistClient == nil {
			t.Errorf("assistClient shouldn't be empty")
		}

		legacyClient := cs.Legacy()
		if legacyClient == nil {
			t.Errorf("legacyClient shouldn't be empty")
		}
		v2Client := cs.V2()
		if v2Client == nil {
			t.Errorf("v2Client shouldn't be empty")
		}
	}
}

func TestHarborV2Client(t *testing.T) {
	secret := "Basic Yjph"
	result := "yes"

	mux := http.NewServeMux()
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

	s := httptest.NewServer(mux)
	defer s.Close()

	c := &harbor.ClientSetConfig{
		URL:      s.URL,
		Password: "a",
		Username: "b",
		Insecure: true,
	}

	cs, err := harbor.NewClientSet(c)
	if err != nil {
		t.Errorf("error when creating the client %v", err)
	}

	res, err := cs.V2().Health.GetHealth(context.TODO(), health.NewGetHealthParams())
	if err != nil {
		t.Errorf("error when calling endpoint %v", err)
	}

	if res.Payload.Status != result {
		t.Errorf("status '%s' doesn't match '%s'", res.Payload.Status, result)
	}
}

func TestHarborLegacyClient(t *testing.T) {
	secret := "Basic Yjph"

	mux := http.NewServeMux()
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

	s := httptest.NewServer(mux)
	defer s.Close()

	c := &harbor.ClientSetConfig{
		URL:      s.URL,
		Password: "a",
		Username: "b",
		Insecure: true,
	}

	cs, err := harbor.NewClientSet(c)
	if err != nil {
		t.Errorf("error when creating the client %v", err)
	}

	_, err = cs.Legacy().Products.PostEmailPing(context.TODO(), products.NewPostEmailPingParams())
	if err != nil {
		t.Errorf("error when calling endpoint %v", err)
	}
}

func TestHarborAssistClient(t *testing.T) {
	secret := "Basic Yjph"
	v := "1.2"

	mux := http.NewServeMux()
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

	s := httptest.NewServer(mux)
	defer s.Close()

	c := &harbor.ClientSetConfig{
		URL:      s.URL,
		Password: "a",
		Username: "b",
		Insecure: true,
	}

	cs, err := harbor.NewClientSet(c)
	if err != nil {
		t.Errorf("error when creating the client %v", err)
	}

	res, err := cs.Assist().Version.GetVersion(context.TODO(), version.NewGetVersionParams())
	if err != nil {
		t.Errorf("error when calling endpoint %v", err)
	}

	if res.Payload.Version != v {
		t.Errorf("version '%s' doesn't match '%s'", res.Payload.Version, v)
	}
}
