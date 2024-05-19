package clients

import "telegram-notification-api/internal/config"

type Clients interface {
	TelegramClient() TelegramClient
}

type client struct {
	telegramClient TelegramClient
}

func (c *client) TelegramClient() TelegramClient {
	return c.telegramClient
}

func NewClients(config config.Config) (Clients, error) {
	tg, err := NewTelegramClient(config.MustGetTelegramBotToken())
	if err != nil {
		return nil, err
	}

	return &client{
		telegramClient: tg,
	}, nil
}
