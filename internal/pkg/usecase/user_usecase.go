package usecase

import (
	"context"

	"github.com/hanifmaliki/go-boilerplate/internal/pkg/entity"
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/model"
	"github.com/hanifmaliki/go-boilerplate/internal/pkg/repository"

	"github.com/jinzhu/copier"
)

type UserUseCase interface {
	GetUser(ctx context.Context) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error)
	UpdateUser(ctx context.Context) (*entity.User, error)
	DeleteUser(ctx context.Context) error
}

func (u *useCase) GetUser(ctx context.Context) (*entity.User, error) {
	data := &entity.User{}
	return data, nil
}

func (u *useCase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return nil, nil
}

func (u *useCase) CreateUser(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error) {
	data := &entity.User{}
	err := u.repository.InTransaction(ctx, func(r repository.Repository) error {
		userRepo := r.UserRepo()
		addressRepo := r.AddressRepo()
		ccRepo := r.CreditCardRepo()
		userRoleRepo := r.UserRoleRepo()

		user := &entity.User{}
		copier.Copy(user, request)
		_, err := userRepo.Create(ctx, user)
		if err != nil {
			return err
		}

		address := &entity.Address{}
		copier.Copy(address, request.Address)
		_, err = addressRepo.Create(ctx, address)
		if err != nil {
			return err
		}

		ccs := []*entity.CreditCard{}
		for _, ccRequest := range request.CreditCards {
			cc := &entity.CreditCard{}
			copier.Copy(cc, ccRequest)
			_, err := ccRepo.Create(ctx, cc)
			if err != nil {
				return err
			}
			ccs = append(ccs, cc)
		}

		userRoles := []*entity.UserRole{}
		for _, urRequest := range request.UserRoles {
			userRole, err := userRoleRepo.Create(ctx, &entity.UserRole{
				UserID: user.ID,
				RoleID: urRequest.RoleID,
			})
			if err != nil {
				return err
			}
			userRoles = append(userRoles, userRole)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *useCase) UpdateUser(ctx context.Context) (*entity.User, error) {
	data := &entity.User{}
	return data, nil
}

func (u *useCase) DeleteUser(ctx context.Context) error {
	return nil
}
