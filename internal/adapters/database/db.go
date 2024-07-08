package database

import (
	"fmt"
	"log"

	"github.com/pitchter/orderRealtime/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "postgres"     // as defined in docker-compose.yml
	password = "postgres" // as defined in docker-compose.yml
	dbname   = "order_menu" // as defined in docker-compose.yml
)

func Init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&entities.MenuItem{}, &entities.Order{})

	DB = db
}
