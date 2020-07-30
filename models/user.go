package models

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound  = errors.New("User Not Found!")
	ErrInvalidLogin  = errors.New("Invalid Login!")
	ErrUsernameTaken = errors.New("The username is already taken!")
)

type User struct {
	id int64
}

func NewUser(username string, hash []byte) (*User, error) {
	exists, err := client.HExists(ctx, "user:by-username", username).Result()
	if exists {
		return nil, ErrUsernameTaken
	}
	id, err := client.Incr(ctx, "user:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(ctx, key, "id", id)
	pipe.HSet(ctx, key, "username", username)
	pipe.HSet(ctx, key, "hash", hash)
	pipe.HSet(ctx, "user:by-username", username, id)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &User{id}, nil
}

func (user *User) GetId() (int64, error) {

	return user.id, nil
}

func (user *User) GetUsername() (string, error) {
	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(ctx, key, "username").Result()
}

func (user *User) GetHash() ([]byte, error) {
	key := fmt.Sprintf("user:%d", user.id)
	return client.HGet(ctx, key, "hash").Bytes()
}

func (user *User) Authenticate(password string) error {
	hash, err := user.GetHash()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrInvalidLogin
	}
	return err
}

func GetUserById(id int64) (*User, error) {
	return &User{id}, nil
}

func GetUserByUsername(username string) (*User, error) {
	id, err := client.HGet(ctx, "user:by-username", username).Int64()
	if err == redis.Nil {

		return nil, ErrUserNotFound
	} else if err != nil {

		return nil, err
	}
	return GetUserById(id)

}

func RegisterUser(username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	_, err = NewUser(username, hash)
	return err
}

func AuthenticateUser(username, password string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, user.Authenticate(password)
}
