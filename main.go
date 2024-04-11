package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ZEQUANR/zhulong/ent"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	pass     = "123456"
	host     = "127.0.0.1"
	port     = "3306"
	database = "home"
)

func main() {
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, database))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

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
