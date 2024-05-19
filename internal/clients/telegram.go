package clients

import (
	"context"

	"github.com/go-telegram/bot"
)

type TelegramClient interface {
	SendMessage(
		ctx context.Context,
		receiverId int64,
		message string,
		mediaContent *string,
		disableNotification bool,
	) error
}

type telegramClient struct {
	b *bot.Bot
}

func NewTelegramClient(token string) (TelegramClient, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, err
	}
	return &telegramClient{b: b}, nil
}

func (c *telegramClient) SendMessage(
	ctx context.Context,
	receiverId int64,
	message string,
	mediaContent *string,
	disableNotification bool,
) error {
	_, err := c.b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:              receiverId,
		Text:                message,
		DisableNotification: disableNotification,
	})
	return err
}
