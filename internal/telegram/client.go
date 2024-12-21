package telegram

import (
	"bot/internal/app"
	"bot/internal/config"
	"bot/pkg/logger"
	"github.com/go-resty/resty/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

//type IClient interface {
//	Run()
//	Stop()
//}

type Client struct {
	logger     logger.Logger
	bot        *tgbotapi.BotAPI
	commands   map[Command]func(*tgbotapi.Message)
	httpClient *resty.Client
	config     config.DesConfig
}

func NewClient(apiKey string, debug bool, logger logger.Logger, config config.DesConfig) app.App {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		logger.Fatal(err)
	}

	bot.Debug = debug

	logger.Infof("Authorized on account %s", bot.Self.UserName)

	client := &Client{
		logger:     logger,
		bot:        bot,
		commands:   make(map[Command]func(*tgbotapi.Message)),
		httpClient: resty.New(),
		config:     config,
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
	c.commands[CommandDes] = c.handleDes
	c.commands[CommandAes] = c.handleAes
	c.commands[CommandInvMix] = c.handleInvMix
	c.commands[CommandFeistel] = c.handleFeistel
	c.commands[CommandHash] = c.handleHash
	c.commands[CommandA5] = c.handleA5
	c.commands[Command3Des] = c.handle3Des
	c.commands[CommandBlowFish] = c.handleBlowFish
	c.commands[CommandSBlock] = c.handleSBlock
	c.commands[CommandMd5] = c.handleMd5
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
		c.logger.Infof("Unknown command: %s", command)
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
