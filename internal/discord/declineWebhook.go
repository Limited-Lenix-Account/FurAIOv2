package discord

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/disgoorg/disgo/discord"
	disgo "github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/snowflake/v2"
)

func PaLiquorDeclineWebhook(data PaDiscordData) {
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

	discordEmbed.SetTitle("Checkout Declined! ❌")
	discordEmbed.SetThumbnail(data.ProductImage)
	discordEmbed.SetFooterText("Built in Go • Made by Lenix")
	discordEmbed.SetFooterIcon(avatarURL)
	discordEmbed.SetTimestamp(time.Now().Local())
	//
	discordEmbed.AddField("Name 🏷️", data.ProductName, true)
	discordEmbed.AddField("Mode :hotel:", data.Mode, true)
	discordEmbed.AddField("\u200b", "\u200b", true) // blank field
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
	discordEmbed.AddField("Task Time :stopwatch:", fmt.Sprintf("%2f", data.TaskTime), true)
	discordEmbed.AddField("Decline Reason :pencil2:", data.DeclineReason, true)
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
