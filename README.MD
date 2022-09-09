# telegram robot msg sender

It's used for send teglegram message to your self 


## Example 
``` go
func main(){
    youChatId:=12345678
    token:="52646xxxx:xxxxxxxxxxxxdKJ1gATeWEPgotSe7M"
    tg, err := NewTgBot(youChatId, token)
	assert.NoError(t, err)
	err=tg.Send("hello")
	assert.NoError(t,err)
}
```
## how to get your chat id and token

- chat id: Forward a message to the [bot](https://t.me/userinfobot),it will send you the user id
- token:  Just talk to [BotFather](https://t.me/BotFather) (described below) and follow a few simple steps.