package medusasdkgolang

import "net/http"

type Config struct {
	MaxRetries int8         // The amount of times a request is retried.
	BaseUrl    string       // The url to which requests are made to.
	ApiKey     string       // For request on authorized routes in admin
	Cookie     *http.Cookie // For making authorized request using cookie "connect.id"
}

func NewConfig() *Config {
	c := new(Config)
	// add default value
	c.MaxRetries = 0
	c.BaseUrl = "http://localhost:9000"
	return c
}

func (c *Config) SetMaxRetries(maxRetries int8) *Config {
	c.MaxRetries = maxRetries
	return c
}

func (c *Config) SetBaseUrl(baseUrl string) *Config {
	c.BaseUrl = baseUrl
	return c
}

func (c *Config) SetApiKey(apiKey string) *Config {
	c.ApiKey = apiKey
	return c
}

func (c *Config) SetCookie(cookie *http.Cookie) *Config {
	c.Cookie = cookie
	return c
}
