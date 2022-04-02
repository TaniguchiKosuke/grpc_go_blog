package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"grpc_go_blog/domain/entity"
)

type DB struct {
	Conn *gorm.DB
}

func NewDB() *DB {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if err := conn.AutoMigrate(&entity.User{}); err != nil {
		log.Println(err.Error())
		return nil
	}

	db := new(DB)
	db.Conn = conn

	return db
}
