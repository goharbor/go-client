package harbor_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/harbor/test"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	"github.com/stretchr/testify/assert"
)

var (
	httpsSchema = "https"
	httpSchema  = "http"
)

type clientType string

var (
	v2Type clientType = "v2"
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
	clientType        clientType
}

var params = []testParam{
	testParam{
		url:               "//10.0.0.1:443",
		transport:         harbor.InsecureTransport,
		authInfo:          nil,
		expectedHost:      "10.0.0.1:443",
		expectedBasePath:  v2client.DefaultBasePath,
		expectedScheme:    httpsSchema,
		expectedTransport: harbor.InsecureTransport,
		expectedAuthInfo:  nil,
		clientType:        v2Type,
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
		clientType:        v2Type,
	},
}

func TestToConfig(t *testing.T) {
	for _, test := range params {
		assert := assert.New(t)
		u, err := url.Parse(test.url)
		assert.Nil(err)

		c := &harbor.Config{
			URL:       u,
			Transport: test.transport,
		}

		if test.clientType == v2Type {
			v2c := c.ToV2Config()
			assert.NotNil(v2c)
			assert.Equal(test.expectedHost, v2c.URL.Host)
			assert.Equal(test.expectedScheme, v2c.URL.Scheme)
			assert.Equal(test.expectedBasePath, v2c.URL.Path)
			assert.Equal(test.expectedTransport, v2c.Transport)
			assert.Equal(test.expectedAuthInfo, v2c.AuthInfo)
		} else {
			t.Errorf("Unexpected client type %s", test.clientType)
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
		assert := assert.New(t)

		c := &harbor.ClientSetConfig{
			URL:      test.url,
			Password: "a",
			Username: "b",
			Insecure: true,
		}

		cs, err := harbor.NewClientSet(c)
		assert.NotNil(cs)
		assert.Nil(err)

		v2Client := cs.V2()
		assert.NotNil(v2Client)
	}
}

var (
	result   = "yes"
	v        = "1.2"
	username = "a"
	password = "b"
)

func TestClients(t *testing.T) {
	assert := assert.New(t)

	s := test.NewServer()
	defer s.Close()

	c := &harbor.ClientSetConfig{
		URL:      s.URL,
		Password: username,
		Username: password,
		Insecure: true,
	}

	cs, err := harbor.NewClientSet(c)
	assert.Nil(err)

	// test v2 client
	resGetHealth, err := cs.V2().Health.GetHealth(context.TODO(), health.NewGetHealthParams())
	assert.Nil(err)

	assert.Equal(result, resGetHealth.Payload.Status)
}
