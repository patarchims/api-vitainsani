package entity

import (
	"context"
	"vincentcoreapi/modules/user"
)

// UserUseCase
type UserUseCase interface {
	GetByUser(ctx context.Context, userName string) (res user.ApiUser, exist bool)
}

// UserRepository
type UserRepository interface {
	GetByID(ctx context.Context, userID string) (res user.ApiUser, exist bool)
	GetByUser(ctx context.Context, userName string) (res user.ApiUser, exist bool)
}
