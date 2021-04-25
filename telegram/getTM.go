package telegram

import (
	"BotInfoDemon/botdata"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var TelegramToken string = ""

const (
	Newline           = "%0A"
	telegramBaseUrl   = "https://api.telegram.org/bot"
	methodSendMessage = "sendMessage"
	methodGetUpdetes  = "getUpdates"
)

type GetUpdatesT struct {
	Ok     bool                `json: "ok"`
	Result []GetUpdatesResultT `json: "result"`
}

type GetUpdatesResultT struct {
	UpdateID int                `json: "update_id"`
	Message  GetUpdatesMessageT `json: "massage,omitempty"`
}

type GetUpdatesMessageT struct {
	MessageId int `json: "message_id"`
	From      struct {
		ID           int    `json: "id"`
		FirstName    string `json: "first_name"`
		LastName     string `json: "last_name"`
		Username     string `json: "username"`
		LanguageCode string `json: "language_code"`
	} `json: "from"`
	Chat struct {
		ID        int    `json: "id"`
		FirstName string `json: "first_name"`
		LastName  string `json: "last_name"`
		Username  string `json: "username"`
		Type      string `json: "type"`
	} `json: "chat"`

	Date int    `json: "date"`
	Text string `json: "text"`
}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + TelegramToken + "/" + methodName
}

func getSingleTextBlock(mes botdata.Bot_data) string {
	if mes.Title == "" {
		return mes.Source
	}
	textALL := "<b>Источник: " + mes.Source +
		"</b>%0AСсылка: <a href=\"" + mes.Link +
		"\">" + mes.Title +
		"</a>%0A<code>Время: " + mes.Time[:len(mes.Time)-10] +
		//  "%0AПревью:" + mes.Text +
		"</code>" + Newline + Newline

	return textALL
}

func getBodyByUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}

func getPeriod() string {
	end_time := time.Now()
	st_time := end_time.Add(-time.Minute * 30)

	return " с " + st_time.Format(time.Stamp)[7:12] + " по " + end_time.Format(time.Stamp)[7:12] + Newline
}

func GetToken(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

// func getMessageJson(text string, chatID int) []byte {
// 	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text":%s, "parse_mode":"HTML", "disable_web_page_preview":False}`, chatID, text))
// 	tx := bytes.NewReader(data)
// 	url := getUrlByMethod(methodSendMessage)
// 	response, err := http.Post(url, "application/json", tx)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	return body

// }
