package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/model"
)

func createAdministrators(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAccount("admin").
		SetPassword("e10adc3949ba59abbe56e057f20f883e").
		SetRole(model.Admin).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Administrators.
		Create().
		SetAvatar("").
		SetName("李宏伟").
		SetCollege("计算机系").
		SetPhone("17315485452").
		SetNumber("065000").
		SetUsers(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)

	return nil
}

func createTeachers(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAccount("teacher").
		SetPassword("e10adc3949ba59abbe56e057f20f883e").
		SetRole(model.Teacher).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Teachers.
		Create().
		SetAvatar("").
		SetName("李向明").
		SetCollege("精工系").
		SetPhone("151458459566").
		SetNumber("041500").
		SetUsers(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)

	return nil
}

func createStudents(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAccount("student").
		SetPassword("e10adc3949ba59abbe56e057f20f883e").
		SetRole(model.Student).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Students.
		Create().
		SetAvatar("").
		SetName("罗冲围").
		SetCollege("通信系").
		SetPhone("19854551254").
		SetMajor("母猪产后护理").
		SetClass("本科八班").
		SetNumber("041500").
		SetUsers(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)

	return nil
}

func main() {
	// ctx := context.Background()
	// client := driver.MysqlClient

	// createAdministrators(ctx, client)
	// createTeachers(ctx, client)
	// createStudents(ctx, client)

	fmt.Println(time.Now())
}
