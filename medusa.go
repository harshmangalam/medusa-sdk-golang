package medusasdkgolang

type Medusa struct {
	MaxRetries int8   // The amount of times a request is retried.
	BaseUrl    string // The url to which requests are made to.
	ApiKey     string //Optional api key used for authenticating admin requests .

}

func NewMedusa() *Medusa {
	m := new(Medusa)
	// add default value
	m.MaxRetries = 0
	m.BaseUrl = "http://localhost:9000"
	return m
}

func (m *Medusa) AddMaxRetries(maxRetries int8) *Medusa {
	m.MaxRetries = maxRetries
	return m
}
