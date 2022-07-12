package main

import (
	"strings"

	"github.com/DisgoOrg/disgohook"

	"github.com/DisgoOrg/disgohook/api"

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

	commando.
		Register("payy.send_embeds").
		SetShortDescription("send embeds to discord webhook. See docs at github.com/p4yy/simple-discordwebhook").
		SetDescription("This command will sending embeds to discord webhook").
		AddFlag("url", "WebhookID/WebhookToken", commando.String, "WebhookID/WebhookToken").
		AddFlag("thumbnail", "url thumbnail", commando.String, "https://cdn.discordapp.com/attachments/992818839320547409/996094836824359032/unknown.png").
		AddFlag("message", "message text to send", commando.String, "Hello world!").
		AddFlag("title", "Set Title embed", commando.String, "This is title").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			index_url := strings.Index(fmt.Sprint(flags["url"]), "false}")
			index_thumbnail := strings.Index(fmt.Sprint(flags["thumbnail"]), "false}")
			index_msg := strings.Index(fmt.Sprint(flags["message"]), "false}")
			index_title := strings.Index(fmt.Sprint(flags["title"]), "false}")
			text_url := strings.Replace(string(fmt.Sprint(flags["url"])[index_url+7:len(fmt.Sprint(flags["url"]))-1]), "\\n", "\n", -1)
			text_thumbnail := strings.Replace(string(fmt.Sprint(flags["thumbnail"])[index_thumbnail+7:len(fmt.Sprint(flags["thumbnail"]))-1]), "\\n", "\n", -1)
			text_message := strings.Replace(string(fmt.Sprint(flags["message"])[index_msg+7:len(fmt.Sprint(flags["message"]))-1]), "\\n", "\n", -1)
			text_title := strings.Replace(string(fmt.Sprint(flags["title"])[index_title+7:len(fmt.Sprint(flags["title"]))-1]), "\\n", "\n", -1)
			webhook, err := disgohook.NewWebhookClientByToken(nil, nil, text_url)
			if err != nil {
				fmt.Printf("failed to create webhook: %s", err)
				return
			}

			if _, err := webhook.SendEmbeds(api.NewEmbedBuilder().
				AddField(text_title, text_message, false).
				SetColor(1127128).
				SetThumbnail(text_thumbnail).
				Build(),
			); err != nil {
				fmt.Printf("Didn't found url thumbnail")
				webhook.SendEmbeds(api.NewEmbedBuilder().
					AddField(text_title, text_message, false).
					SetColor(1127128).
					Build(),
				)
			}
		})
	// configure info command
	commando.
		Register("payy.send_message").
		SetShortDescription("send message to discord webhook. See docs at github.com/p4yy/simple-discordwebhook").
		SetDescription("This command will sending message to discord webhook").
		AddFlag("url", "WebhookID/WebhookToken", commando.String, "WebhookID/WebhookToken").
		AddFlag("message", "message text to send", commando.String, "Hello world!").
		AddFlag("botname", "Set name of bot", commando.String, "payy").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			index_url := strings.Index(fmt.Sprint(flags["url"]), "false}")
			index_msg := strings.Index(fmt.Sprint(flags["message"]), "false}")
			index_botname := strings.Index(fmt.Sprint(flags["botname"]), "false}")
			username := strings.Replace(string(fmt.Sprint(flags["botname"])[index_botname+7:len(fmt.Sprint(flags["botname"]))-1]), "\\n", "\n", -1)
			url := "https://discord.com/api/webhooks/" + strings.Replace(string(fmt.Sprint(flags["url"])[index_url+7:len(fmt.Sprint(flags["url"]))-1]), "\\n", "\n", -1)
			content := strings.Replace(string(fmt.Sprint(flags["message"])[index_msg+7:len(fmt.Sprint(flags["message"]))-1]), "\\n", "\n", -1)
			message := discordwebhook.Message{
				Username: &username,
				Content:  &content,
			}
			discordwebhook.SendMessage(url, message)
		})
	// parse command-line arguments
	commando.Parse(nil)
}
