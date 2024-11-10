package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Client) sendMessage(chatID int64, text string, replyTo int) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyTo
	_, err := c.bot.Send(msg)
	if err != nil {
		c.logger.Fatal(err)
	}
	//c.logger.Infof("[%s]", send.From.UserName)
}
