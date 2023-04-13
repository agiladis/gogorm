package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gogorm/entity"
)

const (
	DB_USERNAME = "postgres"
	DB_PORT     = "5432"
	DB_HOST     = "localhost"
	DB_PASSWORD = "postgres"
	DB_DATABASE = "gorm"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		// Log to see changes
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	// db.Debug().AutoMigrate(&entity.User{}, &entity.Product{})
	// db.Debug().Migrator().DropTable(&entity.Product{})
	// db.Debug().Migrator().DropTable(&entity.User{})

	// CreateData()
	UpdateData()
}

func CreateData() {
	// INSERT 1 DATA INTO TABLE
	// user := entity.User{
	// 	Email: "agil@gmail.com",
	// }

	// if err := db.Debug().Create(&user).Error; err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

	// INSERT MULTIPLE DATA INTO TABLE
	users := []entity.User{
		{Email: "1@gmail.com"},
		{Email: "2@gmail.com"},
		{Email: "3@gmail.com"},
		{Email: "4@gmail.com"},
		{Email: "5@gmail.com"},
	}

	if err := db.Debug().CreateInBatches(&users, len(users)).Error; err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func UpdateData() {
	user := entity.User{}

	if err := db.Debug().Model(&user).Where("id = ?", 3).Updates(entity.User{
		Email: "gantiemail@gmail.com",
	}).Error; err != nil {
		panic(err)
	}

	fmt.Println(user)
}
