package engines

import (
	"BotInfoDemon/botdata"
	"BotInfoDemon/telegram"
	"log"
)

const (
	MAX_MES_NUMBER = 10
)

func SiteEng(eng Eng_conf) []botdata.Bot_data {
	/*
		Get configuration struct with actions and their key words
		and make botdata
		return Bot_data slice

	*/
	result := []botdata.Bot_data{}

	if len(eng.ActionList) != len(eng.ConfList) {
		log.Fatalf("Len of ActionList %d and ConfList %d not the same!", len(eng.ActionList), len(eng.ConfList))
	}
	for i := range eng.ActionList {
		f := eng.ActionList[i].(func(botdata.Sconf) []botdata.Bot_data)

		result = append(result, f(eng.ConfList[i])...)

	}
	if len(result) == 0 {
		result = append(result, botdata.Empty_news)
		result[0].Topic = eng.Topic
	}

	return result
}

func TgEng(eng Eng_conf) {
	/*
		Get configuration struct with actions and their key words
		and make botdata and send today users by MAX_MES_NUMBER number messages

	*/
	rez := SiteEng(eng)

	usersID := telegram.GetUsers()
	for {
		if len(rez) > MAX_MES_NUMBER {
			sub_rez := rez[:MAX_MES_NUMBER]
			rez = rez[MAX_MES_NUMBER:]
			for i := range usersID {
				telegram.Send_mes(sub_rez, usersID[i])
			}

		} else {
			for i := range usersID {
				telegram.Send_mes(rez, usersID[i])
			}
			break
		}
	}

}
