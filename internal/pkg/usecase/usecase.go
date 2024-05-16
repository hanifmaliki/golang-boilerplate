package usecase

import "github.com/hanifmaliki/golang-boilerplate/internal/pkg/repository"

type UseCase interface {
	UserUseCase
}

type useCase struct {
	repository repository.Repository
}

func NewUseCase(r repository.Repository) UseCase {
	return &useCase{
		repository: r,
	}
}
