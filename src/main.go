package main

import (
	"./cmd"
	"./framework"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	commandPrefix string
	botID         string
	cmdHandler    *framework.CommandHandler
)

func main() {
	discord, err := discordgo.New("Bot <insert token here>")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	cmdHandler = framework.NewCommandHandler()
	registerCommands()

	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "a Pokemon's not cooperating machine")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("GoKyogreBot has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}
	if message.Content[:1] != commandPrefix {
		return
	}

	content := message.Content
	args := strings.Fields(content)
	name := strings.ToLower(args[0][1:])
	fmt.Println(name)
	command, found := cmdHandler.Get(name)
	if !found {
		fmt.Println("Command not found.")
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := framework.NewContext(discord, guild, channel, user, message, cmdHandler)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)

	//content := message.Content

	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}

func registerCommands() {
	cmdHandler.Register("bob", cmd.BobCommand)
}
