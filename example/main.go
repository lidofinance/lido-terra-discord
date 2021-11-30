// example module just demonstrates how LidoDiscordBot works
// to join the test channel follow the link https://discord.gg/zcpWa9mZry
package main

import (
	"fmt"
	"github.com/lidofinance/lido-terra-discord/discord"
	"log"
)

func main() {
	token := "OTEwNTEyNzI2MTM5MzU5Mjgy.YZT7Dg.yRKAAl6SBevw_N7j5vwyaBoglzI"
	channelID := "910513741567754295"
	guildID := "910513741567754290"
	bot, err := discord.NewDefaultLidoBot(token, guildID, channelID)
	if err != nil {
		log.Fatalln("failed to create bot:", err)
	}

	participants, err := bot.GetParticipants()
	if err != nil {
		log.Fatalln("failed to get participants list:", err)
	}
	fmt.Println(participants)
	var pic discord.Picture
	message := "test"
	URL := "https://blog.lido.fi/content/images/2021/10/1920x1080.png"

	pictype := "default"
	switch pictype {
	//you can send picture either from local file
	case "local":
		{
			FileName := "/tmp/item.png"
			pic, err = discord.NewFSPicture(FileName)
			if err != nil {
				log.Fatalln("failed to create picture instance:", err)
			}
		}
	// or http url
	case "http":
		{
			pic, err = discord.NewURLPicture("1.jpg", URL)
			if err != nil {
				log.Fatalf("failed to get image with url %s: %v\n", URL, err)
			}
		}
	// or embed url into the message text
	default:
		message = URL + " " + message
	}

	err = bot.SendMessage(participants, message, pic)
	if err != nil {
		log.Fatalln("failed to send message:", err)
	}
}
