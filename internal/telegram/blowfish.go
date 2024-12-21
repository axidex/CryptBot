package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const BlowFishHint = "/blowfish l=4 r=9 k1=1 k2=5 k3=3 k4=4 k5=5"

func (c *Client) handleBlowFish(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredKeys := []string{"l", "r", "k1", "k2", "k3", "k4", "k5"}
	params, err := parseIntParams(message.Text, requiredKeys)
	if err != nil {
		c.sendMessage(message.Chat.ID, BlowFishHint, message.MessageID)
		return
	}

	l := params["l"]
	r := params["r"]
	k1 := params["k1"]
	k2 := params["k2"]
	k3 := params["k3"]
	k4 := params["k4"]
	k5 := params["k5"]

	if l <= 0 || r <= 0 || k1 < 0 || k2 < 0 || k3 < 0 || k4 < 0 || k5 < 0 {
		c.sendMessage(message.Chat.ID, BlowFishHint, message.MessageID)
		return
	}

	result := problems.BlowFish(l, r, k1, k2, k3, k4, k5)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
