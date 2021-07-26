package entity

import (
	"github.com/todo-host/todo-host-api/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

func GetUserByUsername(userName string) (*User, error) {
	user := &User{Username: userName}
	err := config.POSTGRES.DB.Model(user).Where("username = ?", user.Username).Select()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func PostUser(userName string, password string, email string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	hashedPasswordString := string(hashedPassword)

	if err != nil {
		return nil, err
	}

	user := &User{
		Username: userName,
		Password: hashedPasswordString,
		Email:    email,
	}

	_, err = config.POSTGRES.DB.Model(user).Insert()

	if err != nil {
		return nil, err
	}

	return user, err
}
