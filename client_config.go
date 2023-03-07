package frasers

type FrasersClientConfig interface {
	FrasersBaseURL() string
	FrasersAPIKey() string
}
