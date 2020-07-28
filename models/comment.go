package models

import "context"

var ctx context.Context = context.TODO()

func GetComments() ([]string, error) {
	return client.LRange(ctx, "comments", 0, 10).Result()
}

func PostComment(comment string) error {
	return client.LPush(ctx, "comments", comment).Err()
}
