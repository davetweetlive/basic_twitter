package models

import (
	"errors"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("User Not Found!")
	ErrInvalidLogin = errors.New("Invalid Login!")
)

func RegisterUser(username, password string) error {
	cost := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}

	return client.Set(ctx, "user:"+username, hash, 0).Err()
}

func AuthenticateUser(username, password string) error {
	hash, err := client.Get(ctx, "user:"+username).Bytes()
	if err == redis.Nil {

		return ErrUserNotFound
	} else if err != nil {

		return err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {

		return ErrInvalidLogin
	}
	return nil
}
