package driver

import (
	"context"
	"fmt"
	"log"

	"github.com/ZEQUANR/zhulong/ent"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlClient *ent.Client
	user        = "root"
	pass        = "123456"
	host        = "127.0.0.1"
	port        = "3306"
	database    = "home"
)

func initMysql() {
	var err error
	MysqlClient, err = ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", user, pass, host, port, database))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := MysqlClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
