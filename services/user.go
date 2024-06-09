package services

import (
	"context"
	"fmt"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/user"
	"github.com/ZEQUANR/zhulong/model"
	"github.com/ZEQUANR/zhulong/model/api"
	"github.com/ZEQUANR/zhulong/utils"
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

func QueryTeacherList() ([]api.TeacherList, error) {
	ctx := context.Background()

	users, err := client.User.
		Query().
		Where(user.Role(model.Teacher)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryTeacherList %w", err)
	}

	var arr []api.TeacherList
	for _, elem := range users {
		info, err := elem.
			QueryTeachers().
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("func: QueryTeacherList %w", err)
		}

		arr = append(arr, api.TeacherList{
			Id:      elem.ID,
			Name:    info.Name,
			Phone:   info.Phone,
			Number:  info.Number,
			College: info.College,
		})
	}

	return arr, nil
}

func CreateAdministrators(data api.RegisterAdmin) error {
	ctx := context.Background()

	u, err := client.User.
		Create().
		SetAccount(data.Account).
		SetPassword(utils.Md5Encode(data.Password)).
		SetRole(model.Admin).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: createAdministrators: %w", err)
	}

	_, err = client.Administrators.
		Create().
		SetAvatar("").
		SetName(data.Name).
		SetCollege(data.College).
		SetPhone(data.Phone).
		SetNumber(data.Number).
		SetUsers(u).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: createAdministrators: %w", err)
	}

	return nil
}

func CreateTeachers(data api.RegisterTeacher) error {
	ctx := context.Background()

	u, err := client.User.
		Create().
		SetAccount(data.Account).
		SetPassword(utils.Md5Encode(data.Password)).
		SetRole(model.Teacher).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: CreateTeachers: %w", err)
	}

	_, err = client.Teachers.
		Create().
		SetAvatar("").
		SetName(data.Name).
		SetCollege(data.College).
		SetPhone(data.Phone).
		SetNumber(data.Number).
		SetUsers(u).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: CreateTeachers: %w", err)
	}

	return nil
}

func CreateStudents(data api.RegisterStudent) error {
	ctx := context.Background()

	u, err := client.User.
		Create().
		SetAccount(data.Account).
		SetPassword(utils.Md5Encode(data.Password)).
		SetRole(model.Student).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: CreateStudents: %w", err)
	}

	_, err = client.Students.
		Create().
		SetAvatar("").
		SetName(data.Name).
		SetCollege(data.College).
		SetPhone(data.Phone).
		SetMajor(data.Major).
		SetClass(data.Class).
		SetNumber(data.Number).
		SetUsers(u).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("func: CreateStudents: %w", err)
	}

	return nil
}

// func UpdateAdministratorsById(id int, admin api.Administrator) (*ent.Administrators, error) {
// 	ctx := context.Background()

// 	user, err := client.User.
// 		UpdateOneID(id).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
// 	}

// 	admins, err := user.
// 		QueryAdministrators().
// 		Only(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
// 	}

// 	result, err := admins.
// 		Update().
// 		SetName(admin.Name).
// 		SetCollege(admin.College).
// 		SetPhone(admin.Phone).
// 		SetIdentity(admin.Identity).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateAdministratorsById %w", err)
// 	}

// 	return result, nil
// }

// func UpdateTeachersById(id int, teacher api.Teacher) (*ent.Teachers, error) {
// 	ctx := context.Background()

// 	user, err := client.User.
// 		UpdateOneID(id).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
// 	}

// 	teachers, err := user.
// 		QueryTeachers().
// 		Only(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
// 	}

// 	result, err := teachers.
// 		Update().
// 		SetName(teacher.Name).
// 		SetCollege(teacher.College).
// 		SetPhone(teacher.Phone).
// 		SetIdentity(teacher.Identity).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateTeachersById %w", err)
// 	}

// 	return result, nil
// }

// func UpdateStudentsById(u *ent.User) (*ent.Students, error) {
// 	ctx := context.Background()

// 	result, err := u.Update().SetAccount("").Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("func: UpdateStudentsById %w", err)
// 	}

// 	return result, nil
// }
