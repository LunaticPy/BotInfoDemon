package main

import (
	"BotInfoDemon/engines"
	"BotInfoDemon/telegram"
)

func main() {

	telegram.TelegramToken = telegram.GetToken("bottoken.txt")

	engines.TgEng(engines.DefaultEng)

}
