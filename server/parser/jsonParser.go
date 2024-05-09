package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"secondLife/types"
	"strings"
)

func ReplaceSpacesWithPlus(s string) string {
	return strings.ReplaceAll(s, " ", "+")
}

func JsonParser(body []byte) ([]byte, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if errorMessage, ok := data["error"].(map[string]interface{}); ok {
		return nil, fmt.Errorf("YouTube API error: %v", errorMessage)
	}

	items, ok := data["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("no items found in YouTube API response")
	}

	var videos []types.Video
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid item format in YouTube API response")
		}

		videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
		title := itemMap["snippet"].(map[string]interface{})["title"].(string)
		channelTitle := itemMap["snippet"].(map[string]interface{})["channelTitle"].(string)
		imageURL := itemMap["snippet"].(map[string]interface{})["thumbnails"].(map[string]interface{})["default"].(map[string]interface{})["url"].(string)

		video := types.Video{
			Image:  imageURL,
			Title:  title,
			URL:    fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID),
			Author: channelTitle,
		}
		videos = append(videos, video)
	}

	result, err := json.Marshal(videos)
	if err != nil {
		return nil, err
	}

	log.Printf("Raw JSON data: %s", result)

	return result, nil
}
