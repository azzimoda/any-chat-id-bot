# Any Chat ID Bot

This Telegram bot answers to any message, including forwarded ones, with your and its sender's chat IDs.

The bot can answer:
- your chat ID — to any message,
- chat ID of the user which sent the forwarded message,
- chat ID of the channel which the message is forwarded from.

The bot hosted by me: [@any_chat_id_bot](https://any_chat_id_bot.t.me)

---

## Technology stack

- Go 1.25
  - `go-telegram/bot`

## How to run it

1. Install go of version 1.25.x or higher.
2. Clone the repository.
3. Create a `.env` file with the following content:
   ```
   TELEGRAM_BOT_TOKEN=your_bot_token
   ```
4. Run the bot:
   ```
   go run main.go
   ```

---

## Contacts

My Telegram channel [RU]: [@mazzaLLM](https://mazzaLLM.t.me)
