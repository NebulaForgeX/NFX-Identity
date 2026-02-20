package enums

type ClientsApiKeyStatus string

const (
	ClientsApiKeyStatusActive  ClientsApiKeyStatus = "active"
	ClientsApiKeyStatusRevoked ClientsApiKeyStatus = "revoked"
	ClientsApiKeyStatusExpired ClientsApiKeyStatus = "expired"
)
