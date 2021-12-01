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

// InsecureTransport provides a insecure RoundTripper and disable the HTTP2 try.
var InsecureTransport http.RoundTripper = &http.Transport{
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

// Config contains configs for constructing a client
type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// ClientSet contains clients for V2, Assist, Legacy
type ClientSet struct {
	AssistClient *assistclient.HarborAPI
	V2Client     *v2client.HarborAPI
	LegacyClient *legacyclient.HarborAPI
}

// ToAssistConfig convert the Config to assistclient's Config
func (c *Config) ToAssistConfig() assistclient.Config {
	if c.URL == nil {
		c.URL = &url.URL{}
	}

	if len(c.URL.Host) == 0 {
		c.URL.Host = assistclient.DefaultHost
	}
	if len(c.URL.Path) == 0 {
		c.URL.Path = assistclient.DefaultBasePath
	}
	if len(c.URL.Scheme) == 0 {
		// default to https if not set
		c.URL.Scheme = assistclient.DefaultSchemes[1]
	}

	if c.Transport == nil {
		c.Transport = InsecureTransport
	}

	return assistclient.Config{
		URL:       c.URL,
		Transport: c.Transport,
		AuthInfo:  c.AuthInfo,
	}
}

// ToV2Config convert the Config to v2client's Config
func (c *Config) ToV2Config() v2client.Config {
	if c.URL == nil {
		c.URL = &url.URL{}
	}

	if len(c.URL.Host) == 0 {
		c.URL.Host = v2client.DefaultHost
	}
	if len(c.URL.Path) == 0 {
		c.URL.Path = v2client.DefaultBasePath
	}
	if len(c.URL.Scheme) == 0 {
		// default to https if not set
		c.URL.Scheme = v2client.DefaultSchemes[1]
	}

	if c.Transport == nil {
		c.Transport = InsecureTransport
	}

	return v2client.Config{
		URL:       c.URL,
		Transport: c.Transport,
		AuthInfo:  c.AuthInfo,
	}
}

// ToLegacyConfig convert the Config to legacyclient's Config
func (c *Config) ToLegacyConfig() legacyclient.Config {
	if c.URL == nil {
		c.URL = &url.URL{}
	}

	if len(c.URL.Host) == 0 {
		c.URL.Host = legacyclient.DefaultHost
	}
	if len(c.URL.Path) == 0 {
		c.URL.Path = legacyclient.DefaultBasePath
	}
	if len(c.URL.Scheme) == 0 {
		// default to https if not set
		c.URL.Scheme = legacyclient.DefaultSchemes[1]
	}

	if c.Transport == nil {
		c.Transport = InsecureTransport
	}

	return legacyclient.Config{
		URL:       c.URL,
		Transport: c.Transport,
		AuthInfo:  c.AuthInfo,
	}
}

// NewAssistClient return assist Client
func NewClientSet(c *Config) *ClientSet {
	if c == nil {
		c = &Config{}
	}

	cs := &ClientSet{}
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
