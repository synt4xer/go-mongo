package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/synt4xer/go-mongo/internal/domain"
	"github.com/synt4xer/go-mongo/internal/repository"
)

type UserUseCase struct {
	userRepository *repository.UserRepository
	validator      *validator.Validate
}

func ProvideUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: repo, validator: validator.New()}
}

func (uc *UserUseCase) Save(ctx context.Context, user *domain.User) error {
	if err := uc.validator.Struct(user); err != nil {
		return err
	}

	return uc.userRepository.Save(ctx, user)
}

func (uc *UserUseCase) Update(ctx context.Context, id string, user *domain.User) error {
	_, err := uc.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	err = uc.validator.Struct(user)

	if err != nil {
		return err
	}

	err = uc.userRepository.Update(ctx, id, user)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, id string) error {
	err := uc.userRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) GetAll(ctx context.Context, search string) ([]domain.User, error) {
	users, err := uc.userRepository.GetAll(ctx, search)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := uc.userRepository.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
