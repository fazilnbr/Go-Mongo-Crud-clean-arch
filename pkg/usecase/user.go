package usecase

import (
	"context"

	domain "github.com/SethukumarJ/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/SethukumarJ/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/SethukumarJ/go-gin-clean-arch/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// Delete implements interfaces.UserUseCase
func (*userUseCase) Delete(ctx context.Context, user domain.Users) error {
	panic("unimplemented")
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.UserResponse, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

// func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
// 	err := c.userRepo.Delete(ctx, user)

// 	return err
// }
