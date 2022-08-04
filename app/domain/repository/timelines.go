package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timelines interface {
	GetPublicTimelines(ctx context.Context, max_id string, since_id string, limit string) ([]*object.Status, error)
}
