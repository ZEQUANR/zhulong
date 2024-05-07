package services

import (
	"context"
	"fmt"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/user"
	"github.com/ZEQUANR/zhulong/model/api"
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

func UpdateAdministratorsById(id int, admin api.Administrator) (*ent.Administrators, error) {
	ctx := context.Background()

	user, err := client.User.
		UpdateOneID(id).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
	}

	admins, err := user.
		QueryAdministrators().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
	}

	result, err := admins.
		Update().
		SetName(admin.Name).
		SetCollege(admin.College).
		SetPhone(admin.Phone).
		SetIdentity(admin.Identity).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
	}

	return result, nil
}

func UpdateTeachersById(id int, teacher api.Teacher) (*ent.Teachers, error) {
	ctx := context.Background()

	user, err := client.User.
		UpdateOneID(id).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
	}

	teachers, err := user.
		QueryTeachers().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
	}

	result, err := teachers.
		Update().
		SetName(teacher.Name).
		SetCollege(teacher.College).
		SetPhone(teacher.Phone).
		SetIdentity(teacher.Identity).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
	}

	return result, nil
}

func UpdateStudentsById(id int, student api.Student) (*ent.Students, error) {
	ctx := context.Background()

	user, err := client.User.
		UpdateOneID(id).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateStudentsById %w", err)
	}

	students, err := user.
		QueryStudents().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateStudentsById %w", err)
	}

	result, err := students.
		Update().
		SetName(student.Name).
		SetCollege(student.College).
		SetPhone(student.Phone).
		SetSubject(student.Subject).
		SetClass(student.Class).
		SetIdentity(student.Identity).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UpdateStudentsById %w", err)
	}

	return result, nil
}
