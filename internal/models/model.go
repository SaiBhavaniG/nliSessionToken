package models

import "time"

type GetTokenRequest struct {
	DeviceID string
}
type GetTokenResponse struct {
	DeviceID  string
	NLIToken  string
	CreatedAt time.Time
}
