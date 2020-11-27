package main // local ip : 192.168.1.44

import (
	"log"
	//	"net/http"
	"fmt"

	"github.com/joho/godotenv"
	R "github.com/mail_encryption_en/routes"
	//	"golang.org/x/oauth2"
	//	"golang.org/x/oauth2/google"
	//R "github.com/jintecheng/yu/v1/routes"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Server On")
	R.Server()
}
