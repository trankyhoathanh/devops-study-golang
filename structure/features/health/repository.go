package health

import (
	"context"
)

type Repository interface {
	CheckHealth(ctx context.Context) (error)
}