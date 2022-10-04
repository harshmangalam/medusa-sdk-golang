package medusasdkgolang

type Config struct {
	MaxRetries int8   // The amount of times a request is retried.
	BaseUrl    string // The url to which requests are made to.
	ApiKey     string // For request on authorized routes
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
