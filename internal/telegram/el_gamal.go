package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ElGamalHint = "/elgamal p=19 g=3 k=7 x=8 M=5"

func (c *Client) handleElGamal(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"p", "g", "k", "x", "M"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, ElGamalHint, message.MessageID)
		return
	}

	// Perform ElGamal calculations
	p := params["p"]
	g := params["g"]
	k := params["k"]
	x := params["x"]
	M := params["M"]

	result := problems.ElGamal(p, g, k, x, M)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
