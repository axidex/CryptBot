package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const EnigmaHint = "Usage: /diffie n=17 q=11 x=2 y=3"

func (c *Client) handleEnigma(message *tgbotapi.Message) {
	// Parse parameters from the message text
	requiredIntKeys := []string{"rot"}
	intParams, err := parseIntParams(message.Text, requiredIntKeys)
	if err != nil {
		c.logger.Errorf("parseIntParams: %v", err)
		c.sendMessage(message.Chat.ID, EnigmaHint, message.MessageID)
		return
	}

	requiredStrKeys := []string{"ref", "pp", "str"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, EnigmaHint, message.MessageID)
		return
	}

	rot := intParams["rot"]
	ref := strParams["ref"]
	pp := strParams["pp"]
	str := strParams["str"]

	result, err := problems.Enigma(rot, ref, str, pp)
	if err != nil {
		c.logger.Errorf("%d\t%s\t%s\t%s\t ", rot, ref, pp, str)
		c.logger.Error(err)
		c.sendMessage(message.Chat.ID, EnigmaHint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
