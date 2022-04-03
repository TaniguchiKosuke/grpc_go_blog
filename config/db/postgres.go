package db

import (
	"fmt"
	"grpc_go_blog/domain/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresSQL() *DB{
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		host = "postgres"
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		user = "kosuke"
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		password = "grpc_go_blog"
	}

	dbname, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		dbname = "grpc_go_blog_db"
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		port = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port) + " sslmode=disable TimeZone=Asia/Tokyo"
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if err = conn.AutoMigrate(&entity.User{}); err != nil {
		log.Println(err)
		return nil
	}

	db := new(DB)
	db.Conn = conn

	return db
}