package main

import (
	"context"
	"log"

	"github.com/go-pg/pg"
)

var database *DataBase

type DataBase struct {
	DB *pg.DB
}

func initDatabase(options *pg.Options) {
	db := CreateDatabaseConnection(options)
	log.Println("Database init succes!")
	database = &DataBase{
		DB: db,
	}
}

func CreateDatabaseConnection(options *pg.Options) *pg.DB {
	db := pg.Connect(options)

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		panic(err)
	}
	return db
}
