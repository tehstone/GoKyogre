package main

import (
 "fmt"
 "github.com/bwmarrin/discordgo"
)

var (
    commandPrefix string
    botID         string
)

func main() {
  discord, err := discordgo.New("Bot <insert token here>")
  errCheck("error creating discord session", err)
  user, err := discord.User("@me")
  errCheck("error retrieving account", err)
  
  botID = user.ID
  discord.AddHandler(commandHandler)
  discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
        err = discord.UpdateStatus(0, "A friendly helpful bot!")
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

    //content := message.Content

    fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}