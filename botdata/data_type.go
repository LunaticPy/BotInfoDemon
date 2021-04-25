package botdata

import (
	"fmt"
	"time"
)

type Bot_data struct {
	Topic  string
	Title  string
	Time   string
	Text   string
	Source string
	Link   string
}

func (d Bot_data) String() string {
	return fmt.Sprintf("title: %s\nTime: %s\ntext: %s\nsource: %s\nlink: %s\n", d.Title, d.Time, d.Text, d.Source, d.Link)
}

type Sconf struct {
	Topic     string
	Key_words []string
	Clock     time.Time
}

var Empty_news = Bot_data{"", "", "", "", "Нет новостей по данной теме", ""}
var Last_news = Bot_data{"", "", "", "", "Есть еще, но их слишком много", ""}
var Sep_news = Bot_data{"", "", "", "", "%0AЕсть еще, смотри дальше", ""}
