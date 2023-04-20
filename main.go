package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gogorm/router"
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

func init() {
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
}

// @title           Hacktiv8 Users API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3030
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	// db.Debug().AutoMigrate(&entity.User{}, &entity.Product{})
	// db.Debug().Migrator().DropTable(&entity.Product{})
	// db.Debug().Migrator().DropTable(&entity.User{})

	router.StartServer(db.Debug()).Run(":3030")

	// CreateData()
	// UpdateData()
	// GetAllData()
	// GetDataByID()
	// DeleteData()
}

// func DeleteData() {
// 	user := entity.User{}
// 	db.Debug().Where("id = ?", 1).Delete(&user)
// 	// println(user)
// }

// func GetDataByID() {
// 	user := entity.User{}

// 	db.Debug().Where("id = ?", 3).First(&user)

// 	fmt.Println(user)
// }

// func GetAllData() {
// 	users := []entity.User{}

// 	db.Debug().Find(&users)

// 	fmt.Println(users)
// }

// func CreateData() {
// 	// INSERT 1 DATA INTO TABLE
// 	// user := entity.User{
// 	// 	Email: "",
// 	// }

// 	// if err := db.Debug().Create(&user).Error; err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println(user)

// 	// INSERT MULTIPLE DATA INTO TABLE
// 	users := []entity.User{
// 		{Email: ""},
// 		{Email: ""},
// 		{Email: ""},
// 	}

// 	if err := db.Debug().CreateInBatches(&users, len(users)).Error; err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(users)
// }

// func UpdateData() {
// 	user := entity.User{}

// 	if err := db.Debug().Model(&user).Where("id = ?", 3).Updates(entity.User{
// 		Email: "gantiemail@gmail.com",
// 	}).Error; err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(user)
// }
