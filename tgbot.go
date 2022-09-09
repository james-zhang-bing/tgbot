package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *TgBot

func init() {

}

type TgBot struct {
	ChatID int64
	BotToken string
	api      *tgbotapi.BotAPI
}

func NewTgBot(chatid int64, token string) (*TgBot, error) {
	if token==""||chatid==0{
		return nil,fmt.Errorf("tgbot required chat id and bot's token")
	}
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot := &TgBot{
		ChatID: chatid,
		BotToken: token,
		api:      api,
	}
	
	return bot, bot.Send("start")
}

func (l *TgBot) Send(msg string) error {
	tmsg := tgbotapi.NewMessage(l.ChatID, msg)
	_,err:=l.api.Send(tmsg)
	return err
}
