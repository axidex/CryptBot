package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const AesHint = "/aes text=Перу key=Ключ"

func (c *Client) handleAes(message *tgbotapi.Message) {

	requiredStrKeys := []string{"text", "key"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, AesHint, message.MessageID)
		return
	}

	text := strParams["text"]
	key := strParams["key"]

	if len(text) != 8 || len(key) != 8 {
		c.logger.Errorf("handleAes: invalid AES text=%s, key=%s | %d, %d", text, key, len(text), len(key))
		c.sendMessage(message.Chat.ID, AesHint, message.MessageID)
		return
	}

	result := problems.Aes(text, key)

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
