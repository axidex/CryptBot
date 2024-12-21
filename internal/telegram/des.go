package telegram

import (
	"github.com/axidex/CryptBot/internal/problems"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const DesHint = "/des text=гана key=куб"

func (c *Client) handleDes(message *tgbotapi.Message) {
	requiredStrKeys := []string{"text", "key"}
	strParams, err := parseStrParams(message.Text, requiredStrKeys)
	if err != nil {
		c.logger.Errorf("parseStrParams: %v", err)
		c.sendMessage(message.Chat.ID, DesHint, message.MessageID)
		return
	}

	text := strParams["text"]
	key := strParams["key"]

	textBinaryResult, err := problems.TransformText(text)
	if err != nil {
		c.logger.Errorf("TransformText: %v", err)
		c.sendMessage(message.Chat.ID, DesHint, message.MessageID)
		return
	}

	keyBinaryResult, err := problems.TransformText(key)
	if err != nil {
		c.logger.Errorf("TransformText: %v", err)
		c.sendMessage(message.Chat.ID, DesHint, message.MessageID)
		return
	}

	hexText := problems.BinToHexStr(textBinaryResult.Value)
	hexKey := problems.BinToHexStr(keyBinaryResult.Value)

	resp, err := c.httpClient.R().
		SetQueryParam("key", hexKey).
		SetQueryParam("data", hexText).
		Post(fmt.Sprintf("http://%s:%d/encrypt", c.config.DesHost, c.config.DesPort))
	if err != nil {
		c.logger.Errorf("Des: %v", err)
		c.sendMessage(message.Chat.ID, DesHint, message.MessageID)
		return
	}
	if resp.StatusCode() != 200 {
		c.logger.Errorf("Des: %v | %d", resp.Body(), resp.StatusCode())
		c.sendMessage(message.Chat.ID, DesHint, message.MessageID)
		return
	}

	trimmedResp := strings.Trim(resp.String(), `"`)
	normalResp := strings.ReplaceAll(trimmedResp, `\n`, "\n")

	c.sendMessage(message.Chat.ID, normalResp, message.MessageID)
}
