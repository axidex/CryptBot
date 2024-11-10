package telegram

import (
	"bot/internal/app"
	"bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

//type IClient interface {
//	Run()
//	Stop()
//}

type Client struct {
	logger   logger.Logger
	bot      *tgbotapi.BotAPI
	commands map[Command]func(*tgbotapi.Message)
}

func NewClient(apiKey string, debug bool, logger logger.Logger) app.App {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		logger.Fatal(err)
	}

	bot.Debug = debug

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	client := &Client{
		logger:   logger,
		bot:      bot,
		commands: make(map[Command]func(*tgbotapi.Message)),
	}
	// Register commands
	client.registerCommands()

	return client
}

func (c *Client) registerCommands() {
	c.commands[CommandStart] = c.handleStart
	c.commands[CommandHelp] = c.handleHelp
	c.commands[CommandElGamal] = c.handleElGamal
	c.commands[CommandDiffie] = c.handleDiffie
	c.commands[CommandRSA] = c.handleRSA
	c.commands[CommandEnigma] = c.handleEnigma
}

func (c *Client) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			c.logger.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
			c.processMessage(update.Message)
		}
	}
}

func (c *Client) processMessage(message *tgbotapi.Message) {
	command := GetCommand(strings.Split(strings.ToLower(message.Text), " ")[0])

	if handler, exists := c.commands[command]; exists {
		handler(message)
	} else {
		c.defaultReplyToUser(message)
	}
}

func (c *Client) defaultReplyToUser(message *tgbotapi.Message) {
	reply := "I don't recognize that command. Use /help to see available commands."
	c.sendMessage(message.Chat.ID, reply, message.MessageID)
}

func (c *Client) Stop() {
	c.logger.Infof("Stopping receiving updates from bot")
	c.bot.StopReceivingUpdates()
}
