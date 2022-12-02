package submit

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func SubmitAnswer(year, day, part int, answer string) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	resp := postResponse(url, fmt.Sprintf("level=%d&answer=%s", part, answer))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Error: Could not submit answer. Response: %s", resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := doc.Find("main").Text()
	log.Println(text)
}

func postResponse(url, data string) *http.Response {
	client := http.Client{
		Timeout: time.Duration(30) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	body := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 300,
	}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return resp
}
