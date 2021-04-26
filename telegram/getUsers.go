package telegram

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func GetUsers() (usersID []int) {
	/*
		Get users from current day, who write /start
	*/

	body := getBodyByUrl(getUrlByMethod(methodGetUpdetes))
	getupdates := GetUpdatesT{}

	err := json.Unmarshal(body, &getupdates)
	if err != nil {
		fmt.Println("Unmarshal err: ", err.Error())
		return
	}

	for _, update := range getupdates.Result {
		if strings.ToLower(update.Message.Text) == "/start" && time.Unix(int64(update.Message.Date), 0).Day() == time.Now().Day() {
			var flag bool = true
			for _, val := range usersID {
				flag = (val != update.Message.Chat.ID)
			}
			if flag {
				usersID = append(usersID, update.Message.Chat.ID)
			}

		}

	}

	return

}
