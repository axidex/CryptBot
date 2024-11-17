package telegram

import (
	"bot/internal/problems"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const InvMixHint = "/invMix char=h"

func (c *Client) handleInvMix(message *tgbotapi.Message) {

	requiredStrKeys := []string{"char"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, InvMixHint, message.MessageID)
		return
	}

	char := strParams["char"]

	if len(char) != 1 {
		c.logger.Errorf("handleInvMix: invalid length char %q", char)
		c.sendMessage(message.Chat.ID, InvMixHint, message.MessageID)
		return
	}

	invMix := problems.InvMixCol{}
	err = invMix.Solve(char)
	if err != nil {
		c.logger.Errorf("handleInvMix: %v", err)
		c.sendMessage(message.Chat.ID, InvMixHint, message.MessageID)
		return
	}

	c.sendMessage(message.Chat.ID, invMix.GetSolution(), message.MessageID)
}
