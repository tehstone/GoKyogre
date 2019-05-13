package util

import (
	"../framework"
	"errors"
	"time"
	"github.com/bwmarrin/discordgo"
)

func QueryInput(ctx framework.Context, prompt string, timeout time.Duration) (*discordgo.Message, error) {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, prompt)
	if err != nil {
		return nil, err
	}
	defer func() {
		ctx.Discord.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}()

	timeoutChan := make(chan int)
	go func() {
		time.Sleep(timeout)
		timeoutChan <- 0
	}()

	for {
		select {
		case usermsg := <-NextMessageCreateC(ctx.Discord):
			if usermsg.Author.ID != ctx.User.ID {
				continue
			}
			ctx.Discord.ChannelMessageDelete(usermsg.ChannelID, usermsg.ID)
			return usermsg.Message, nil
		case <-timeoutChan:
			return nil, errors.New("Timed out")
		}
	}
}

// NextMessageCreateC returns a channel for the next MessageCreate event
func NextMessageCreateC(s *discordgo.Session) chan *discordgo.MessageCreate {
	out := make(chan *discordgo.MessageCreate)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageCreate) {
		out <- e
	})
	return out
}

// NextMessageReactionAddC returns a channel for the next MessageReactionAdd event
func NextMessageReactionAddC(s *discordgo.Session) chan *discordgo.MessageReactionAdd {
	out := make(chan *discordgo.MessageReactionAdd)
	s.AddHandlerOnce(func(_ *discordgo.Session, e *discordgo.MessageReactionAdd) {
		out <- e
	})
	return out
}
