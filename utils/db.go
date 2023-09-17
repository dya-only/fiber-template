package utils

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go-template/ent"
	"log"
	"os"
)

var DbConn *ent.Client

func CreateDbConnection() {
	dsn := "root:root@tcp(localhost:3306)/entgo?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Failed to connect database. \n", err)
		os.Exit(1)
	}

	log.Println("DB Connected.")
	DbConn = client

	ctx := context.Background()
	if err := DbConn.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
