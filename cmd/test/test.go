package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ZEQUANR/zhulong/driver"
	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/user"
)

func createAdministrators(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAccount("admin").
		SetPassword("e10adc3949ba59abbe56e057f20f883e").
		SetRole(0).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Administrators.
		Create().
		SetName("李宏伟").
		SetCollege("计算机系").
		SetIdentity("06500").
		SetPhone("17315485452").
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
		SetRole(1).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Teachers.
		Create().
		SetName("李向明").
		SetCollege("精工系").
		SetPhone("151458459566").
		SetIdentity("041500").
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
		SetRole(2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user:", a8m)

	card1, err := client.Students.
		Create().
		SetName("罗冲围").
		SetCollege("通信系").
		SetPhone("19854551254").
		SetSubject("母猪产后护理").
		SetClass("本科八班").
		SetIdentity("845100").
		SetUsers(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card:", card1)

	return nil
}

func createThesiss(ctx context.Context, client *ent.Client) error {

	a8m, err := client.User.
		Query().
		Where(user.ID(1)).
		Only(ctx)
	if err != nil {
		fmt.Println("%w", err)
		return fmt.Errorf("%w", err)
	}

	pedro, err := client.Thesis.
		Create().
		SetName("pedro").
		SetURL("/uirl/sdi/saortt/124ve.mp4").
		SetType(0).
		SetStatus(0).
		SetUploaders(a8m).
		Save(ctx)
	if err != nil {
		fmt.Println("%w", err)
		return fmt.Errorf("%w", err)
	}

	lola, err := client.Thesis.
		Create().
		SetName("lola").
		SetURL("/uirl/sdi/safrr/adw5.mp4").
		SetType(0).
		SetStatus(0).
		SetUploaders(a8m).
		Save(ctx)
	if err != nil {
		fmt.Println("%w", err)
		return fmt.Errorf("%w", err)
	}

	log.Println(pedro, lola)
	return nil
}

func main() {
	ctx := context.Background()
	client := driver.MysqlClient

	createAdministrators(ctx, client)
	createTeachers(ctx, client)
	createStudents(ctx, client)

	createThesiss(ctx, client)
}
