package main

import (
	"fmt"
	"gogin/internal/routers"
	"gogin/pkg"
	"log"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	database, err := pkg.Pgdb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(database)

	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
