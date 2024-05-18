package usecase

import (
	"context"

	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/entity"
	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/model"
	"github.com/hanifmaliki/golang-boilerplate/internal/pkg/repository"
	pkg_model "github.com/hanifmaliki/golang-boilerplate/pkg/model"

	"github.com/jinzhu/copier"
)

type UserUseCase interface {
	GetUser(ctx context.Context, conds *entity.User, query *pkg_model.Query) (*entity.User, error)
	GetUsers(ctx context.Context, query pkg_model.Query) ([]*entity.User, error)
	CreateUser(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error)
	UpdateUser(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context) error
}

func (u *useCase) GetUser(ctx context.Context, conds *entity.User, query *pkg_model.Query) (*entity.User, error) {
	return u.repository.UserRepo().FindOne(ctx, conds, query)
}

func (u *useCase) GetUsers(ctx context.Context, query pkg_model.Query) ([]*entity.User, error) {
	datas := []*entity.User{}
	return datas, nil
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
		user.CreatedBy = ""
		user.UpdatedBy = ""
		_, err := userRepo.Create(ctx, user)
		if err != nil {
			return err
		}

		address := &entity.Address{}
		copier.Copy(address, request.Address)
		address.UserID = user.ID
		address.CreatedBy = ""
		address.UpdatedBy = ""
		_, err = addressRepo.Create(ctx, address)
		if err != nil {
			return err
		}

		ccs := []*entity.CreditCard{}
		for _, ccRequest := range request.CreditCards {
			cc := &entity.CreditCard{}
			cc.UserID = user.ID
			cc.Number = ccRequest.Number
			cc.CreatedBy = ""
			cc.UpdatedBy = ""
			_, err := ccRepo.Create(ctx, cc)
			if err != nil {
				return err
			}
			ccs = append(ccs, cc)
		}

		userRoles := []*entity.UserRole{}
		for _, urRequest := range request.UserRoles {
			userRole := &entity.UserRole{}
			userRole.UserID = user.ID
			userRole.RoleID = urRequest.RoleID
			userRole.CreatedBy = ""
			userRole.UpdatedBy = ""
			_, err := userRoleRepo.Create(ctx, userRole)
			if err != nil {
				return err
			}
			userRoles = append(userRoles, userRole)
		}

		data = user
		data.Address = address
		data.CreditCards = ccs
		data.UserRoles = userRoles

		return nil
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *useCase) UpdateUser(ctx context.Context, request *model.CreateUserRequest) (*entity.User, error) {
	data := &entity.User{}
	return data, nil
}

func (u *useCase) DeleteUser(ctx context.Context) error {
	return nil
}
