package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const StartHint = "Welcome! Use /help to see available commands."

func (c *Client) handleStart(message *tgbotapi.Message) {
	c.sendMessage(message.Chat.ID, StartHint, message.MessageID)
}

func (c *Client) handleHelp(message *tgbotapi.Message) {
	var hints []string
	for command, _ := range c.commands {
		if command == CommandHelp || command == CommandStart {
			continue
		}
		hints = append(hints, command.GetHint())

	}
	c.sendMessage(message.Chat.ID, strings.Join(hints, "\n"), message.MessageID)
}
