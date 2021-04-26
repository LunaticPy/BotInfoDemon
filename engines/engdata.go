package engines

import (
	"BotInfoDemon/actions"
	"BotInfoDemon/botdata"
	"time"
)

type Eng_conf struct {
	Topic      string
	ActionList []interface{}
	ConfList   []botdata.Sconf
}

var TestConf = []botdata.Sconf{
	botdata.Sconf{
		Topic:     "Covid-19",
		Key_words: []string{"Moskva", "COVID-19"},
		Clock:     time.Now()},
	botdata.Sconf{Topic: "Covid-19",
		Key_words: []string{"Moskva", "bolnitsa", "rospotrebnadzor"},
		Clock:     time.Now()},
}

var TestAct = []interface{}{
	actions.Get_Gnews,
	actions.Get_Gnews,
}

var DefaultEng = Eng_conf{
	Topic:      "Covid-19",
	ActionList: TestAct,
	ConfList:   TestConf}
