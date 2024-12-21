package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const HashHint = "/hash word=сотовыйтелефон h0=7 p=11"

func (c *Client) handleHash(message *tgbotapi.Message) {

	requiredStrKeys := []string{"word"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, HashHint, message.MessageID)
		return
	}
	requiredIntKeys := []string{"h0", "p"}
	intParams, err := parseIntParams(message.Text, requiredIntKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, HashHint, message.MessageID)
		return
	}

	word := strParams["word"]
	h0 := intParams["h0"]
	p := intParams["p"]

	if len(word) == 0 || h0 < 0 || p < 0 {
		c.logger.Errorf("handleHash: invalid Hash word=%s, h0=%d, p=%d", word, h0, p)
		c.sendMessage(message.Chat.ID, HashHint, message.MessageID)
		return
	}

	result, err := problems.Hash(word, h0, p)
	if err != nil {
		c.logger.Errorf("hashErr: %v", err)
		c.sendMessage(message.Chat.ID, HashHint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
