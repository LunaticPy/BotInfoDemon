package telegram

import (
	"BotInfoDemon/botdata"
	"strconv"
)

func Send_mes(mes []botdata.Bot_data, chatID int) {
	/*
		Send Bot_data to chatID
	*/

	curTopic := mes[0].Topic

	text := "Тема: " + curTopic + getPeriod()

	for i := range mes {
		if mes[i].Topic != curTopic {
			curTopic = mes[i].Topic
			text = text + "Тема: " + curTopic + Newline
		}
		text = text + getSingleTextBlock(mes[i])
	}

	url := getUrlByMethod(methodSendMessage)
	url = url + "?chat_id=" + strconv.Itoa(chatID) + "&text=" + text + "&parse_mode=HTML&disable_web_page_preview=1"
	getBodyByUrl(url)

}
