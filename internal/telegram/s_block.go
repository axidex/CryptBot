package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const SBlockHint = "/sblock a=11 b=17 c=15 z0=12 n=7"

func (c *Client) handleSBlock(message *tgbotapi.Message) {
	requiredKeys := []string{"a", "b", "c", "z0", "n"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, SBlockHint, message.MessageID)
		return
	}

	a := params["a"]
	b := params["b"]
	cc := params["c"]
	z0 := params["z0"]
	n := params["n"]

	if a <= 0 || b <= 0 || cc <= 0 || z0 <= 0 || n <= 0 {
		c.sendMessage(message.Chat.ID, SBlockHint, message.MessageID)
		return
	}

	result := problems.SBlock(a, b, cc, z0, n)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
