package actions

import (
	"BotInfoDemon/botdata"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Get_Gnews(conf botdata.Sconf) (result []botdata.Bot_data) {
	/*
		Get config data with key words and parse Gnews generated rss
		return Bot_data slice
	*/
	const (
		b_url   = "https://news.google.com/rss/search?q="
		sep     = "%20"
		end_url = "&hl=ru&gl=RU&ceid=RU%3Aru"
	)

	body := getBody(b_url, end_url, sep, conf)
	defer body.Close()
	_ = body

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	// rss
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		var botd botdata.Bot_data
		botd.Topic = conf.Topic
		botd.Title = s.Find("title").Text()
		re := regexp.MustCompile(`.+ -`)
		botd.Title = string(re.Find([]byte(botd.Title)))
		botd.Title = botd.Title[:len(botd.Title)-2]

		botd.Text = s.Find("description").Text()
		re = regexp.MustCompile(`href=".+target`)
		botd.Link = string(re.Find([]byte(botd.Text)))
		botd.Link = botd.Link[6 : len(botd.Link)-8]

		re = regexp.MustCompile(`color=.+</font>`)
		botd.Source = string(re.Find([]byte(botd.Text)))
		botd.Source = botd.Source[16 : len(botd.Source)-7]

		re = regexp.MustCompile(`">.+</a>`)
		botd.Text = string(re.Find([]byte(botd.Text)))
		botd.Text = botd.Text[2 : len(botd.Text)-4]
		// botd.Text = ""

		tt := s.Find("pubDate").Text()
		timer, _ := time.Parse(time.RFC1123, tt)
		botd.Time = timer.Local().String()

		if conf.Clock.Add(-time.Hour * 10).Before(timer) {
			// fmt.Println(botd)

			result = append(result, botd)
		}
	})

	if len(result) == 0 {
		fmt.Println("Нет новостей: ", time.Now())
	}

	return result

}
