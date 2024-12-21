package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Des3Hint = "/3des l=5 r=1 k1=3 k2=7 k3=5"

func (c *Client) handle3Des(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"l", "r", "k1", "k2", "k3"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, Des3Hint, message.MessageID)
		return
	}

	l := params["l"]
	r := params["r"]
	k1 := params["k1"]
	k2 := params["k2"]
	k3 := params["k3"]

	if l < 0 || r < 0 || k1 < 0 || k2 < 0 || k3 < 0 {
		c.sendMessage(message.Chat.ID, Des3Hint, message.MessageID)
		return
	}

	result := problems.DES3(l, r, k1, k2, k3)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
