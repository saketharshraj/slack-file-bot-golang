package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error in loading env variables : %v", err)
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"./beach.jpg"}
	
	for i:=0; i<len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("Error in uploading : %v\n", err)
		}

		fmt.Printf("Name : %s\n URL : %s\n", file.Name, file.URL)
	}
}