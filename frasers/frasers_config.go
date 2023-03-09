package frasers

type FrasersClientConfig interface {
	FrasersBaseURL() string
	FrasersAPIKey() string
	FrasersPropertyBaseURL() string
	FrasersPropertyGrantTypeClient() string
	FrasersPropertyClientID() string
	FrasersPropertyClientSecretKey() string
}
