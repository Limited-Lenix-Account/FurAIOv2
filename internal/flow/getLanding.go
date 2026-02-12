package flow

import (
	"FurAIOIgnited/util"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Will Get Passkey Landing Page and returns Int Status Code and A List of Attendee Options
func GetLanding(task *UserTask) (int, []util.AttendeeOptions) {
	task.Hotel.StemURL = fmt.Sprintf("https://book.passkey.com/event/%s/owner/%s", task.Hotel.EventID, task.Hotel.OwnerID)

	url := fmt.Sprintf("https://book.passkey.com/event/%s/owner/%s/home", task.Hotel.EventID, task.Hotel.OwnerID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := task.Client.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	task.Client.CSRF = util.GetCSRFToken(task.Client.client.Jar)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Cannot Read Resp into Body")
	}
	opts := util.GetAttendeeOptions(doc)
	return resp.StatusCode, opts

}
