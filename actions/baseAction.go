package actions

import (
	"BotInfoDemon/botdata"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getBody(b_url, end_url, sep string, conf botdata.Sconf) io.ReadCloser {
	/*
		Make io.ReadCloser from url parts and configurations
		return io.ReaderCloser
	*/

	key := strings.Join(conf.Key_words, sep)
	buf := []string{b_url, key, end_url}
	URL := strings.Join(buf, "")
	fmt.Println(URL)
	ans, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}

	if ans.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", ans.StatusCode, ans.Status)
	}

	return ans.Body

}
