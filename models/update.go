package models

import (
	"context"
	"fmt"
	"strconv"
)

var ctx context.Context = context.TODO()

type Update struct {
	id int64
}

func NewUpdate(userId int64, body string) (*Update, error) {
	id, err := client.Incr(ctx, "update:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("update:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(ctx, key, "id", id)
	pipe.HSet(ctx, key, "user_id", userId)
	pipe.HSet(ctx, key, "body", body)
	pipe.LPush(ctx, "updates", id)
	pipe.LPush(ctx, fmt.Sprintf("user:%d:updates", userId), id)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &Update{id}, nil
}
func (update *Update) GetBody() (string, error) {
	key := fmt.Sprintf("update:%d", update.id)
	return client.HGet(ctx, key, "body").Result()
}

func (update *Update) GetUser() (*User, error) {
	key := fmt.Sprintf("update:%d", update.id)
	userId, err := client.HGet(ctx, key, "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserById(userId)
}

func queryUpdates(key string) ([]*Update, error) {
	updateIds, err := client.LRange(ctx, key, 0, 10).Result()
	if err != nil {
		return nil, err
	}
	updates := make([]*Update, len(updateIds))
	for i, strId := range updateIds {
		id, err := strconv.Atoi(strId)
		if err != nil {
			return nil, err
		}
		updates[i] = &Update{int64(id)}
	}
	return updates, nil
}

func GetAllUpdates() ([]*Update, error) {
	return queryUpdates("updates")
}

func GetUpdates(userId int64) ([]*Update, error) {
	key := fmt.Sprintf("user:%d:updates", userId)
	return queryUpdates(key)

}

func PostUpdate(userId int64, body string) error {
	_, err := NewUpdate(userId, body)
	return err
}
