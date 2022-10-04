package medusasdkgolang

type Medusa struct {
	MaxRetries int8   // The amount of times a request is retried.
	BaseUrl    string // The url to which requests are made to.
	ApiKey     string //Optional api key used for authenticating admin requests .

}

func NewMedusa() *Medusa {
	m := new(Medusa)
	return m
}
