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

const (
	avatarURL  = "https://pbs.twimg.com/profile_images/1805435885149650944/9D7uUhny_400x400.jpg"
	webhookURL = ""
)

func HotelCheckoutWebhook(data SuccessfulHotelData) {
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

	discordEmbed.SetTitle("Hotel Secured! ✅")
	// discordEmbed.SetURL(productUrl)
	discordEmbed.SetDescription("Click the link above to view inventory page")
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

func PaLiquorCheckoutWebhook(data PaDiscordData) {
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

	discordEmbed.SetTitle("Popping Bottles! 🥂")
	discordEmbed.SetThumbnail(data.ProductImage)
	discordEmbed.SetFooterText("Built in Go • Made by Lenix")
	discordEmbed.SetFooterIcon(avatarURL)
	discordEmbed.SetTimestamp(time.Now().Local())
	//
	discordEmbed.AddField("Name 🍷", data.ProductName, true)
	discordEmbed.AddField("Mode :hotel:", data.Mode, true)
	discordEmbed.AddField("Profile Name 🔖", data.ProfileName, true)
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("Size", data.ProductVol, true)
	discordEmbed.AddField("Total Cost (after tax) :moneybag:", data.Subtotal, true)
	discordEmbed.AddField("Profile Email :e_mail:", data.ProfileEmail, true)
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
	discordEmbed.AddField("Checkout Time :stopwatch:", strconv.FormatFloat(data.CheckoutTime, 'f', -1, 64), true)
	discordEmbed.AddField("Task Time :stopwatch:", strconv.FormatFloat(data.TaskTime, 'f', -1, 64), true)
	discordEmbed.AddField("Confirmation Number :pencil2:", data.OrderNumber, true)
	// discordEmbed.AddField("Order Link :computer:", "[Click Here]("+data.OrderLink.String()+")", true)

	_, err := client.Rest().CreateWebhookMessage(sf, channel, discord.WebhookMessageCreate{
		Username:  "Lenix's Helper",
		AvatarURL: avatarURL,
		Embeds:    []discord.Embed{discordEmbed.Build()},
	}, true, snowflake.ID(0))

	if err != nil {
		log.Printf("Error Sending Discord Embed: %s\n", err)
	}
}
