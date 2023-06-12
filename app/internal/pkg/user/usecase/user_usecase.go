package usecase

import (
	"app/internal/pkg/models"
)

type userUsecase struct {
	userRepo models.UserRepositoryI
}

func NewUserUsecase(userRepo models.UserRepositoryI) models.UserUsecaseI {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) GetUserById(id int) (*models.User, error) {
	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) GetUserByLogin(login string) (*models.User, error) {
	user, err := uc.userRepo.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) Authenticate(login string, password string) (bool, error) {
	user, err := uc.userRepo.GetUserByLogin(login)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func (uc *userUsecase) Create(user *models.User) (int, error) {
	id, err := uc.userRepo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *userUsecase) Update(id int, newUser *models.User) error {
	newUser.ID = id
	err := uc.userRepo.Update(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUsecase) Delete(id int) error {
	err := uc.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
