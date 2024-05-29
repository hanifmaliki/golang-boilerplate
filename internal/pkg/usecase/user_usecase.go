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
	GetUser(ctx context.Context, id uint, query *pkg_model.Query) (*entity.User, error)
	GetUsers(ctx context.Context, request *model.GetUserRequest, query *pkg_model.Query) ([]*entity.User, *pkg_model.Pagination, error)
	CreateUser(ctx context.Context, request *model.CreateUserRequest, by string) (*entity.User, error)
	UpdateUser(ctx context.Context, request *model.CreateUserRequest, by string) (*entity.User, error)
	DeleteUser(ctx context.Context, id uint, by string) error
}

func (u *useCase) GetUser(ctx context.Context, id uint, query *pkg_model.Query) (*entity.User, error) {
	conds := &entity.User{}
	conds.ID = id
	return u.repository.UserRepo().FindOne(ctx, conds, query)
}

func (u *useCase) GetUsers(ctx context.Context, request *model.GetUserRequest, query *pkg_model.Query) ([]*entity.User, *pkg_model.Pagination, error) {
	return u.repository.UserRepo().Find(ctx, request, query)
}

func (u *useCase) CreateUser(ctx context.Context, request *model.CreateUserRequest, by string) (*entity.User, error) {
	data := &entity.User{}

	err := u.repository.InTransaction(ctx, func(r repository.Repository) error {
		userRepo := r.UserRepo()
		addressRepo := r.AddressRepo()
		ccRepo := r.CreditCardRepo()
		userRoleRepo := r.UserRoleRepo()

		user := &entity.User{}
		copier.Copy(user, request)
		user.CreatedBy = by
		user.UpdatedBy = by
		err := userRepo.Create(ctx, user)
		if err != nil {
			return err
		}

		address := &entity.Address{}
		copier.Copy(address, request.Address)
		address.UserID = user.ID
		address.CreatedBy = by
		address.UpdatedBy = by
		err = addressRepo.Create(ctx, address)
		if err != nil {
			return err
		}

		ccs := []*entity.CreditCard{}
		for _, ccRequest := range request.CreditCards {
			cc := &entity.CreditCard{}
			cc.UserID = user.ID
			cc.Number = ccRequest.Number
			cc.CreatedBy = by
			cc.UpdatedBy = by
			err := ccRepo.Create(ctx, cc)
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
			userRole.CreatedBy = by
			userRole.UpdatedBy = by
			err := userRoleRepo.Create(ctx, userRole)
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

func (u *useCase) UpdateUser(ctx context.Context, request *model.CreateUserRequest, by string) (*entity.User, error) {
	data := &entity.User{}

	err := u.repository.InTransaction(ctx, func(r repository.Repository) error {
		userRepo := r.UserRepo()
		addressRepo := r.AddressRepo()
		ccRepo := r.CreditCardRepo()
		userRoleRepo := r.UserRoleRepo()

		user := &entity.User{}
		copier.Copy(user, request)
		user.UpdatedBy = by
		err := userRepo.Save(ctx, user)
		if err != nil {
			return err
		}

		address := &entity.Address{}
		copier.Copy(address, request.Address)
		address.UserID = user.ID
		address.UpdatedBy = by
		err = addressRepo.Save(ctx, address)
		if err != nil {
			return err
		}

		ccs := []*entity.CreditCard{}
		for _, ccRequest := range request.CreditCards {
			cc := &entity.CreditCard{}
			cc.UserID = user.ID
			cc.Number = ccRequest.Number
			cc.UpdatedBy = by
			err := ccRepo.Save(ctx, cc)
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
			userRole.UpdatedBy = by
			err := userRoleRepo.Save(ctx, userRole)
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

func (u *useCase) DeleteUser(ctx context.Context, id uint, by string) error {
	err := u.repository.InTransaction(ctx, func(r repository.Repository) error {
		userRepo := r.UserRepo()
		addressRepo := r.AddressRepo()
		ccRepo := r.CreditCardRepo()
		userRoleRepo := r.UserRoleRepo()

		err := userRepo.Update(ctx, &entity.User{Base: pkg_model.Base{DeletedBy: by}}, &entity.User{Base: pkg_model.Base{ID: id}})
		if err != nil {
			return err
		}
		err = userRepo.Delete(ctx, &entity.User{Base: pkg_model.Base{ID: id}})
		if err != nil {
			return err
		}

		err = addressRepo.Update(ctx, &entity.Address{Base: pkg_model.Base{DeletedBy: by}}, &entity.Address{UserID: id})
		if err != nil {
			return err
		}
		err = addressRepo.Delete(ctx, &entity.Address{UserID: id})
		if err != nil {
			return err
		}

		err = ccRepo.Update(ctx, &entity.CreditCard{Base: pkg_model.Base{DeletedBy: by}}, &entity.CreditCard{UserID: id})
		if err != nil {
			return err
		}
		err = ccRepo.Delete(ctx, &entity.CreditCard{UserID: id})
		if err != nil {
			return err
		}

		err = userRoleRepo.Update(ctx, &entity.UserRole{Base: pkg_model.Base{DeletedBy: by}}, &entity.UserRole{UserID: id})
		if err != nil {
			return err
		}
		err = userRoleRepo.Delete(ctx, &entity.UserRole{UserID: id})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
