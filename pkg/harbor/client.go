package harbor

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/runtime"
	assistclient "github.com/goharbor/go-client/pkg/sdk/assist/client"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	legacyclient "github.com/goharbor/go-client/pkg/sdk/v2.0/legacy/client"
)

// insecureTransport provides a insecure RoundTripper and disable the HTTP2 try.
var insecureTransport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // nolint:gomnd
		KeepAlive: 30 * time.Second, // nolint:gomnd
		DualStack: true,
	}).DialContext,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true, // nolint:gosec
	},
	MaxIdleConns:          100,              // nolint:gomnd
	IdleConnTimeout:       90 * time.Second, // nolint:gomnd
	TLSHandshakeTimeout:   10 * time.Second, // nolint:gomnd
	ExpectContinueTimeout: 1 * time.Second,
}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

type ClientSet struct {
	AssistClient *assistclient.HarborAPI
	V2Client     *v2client.HarborAPI
	LegacyClient *legacyclient.HarborAPI
}

func (c *Config) ToAssistConfig() assistclient.Config {
	return assistclient.Config{
		URL:       c.URL,
		Transport: insecureTransport,
		AuthInfo:  c.AuthInfo,
	}
}

func (c *Config) ToV2Config() v2client.Config {
	return v2client.Config{
		URL:       c.URL,
		Transport: insecureTransport,
		AuthInfo:  c.AuthInfo,
	}
}
func (c *Config) ToLegacyConfig() legacyclient.Config {
	return legacyclient.Config{
		URL:       c.URL,
		Transport: insecureTransport,
		AuthInfo:  c.AuthInfo,
	}
}

// NewAssistClient return assist Client
func NewClientSet(c *Config) *ClientSet {
	if c == nil {
		return nil
	}

	cs := &ClientSet{}

	// opinionated config
	c.URL.Scheme = "https"
	c.Transport = insecureTransport

	cs.AssistClient = assistclient.New(c.ToAssistConfig())
	cs.LegacyClient = legacyclient.New(c.ToLegacyConfig())
	cs.V2Client = v2client.New(c.ToV2Config())

	return cs
}

// Assist return AssistClient
func (c *ClientSet) Assist() *assistclient.HarborAPI {
	return c.AssistClient
}

// V2 return V2Client
func (c *ClientSet) V2() *v2client.HarborAPI {
	return c.V2Client
}

// Legacy return LegacyClient
func (c *ClientSet) Legacy() *legacyclient.HarborAPI {
	return c.LegacyClient
}
