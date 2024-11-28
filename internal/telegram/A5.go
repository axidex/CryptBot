package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const A5Hint = "/A5 a=010011 b=1001011 c=10101011110 text=f34ff541f04fed5ee04fff59fa50"

func (c *Client) handleA5(message *tgbotapi.Message) {

	requiredStrKeys := []string{"a", "b", "c", "text"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, FeistelHint, message.MessageID)
		return
	}

	a := strParams["a"]
	b := strParams["b"]
	cc := strParams["c"]
	text := strParams["text"]

	if len(text)%2 != 0 || len(text) == 0 {
		c.logger.Errorf("handleA5: invalid A5 a=%s, b=%s, c=%s, text=%s", a, b, cc, text)
		c.sendMessage(message.Chat.ID, A5Hint, message.MessageID)
		return
	}

	result, err := problems.A5(a, b, cc, text)
	if err != nil {
		c.logger.Errorf("A5Err: %v", err)
		c.sendMessage(message.Chat.ID, A5Hint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
