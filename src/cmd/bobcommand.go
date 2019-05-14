package cmd

import (
	"../framework"
	"fmt"
	"strings"
	"unicode"
)

func BobifyCommand(ctx framework.Context) {
	if len(ctx.Args) == 0 {
		fmt.Println("No args.")
		return
	}
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	// remove command string
	messagesplit := strings.Fields(ctx.Message.Content)
	newmessage := strings.Join(messagesplit[1:], " ")

	ctx.Reply(toSpongeBoBCase(newmessage))
}

func toSpongeBoBCase(message string) string {
	runes := make([]rune, 0, len(message))
	var upper bool
	for _, c := range message {
		if unicode.IsLetter(c) {
			upper = !upper
			if upper {
				c = unicode.ToUpper(c)
			}
		}
		runes = append(runes, c)
	}
	return string(runes)
}
