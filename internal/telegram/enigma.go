package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const EnigmaHint = "/enigma rot=315 ref=C pp=A-U,D-Y str=MJL"

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

	if len(str) > 20 {
		c.sendMessage(message.Chat.ID, EnigmaHint, message.MessageID)
		return
	}

	result, err := problems.Enigma(rot, ref, str, pp)
	if err != nil {
		c.logger.Errorf("%d\t%s\t%s\t%s\t ", rot, ref, pp, str)
		c.logger.Error(err)
		c.sendMessage(message.Chat.ID, EnigmaHint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, result, message.MessageID)
}
