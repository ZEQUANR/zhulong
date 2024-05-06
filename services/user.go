package services

import (
	"context"
	"fmt"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/user"
)

var client = driver.MysqlClient

func QueryUserByAccountPassword(account string, password string) (*ent.User, error) {
	ctx := context.Background()

	u, err := client.User.
		Query().
		Select().
		Where(user.Account(account), user.Password(password)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryUserByAccountPassword %w", err)
	}

	return u, nil
}

func QueryUserById(id int) (*ent.User, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryUserById %w", err)
	}

	return user, nil
}

func QueryStudentsById(id int) (*ent.Students, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryStudentsById %w", err)
	}

	info, err := user.
		QueryStudents().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryStudentsById %w", err)
	}

	return info, nil
}

func QueryTeachersById(id int) (*ent.Teachers, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryTeachersById %w", err)
	}

	info, err := user.
		QueryTeachers().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryTeachersById %w", err)
	}

	return info, nil
}

func QueryAdministratorsById(id int) (*ent.Administrators, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryAdministratorsById %w", err)
	}

	info, err := user.
		QueryAdministrators().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryAdministratorsById %w", err)
	}

	return info, nil
}
