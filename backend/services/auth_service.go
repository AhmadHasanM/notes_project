package services

import (
	"errors"
	"notes_project/models"
	"notes_project/repository"
	"notes_project/utils"
)

func RegisterService(input models.User) error {
	_, err := repository.GetUserByEmail(input.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	input.Password = hashed
	return repository.CreateUser(input)
}

func LoginService(email string, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("wrong email or password")
	}

	// ubah nama fungsi sesuai utils/password.go
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("wrong email or password")
	}

	// generate token JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
