package services

import (
	"context"
	"fmt"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/user"
)

var client = driver.MysqlClient

func QueryUserById(account string, password string) (*ent.User, error) {
	ctx := context.Background()
	u, err := client.User.
		Query().
		Select("id").
		Where(user.Account(account), user.Password(password)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed querying user: %w", err)
	}

	return u, nil
}

func QueryUserInformation(id int) (*ent.User, error) {
	ctx := context.Background()
	u, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed querying user: %w", err)
	}

	return u, nil
}
