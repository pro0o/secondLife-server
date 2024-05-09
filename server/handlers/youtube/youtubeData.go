package youtube

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"secondLife/parser"

	"github.com/joho/godotenv"
)

func GetYoutubeData(query string) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	apiKey := os.Getenv("API_KEY")

	maxResults := 3

	// Prepare the query to include recycling-related keywords
	q := url.QueryEscape(query + " recycling")

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&part=snippet&type=video&q=%s&maxResults=%d", apiKey, q, maxResults)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result, err := parser.JsonParser(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
