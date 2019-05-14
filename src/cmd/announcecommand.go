package cmd

import (
	"../framework"
	"../util"
	"strings"
	"time"
	"github.com/bwmarrin/discordgo"
)

// todo define standard embed colors and timeout durations app-wide
func AnnounceCommand(ctx framework.Context) {
	prompt := "Please reply with the title of your announcement. Or reply with 'skip' to proceed to announcement body."
	titlem, err := util.QueryInput(ctx, prompt, 10*time.Second)
	if err != nil {
		return
	}
	title := titlem.Content
	prompt = "What is your announcement?"
	bodym, err := util.QueryInput(ctx, prompt, 10*time.Second)
	if err != nil {
		return
	}
	body := bodym.Content
	prompt = "Provide the channel id for the channel you'd like to send your announcement"
	cid, err := util.QueryInput(ctx, prompt, 10*time.Second)
	if err != nil || cid == nil {
		return
	}
	channel, err := ctx.Discord.Channel(cid.Content)
	if err != nil {
		return
	}
	embed := new(discordgo.MessageEmbed)
	
	if strings.ToLower(title) != "skip" {
		embed.Title = title
	} 
	embed.Description = body
	embed.Color = 42320
	ctx.Discord.ChannelMessageSendEmbed(channel.ID, embed)
}

