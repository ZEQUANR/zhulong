package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func main() {
	CreateUser(context.Background(), driver.MysqlClient)
}
