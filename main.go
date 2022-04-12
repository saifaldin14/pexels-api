package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var TOKEN = os.Getenv("API_KEY")

	var c = NewClient(TOKEN)
}
