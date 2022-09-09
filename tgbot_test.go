package tgbot

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	token:=os.Getenv("TG_TOKEN")
	chatID:=os.Getenv("CHAT_ID")
	id,err:=strconv.Atoi(chatID)
	assert.NoError(t,err)
	tg, err := NewTgBot(int64(id), token)
	assert.NoError(t, err)
	err=tg.Send("test")
	assert.NoError(t,err)
}
