package app

import (
	"net/http"

	"github.com/m1al04949/news-bot/internal/config"
	"github.com/m1al04949/news-bot/internal/lib/xmlparser/handlers"
	"github.com/m1al04949/news-bot/internal/pkg/setlog"
	"golang.org/x/exp/slog"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var rss = map[string]string{
	"Habr": "https://habrahabr.ru/rss/best/",
}

func RunBot() error {
	cfg := config.MustLoad()

	log := setlog.SetupLogger(cfg.Env)

	log.Info("starting news_bot", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	log.Debug("Authorized on account", slog.String("username", bot.Self.UserName))

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.WebhookURL))
	if err != nil {
		return err
	}

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(cfg.Port, nil)
	log.Info("start listen", slog.String("port", cfg.Port))

	// получаем обновления из канала updates
	for update := range updates {
		if url, ok := rss[update.Message.Text]; ok {
			rss, err := handlers.GetNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"sorry, error happend",
				))
			}
			for _, item := range rss.Items {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					item.URL+"\n"+item.Title,
				))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				`there is only Habr feed availible`,
			))
		}
	}

	return nil
}
