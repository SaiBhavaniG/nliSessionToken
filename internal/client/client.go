package client

import (
	"context"
)

type TokenClient interface {
	SendToken(ctx context.Context, token string) error
}
