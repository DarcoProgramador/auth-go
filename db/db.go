package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;type:varchar(50);default:null"`
	Password string `gorm:"unique;not null;default:null"`
	Email    string `gorm:"unique;not null;type:varchar(100);default:null"`
}

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	return &DB{initializer("postgres://pg:pass@localhost:5432/auth?sslmode=disable")}
}

func initializer(Url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(Url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&User{})

	return db
}
