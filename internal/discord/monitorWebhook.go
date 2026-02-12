package discord

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	disgo "github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
)

func SendMonitorWebhook(data SuccessfulHotelData) {
	// How can I do this and not get an import cycle error
	snowflakeReg := regexp.MustCompile("[0-9]{18,19}")
	it, _ := strconv.Atoi(snowflakeReg.FindString(webhookURL))
	split := strings.Split(webhookURL, "/")

	var (
		sf      = snowflake.ID(it)
		channel = split[len(split)-1]
	)

	client := disgo.New(sf, channel)
	defer client.Close(context.TODO())

	discordEmbed := discord.NewEmbedBuilder()

	discordEmbed.SetTitle("Rooms Found! 🚨")
	// discordEmbed.SetURL(productUrl)
	discordEmbed.SetThumbnail(data.HotelImage)
	discordEmbed.SetFooterText("Built in Go • Made by Lenix")
	discordEmbed.SetFooterIcon(avatarURL)
	discordEmbed.SetTimestamp(time.Now().Local())
	//
	discordEmbed.AddField("Check-In date :calendar:", data.StartDate, true)
	discordEmbed.AddField("Check-Out date :calendar:", data.EndDate, true)
	discordEmbed.AddField("Hotel Name :hotel:", data.HotelName, true)
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("Room Type :sleeping_accommodation:", data.BlockName, true)
	discordEmbed.AddField("Total Cost (after tax) :moneybag:", data.TotalCost, true)
	discordEmbed.AddField("Profile Email :e_mail:", data.ProfileEmail, true)
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("Checkout Time :stopwatch:", strconv.FormatFloat(data.CheckoutTime, 'f', -1, 64), true)
	discordEmbed.AddField("Confirmation Number :pencil2:", data.AckNumber, true)
	discordEmbed.AddField("Order Link :computer:", "[Click Here]("+data.OrderLink.String()+")", true)

	_, err := client.Rest().CreateWebhookMessage(sf, channel, discord.WebhookMessageCreate{
		Username:  "Lenix's Helper",
		AvatarURL: avatarURL,
		Embeds:    []discord.Embed{discordEmbed.Build()},
	}, true, snowflake.ID(0))

	if err != nil {
		log.Printf("Error Sending Discord Embed: %s\n", err)
	}
}
