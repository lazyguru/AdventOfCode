package download

import (
	"errors"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getDay(day int) string {
	if day > 9 {
		return fmt.Sprintf("%d", day)
	}
	return fmt.Sprintf("0%d", day)
}

func downloadInput(year, day int) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	resp := getResponse(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Warning: Could not authenticate with AOC server. Resp Code: %d - %s", resp.StatusCode, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	sday := getDay(day)
	ensurePathExists(year, sday)
	filename := fmt.Sprintf("%d/%s/input.txt", year, sday)
	writeFile(filename, string(body))
}

func writeFile(filename string, body string) {
	f := getFile(filename)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error closing file", err)
		}
	}(f)
	_, err := f.WriteString(body)
	if err != nil {
		log.Fatal("Error writing to file", err)
	}
}

func getFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func getResponse(url string) *http.Response {
	client := http.Client{Timeout: time.Duration(30) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  os.Getenv("AOC_SESSION"),
		MaxAge: 300,
	}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	return resp
}

func downloadPuzzle(year, day int) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	resp := getResponse(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Warning: Could not authenticate with AOC server. Resp Code: %d - %s", resp.StatusCode, resp.Status)
	}
	return processHtml(resp.Body)
}

func processHtml(r io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	html, err := doc.Find("main").Html()
	if err != nil {
		log.Fatal(err)
	}
	return html
}

func SetupPuzzle(year, day, part int) {
	if part == 1 {
		downloadInput(year, day)
	}
	puzzle := downloadPuzzle(year, day)
	markdown := convertToMarkdown(puzzle)
	savePuzzle(year, day, markdown)
}

func savePuzzle(year int, day int, markdown string) {
	sday := getDay(day)
	filename := fmt.Sprintf("%d/%s/puzzle.md", year, sday)
	writeFile(filename, markdown)
}

func convertToMarkdown(html string) string {
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	return markdown
}

func ensurePathExists(year int, sday string) {
	dirname := fmt.Sprintf("%d/%s", year, sday)
	if _, err := os.Stat(dirname); err == nil || !errors.Is(err, os.ErrNotExist) {
		return
	}
	err := os.MkdirAll(dirname, 0750)
	if err != nil {
		log.Fatalf("Error creating directory %s: %s", dirname, err)
	}
}
