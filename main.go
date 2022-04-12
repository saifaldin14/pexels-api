package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var TOKEN = os.Getenv("API_KEY")

	var c = NewClient(TOKEN)
}
