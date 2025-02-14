package interfaces

import (
	"context"

	"github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.UserResponse, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
	
}
