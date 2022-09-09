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