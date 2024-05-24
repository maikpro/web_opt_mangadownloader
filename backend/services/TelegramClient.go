package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maikpro/web_opt_mangadownloader/database"
	"github.com/maikpro/web_opt_mangadownloader/models"
)

func SendChapter(chapter models.Chapter) error {
	bot, err := getBot()
	if err != nil {
		log.Fatal(err)
		return err
	}

	sendMessage(bot, fmt.Sprintf("One Piece chapter: %d - %s - Pages: %d", chapter.Number, chapter.Name, len(chapter.Pages)))

	var mediaGroup []interface{}
	for _, page := range chapter.Pages {
		inputMediaPhoto := tgbotapi.NewInputMediaPhoto(page.Url)
		mediaGroup = append(mediaGroup, inputMediaPhoto)
	}

	err = sendMediaGroup(bot, mediaGroup)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func sleep(ms uint) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func getBot() (*tgbotapi.BotAPI, error) {
	telegramAPITokenString, err := getToken()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(*telegramAPITokenString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return bot, nil
}

func getToken() (*string, error) {
	settings, err := database.GetSettings()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &settings.TelegramToken, nil
}

func getChatId() (*int64, error) {
	settings, err := database.GetSettings()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	chatId, err := strconv.ParseInt(settings.TelegramChatId, 10, 64)
	if err != nil {
		return nil, err
	}

	return &chatId, nil
}

func sendMessage(bot *tgbotapi.BotAPI, text string) error {
	chatId, err := getChatId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	msg := tgbotapi.NewMessage(*chatId, text)
	bot.Send(msg)
	sleep(5500)
	return nil
}

func sendMediaGroup(bot *tgbotapi.BotAPI, mediaGroup []interface{}) error {
	chatId, err := getChatId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	chunkSize := 10
	for i := 0; i < len(mediaGroup); i += chunkSize {
		end := i + chunkSize
		if end > len(mediaGroup) {
			end = len(mediaGroup)
		}
		chunk := mediaGroup[i:end]
		mediaGroupConfig := tgbotapi.NewMediaGroup(*chatId, chunk)
		_, err := bot.Send(mediaGroupConfig)
		if err != nil {
			log.Fatal("Error sending mediaGroup:", err)
			return err
		}
		sleep(5500)
	}
	return nil
}
