package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/kalyasik/blog_backend/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectToDB(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	database, err := gorm.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatalf("Cannot connect to %s database: %v", dbDriver, err)
	}
	log.Printf("We are connected to %s database", dbDriver)

	RunMigrations(database)
	CreateTestDataMigration(database)

	DB = database
}

func InitDatabase() {
	ConnectToDB(
		"postgres",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"),
	)
}

func RunMigrations(db *gorm.DB) {
	/* Create tables and ForeignKey */
	db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Comments{})
	db.Model(&models.Posts{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Comments{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
}

func CreateTestDataMigration(db *gorm.DB) {
	/* Create test data for debugging */
	db.Create(&models.Users{Username: "Bret", Password: "password", Name: "Leanne Graham",
		Email: "Sincere@april.biz"})
	db.Create(&models.Users{Username: "Antonette", Password: "password", Name: "Ervin Howell",
		Email: "Shanna@melissa.tv"})
	db.Create(&models.Users{Username: "Samantha", Password: "password", Name: "Clementine Bauch",
		Email: "Nathan@yesenia.net"})

	db.Create(&models.Posts{Title: "sunt aut facere", Body: "quia et suscipit\nsuscipit recusandae", UserID: 1})
	db.Create(&models.Posts{Title: "qui est esse", Body: "quia et suscipit\nsuscipit recusandae", UserID: 2})
	db.Create(&models.Posts{Title: "eum et est occaecati", Body: "quia et suscipit\nsuscipit recusandae", UserID: 3})

	db.Create(&models.Comments{Name: "Leanne Graham", Email: "Sincere@april.biz", Body: "laudantium enim quasi", PostId: 1, UserId: 1})
	db.Create(&models.Comments{Name: "Antonette", Email: "Shanna@melissa.tv", Body: "laudantium enim quasi", PostId: 2, UserId: 2})
	db.Create(&models.Comments{Name: "Samantha", Email: "Nathan@yesenia.net", Body: "laudantium enim quasi", PostId: 3, UserId: 3})
}
