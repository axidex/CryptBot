package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DiffieHellmanHint = "Usage: /diffie n=17 q=11 x=2 y=3"

func (c *Client) handleDiffie(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"n", "q", "x", "y"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, DiffieHellmanHint, message.MessageID)
		return
	}

	n := params["n"]
	q := params["q"]
	x := params["x"]
	y := params["y"]

	result := problems.DiffieHellman(n, q, x, y)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
