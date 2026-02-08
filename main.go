package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file!")
	}

	b, err := bot.New(
		os.Getenv("TELEGRAM_BOT_TOKEN"),
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			log.Debug().Any("update", update).Send()
			if update.Message != nil {
				chatID := update.Message.Chat.ID
				messageID := update.Message.ID
				replyParameters := models.ReplyParameters{ChatID: chatID, MessageID: messageID}
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:          chatID,
					Text:            fmt.Sprintf("Your chat ID is `%d`", chatID),
					ParseMode:       models.ParseModeMarkdown,
					ReplyParameters: &replyParameters,
				})

				if update.Message.ForwardOrigin != nil {
					if update.Message.ForwardOrigin.MessageOriginUser != nil {
						userID := update.Message.ForwardOrigin.MessageOriginUser.SenderUser.ID
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID:          chatID,
							Text:            fmt.Sprintf("The message forwarded from user with ID `%d`", userID),
							ParseMode:       models.ParseModeMarkdown,
							ReplyParameters: &replyParameters,
						})
					}

					if update.Message.ForwardOrigin.MessageOriginChat != nil {
						forwardChatID := update.Message.ForwardOrigin.MessageOriginChat.SenderChat.ID
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID:          chatID,
							Text:            fmt.Sprintf("The message forwarded from chat with ID `%d`", forwardChatID),
							ParseMode:       models.ParseModeMarkdown,
							ReplyParameters: &replyParameters,
						})
					}

					if update.Message.ForwardOrigin.MessageOriginChannel != nil {
						userID := update.Message.ForwardOrigin.MessageOriginChannel.Chat.ID
						b.SendMessage(ctx, &bot.SendMessageParams{
							ChatID:          chatID,
							Text:            fmt.Sprintf("The message forwarded from channel with chat ID `%d`", userID),
							ParseMode:       models.ParseModeMarkdown,
							ReplyParameters: &replyParameters,
						})
					}
				}
			}
		}),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot!")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}
