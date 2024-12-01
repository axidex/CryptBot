package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const RSAHint = "/rsa p=3 q=7 e=5 x=101"

func (c *Client) handleRSA(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"p", "q", "e", "x"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, RSAHint, message.MessageID)
		return
	}

	p := params["p"]
	q := params["q"]
	e := params["e"]
	x := params["x"]

	if p < 0 || q < 0 || e < 0 || x < 0 {
		c.sendMessage(message.Chat.ID, RSAHint, message.MessageID)
		return
	}

	result := problems.RSA(p, q, e, x)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
