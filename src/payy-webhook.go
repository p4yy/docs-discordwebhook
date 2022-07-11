package main

import (
	"github.com/gtuk/discordwebhook"

	"github.com/thatisuday/commando"

	"fmt"
)

func main() {
	// configure commando
	commando.
		SetExecutableName("payy").
		SetVersion("1.0.0").
		SetDescription("This tool send a message to discord webhooks")

	// configure the root command
	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("type payy help or -h")
		})

	// configure info command
	commando.
		Register("payy.send_discordwebhook").
		SetShortDescription("send message to discord webhook. See docs at github.com/p4yy/docs-discordwebhook").
		SetDescription("This command will sending message to discord webhook").
		AddArgument("botname", "name of bot", "payy").
		AddArgument("url_link", "url", "discord.com/api").
		AddArgument("msg_1", "message text_1", " ").
		AddArgument("msg_2", "message text_2", " ").
		AddArgument("msg_3", "message text_3", " ").
		AddArgument("msg_4", "message text_4", " ").
		AddArgument("msg_5", "message text_5", " ").
		AddArgument("msg_6", "message text_6", " ").
		AddArgument("msg_7", "message text_7", " ").
		AddArgument("msg_8", "message text_8", " ").
		AddArgument("msg_9", "message text_9", " ").
		AddArgument("msg_10", "message text_10", " ").
		AddArgument("msg_11", "message text_11", " ").
		AddArgument("msg_12", "message text_12", " ").
		AddArgument("msg_13", "message text_13", " ").
		AddArgument("msg_14", "message text_14", " ").
		AddArgument("msg_15", "message text_15", " ").
		AddArgument("msg_16", "message text_16", " ").
		AddArgument("msg_17", "message text_17", " ").
		AddArgument("msg_18", "message text_18", " ").
		AddArgument("msg_19", "message text_19", " ").
		AddArgument("msg_20", "message text_20", " ").
		AddArgument("msg_21", "message text_21", " ").
		AddArgument("msg_22", "message text_22", " ").
		AddArgument("msg_23", "message text_23", " ").
		AddArgument("msg_24", "message text_24", " ").
		AddArgument("msg_25", "message text_25", " ").
		AddArgument("msg_26", "message text_26", " ").
		AddArgument("msg_27", "message text_27", " ").
		AddArgument("msg_28", "message text_28", " ").
		AddArgument("msg_29", "message text_29", " ").
		AddArgument("msg_30", "message text_30", " ").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			username := args["botname"].Value
			url := args["url_link"].Value
			content := ""
			//content := args["msg_1"].Value + "\n" + args["msg_2"].Value + "\n" + args["msg_3"].Value + "\n"
			for x := 1; x <= 30; x++ {
				s := fmt.Sprint(x)
				content = content + args["msg_"+s].Value + "\n"
			}
			message := discordwebhook.Message{
				Username: &username,
				Content:  &content,
			}
			discordwebhook.SendMessage(url, message)
		})
	// parse command-line arguments
	commando.Parse(nil)
}
