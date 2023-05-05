package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/thaiha1607/iot-app-backend/config"
)

type FeedDataType []struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
}

func FetchDataFromFeed(feedName string, feedDataObj *FeedDataType) {
	if matched, _ := regexp.MatchString("^[a-z0-9-]+$", feedName); !matched {
		log.Panicln("Invalid feed name!")
	}
	url := config.GetAdafruitIOFeedURL(feedName)
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(string(body))
	if jsonErr := json.Unmarshal(body, &feedDataObj); jsonErr != nil {
		log.Panicln(jsonErr)
	}
}
