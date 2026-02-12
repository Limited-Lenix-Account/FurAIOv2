package flow

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// POSTs the Group ID and returns an Int Status Code, Not Much Happening Here
func GetHome(task *UserTask) int {

	d := fmt.Sprintf(`groupTypeId=%s&accessCode=&_csrf=%s`, task.AttendeeInfo.Value, task.Client.CSRF)
	data := strings.NewReader(d)
	url := fmt.Sprintf("%s/home/group", task.Hotel.StemURL)

	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; AWSALBTG=DVkU1WWmY7is/u5yJPhbUXcTSeRgM/gcl69hlXSiacUVkKPyGV6y3HGbDEhVhd9R11H6Ou2Mr39arGUXA4Og895HYNgmFccedjQJ17E6ksWDluwU/664Bq2W6bYdpWZem5RoeMzi9AX+W2rtADw65ER6Pf8HawHBhf5gks+Ivg5XR7qfShw=; AWSALBTGCORS=DVkU1WWmY7is/u5yJPhbUXcTSeRgM/gcl69hlXSiacUVkKPyGV6y3HGbDEhVhd9R11H6Ou2Mr39arGUXA4Og895HYNgmFccedjQJ17E6ksWDluwU/664Bq2W6bYdpWZem5RoeMzi9AX+W2rtADw65ER6Pf8HawHBhf5gks+Ivg5XR7qfShw=; AWSALB=QQ4vMTDfHjCVyxvuNOM6BGQ1BmIIQ6ZJx83akysoWfs1Uzv4VpJuvtXqYFrQncBSal3Z/mM8FmMHthAJgba0cPUaPgzKIjAYBRlN4hQaHrhcVT+u1AMmq8lowtBV; AWSALBCORS=QQ4vMTDfHjCVyxvuNOM6BGQ1BmIIQ6ZJx83akysoWfs1Uzv4VpJuvtXqYFrQncBSal3Z/mM8FmMHthAJgba0cPUaPgzKIjAYBRlN4hQaHrhcVT+u1AMmq8lowtBV; JSESSIONID=721wR7JRODw6L9vorY7RzEfiOSZI42AUB1OJaREI.10-96-56-148")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Referer", task.Hotel.StemURL+"/home")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
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

	fmt.Println(resp.Request.URL)

	return resp.StatusCode
}
