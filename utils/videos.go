package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

func GetVideos() (videos []Video) {
	fileBytes, err := ioutil.ReadFile("./utils/videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func SaveVideos(videos []Video) {

	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./utils/videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}

}
