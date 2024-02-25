package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load file")
	}

	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")

	fmt.Printf("This is the slack bot token %s \n", SLACK_BOT_TOKEN)
	fmt.Printf("This is the channel id %s \n", CHANNEL_ID)

	api := slack.New(SLACK_BOT_TOKEN)
	channelarr := []string{CHANNEL_ID}
	filearr := []string{"newfile.txt"}

	for i := 0; i < len(filearr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelarr,
			File:     filearr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name %s, URL : %s \n", file.Name, file.URLPrivate)
	}

}
