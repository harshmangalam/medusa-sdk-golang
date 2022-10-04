package medusasdkgolang

type Config struct {
	MaxRetries int8   // The amount of times a request is retried.
	BaseUrl    string // The url to which requests are made to.
	ApiKey     string // For request on authorized routes
}

func NewConfig() *Config {
	m := new(Config)
	// add default value
	m.MaxRetries = 0
	m.BaseUrl = "http://localhost:9000"
	return m
}

func (m *Config) SetMaxRetries(maxRetries int8) *Config {
	m.MaxRetries = maxRetries
	return m
}

func (m *Config) SetBaseUrl(baseUrl string) *Config {
	m.BaseUrl = baseUrl
	return m
}

func (m *Config) SetApiKey(apiKey string) *Config {
	m.ApiKey = apiKey
	return m
}
