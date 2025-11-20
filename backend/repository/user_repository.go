package repository

import (
	"context"
	"errors"
	"fmt"
	"notes_project/config"
	"notes_project/models"
)

func CreateUser(user models.User) error {
	_, err := config.DB.Exec(context.Background(),
        "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
        user.Username, user.Email, user.Password,
    )

	if err != nil {
		fmt.Println("DB ERROR:", err)
	}

	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := config.DB.QueryRow(context.Background(),
		"SELECT id, email, password FROM users WHERE email=$1",
		email,
	)

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}
