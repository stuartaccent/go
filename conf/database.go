package conf

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Gorm postgres dialect interface
)

func ConnectDB() (*gorm.DB, error) {
	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	// Define DB connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, username, dbname, password)

	// connect to db URI
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		log.Println("error", err)
	}

	return db, err
}

func WaitForDB() {
	for {
		if db, err := ConnectDB(); err == nil {
			fmt.Println("Connected to database")
			db.Close()
			break
		}
		time.Sleep(2 * time.Second)
	}
}
