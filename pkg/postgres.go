package pkg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func Pgdb() (*sqlx.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	config := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable", host, user, dbPort, password, dbName)

	fmt.Println(config)
	return sqlx.Connect("postgres", config)
}
