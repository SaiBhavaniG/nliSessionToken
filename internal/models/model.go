package model

type GetTokenRequest struct {
	DeviceID string
}
type GetTokenResponse struct {
	DeviceID  string
	NLIToken  string
	CreatedAt string
}
