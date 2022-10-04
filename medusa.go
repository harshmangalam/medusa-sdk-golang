package medusasdkgolang

import "net/http"

type Medusa struct {
	MaxRetries int8         // The amount of times a request is retried.
	BaseUrl    string       // The url to which requests are made to.
	Cookie     *http.Cookie //cookie connect.sid for authenticated request
}

func NewMedusa() *Medusa {
	m := new(Medusa)
	// add default value
	m.MaxRetries = 0
	m.BaseUrl = "http://localhost:9000"
	return m
}

func (m *Medusa) SetMaxRetries(maxRetries int8) *Medusa {
	m.MaxRetries = maxRetries
	return m
}

func (m *Medusa) SetBaseUrl(baseUrl string) *Medusa {
	m.BaseUrl = baseUrl
	return m
}

func (m *Medusa) SetCookie(cookie *http.Cookie) *Medusa {
	m.Cookie = cookie
	return m
}
