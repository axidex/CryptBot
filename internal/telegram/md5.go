package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Md5Hint = "/md5 a=5 b=1 c=9 d=1 m=6 k=10"

func (c *Client) handleMd5(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"a", "b", "c", "d", "m", "k"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, Md5Hint, message.MessageID)
		return
	}
	a := params["a"]
	b := params["b"]
	cc := params["c"]
	d := params["d"]
	m := params["m"]
	k := params["k"]

	if a <= 0 || b <= 0 || cc < 0 || d < 0 || m < 0 || k < 0 {
		c.sendMessage(message.Chat.ID, Md5Hint, message.MessageID)
		return
	}

	result := problems.Md5(a, b, cc, d, m, k)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
