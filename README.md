# Simple discord webhook

## Command
- payy.send_message
- payy.send_embeds

## Flags
- --url >> ``WebhookID/WebhookToken``
- --title >> ``Set title on payy.send_embeds``
- --botname >> ``You just can set this on payy.send_message``
- --thumbnail >> ``Set url thumbnail on payy.send_embeds``
- --message >> ``Set content to send``

### Example usage
- go run .\payy-webhook.go payy.send_message --botname ``"BOTNAME"`` --url ``"WebhookID/WebhookToken"`` --message ``"Hello\nWorld!"`` 
- go run .\payy-webhook.go payy.send_embeds --url ``"WebhookID/WebhookToken"`` --message ``"Hello\nWorld!"`` --title ``"Hello Title"`` --thumbnail ``"url thumbnail"``
