package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const FeistelHint = "/feistel data=047b16c2 keys=67e99b3c"

func (c *Client) handleFeistel(message *tgbotapi.Message) {

	requiredStrKeys := []string{"data", "keys"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, FeistelHint, message.MessageID)
		return
	}

	data := strParams["data"]
	keys := strParams["keys"]

	if len(data) != 8 || len(keys) != 8 {
		c.logger.Errorf("handleFeistel: invalid Feistel data=%s, keys=%s | %d, %d", data, keys, len(data), len(keys))
		c.sendMessage(message.Chat.ID, FeistelHint, message.MessageID)
		return
	}

	result, err := problems.Feistel(data, keys)
	if err != nil {
		c.logger.Errorf("feistelErr: %v", err)
		c.sendMessage(message.Chat.ID, FeistelHint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
