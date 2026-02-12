package flow

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SubmitInfo(encodedStr string, task *UserTask) {

	var data = strings.NewReader(encodedStr)
	req, err := http.NewRequest("POST", task.Hotel.StemURL+"/guest/info", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=Bj-hdAlXnM0Nk7qqDQi9qRQpdUOuNHg_9nnkJMMu.10-96-56-148; AWSALBTG=F5jL73XakALqyOsXMGVwnL1YRdQoovEzMTg6MWJEGzRZAfWWf1hKsDoAkoQ3WN2JB3pMv4cSIwBLv5f+mmuYJek7e/6ew5NSiDNh8ghAfMutb5ij11gF6nOjgZtpIOOtuw/h57Egnm359F9xe5YNdd2+FsCci4THouoXVWaMmTycYSZb/bU=; AWSALBTGCORS=F5jL73XakALqyOsXMGVwnL1YRdQoovEzMTg6MWJEGzRZAfWWf1hKsDoAkoQ3WN2JB3pMv4cSIwBLv5f+mmuYJek7e/6ew5NSiDNh8ghAfMutb5ij11gF6nOjgZtpIOOtuw/h57Egnm359F9xe5YNdd2+FsCci4THouoXVWaMmTycYSZb/bU=; AWSALB=8uFdZTICl0kjSYx3YrVWJMyHhyDtW7kWLcnEF/09QMa+eI/xxnLwM5RoRoTBWormUVAuUwurBitny24at9H3SBNGove7/GBN4Ux6/L7BiJWrRmz+j7RqN+q5oP8L; AWSALBCORS=8uFdZTICl0kjSYx3YrVWJMyHhyDtW7kWLcnEF/09QMa+eI/xxnLwM5RoRoTBWormUVAuUwurBitny24at9H3SBNGove7/GBN4Ux6/L7BiJWrRmz+j7RqN+q5oP8L")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Referer", task.Hotel.StemURL+"/guest/info")
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

}

//fbed38c6-f759-468c-a2b4-33e8e8a86a85

func SubmitPayment(task *UserTask, payment string) {

	var data = strings.NewReader(payment)
	req, err := http.NewRequest("POST", task.Hotel.StemURL+"/guest/payment", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", task.Hotel.StemURL+"/guest/payment")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=rXwacYZ0D-ydhxWIaZNcLiMMLk51mGU2K1Dj9NPv.10-96-56-87; AWSALBTG=ghKRwNI9MBpe+Ouzb5rI/s0+0KcnbO3CTYMG+5y/EZfRtPXLhxyCb7lkazMC2albY0u7Q4xnOAFOODO2/thsPiPAF4EomWLaofkrc8slvTCSFC0ksqKiiBqbNDx9ibMKn5bXo6q3dgDiI+dkhB1YXmISZvmSYliBlfGTOYw9LSAsDS3D9Fg=; AWSALBTGCORS=ghKRwNI9MBpe+Ouzb5rI/s0+0KcnbO3CTYMG+5y/EZfRtPXLhxyCb7lkazMC2albY0u7Q4xnOAFOODO2/thsPiPAF4EomWLaofkrc8slvTCSFC0ksqKiiBqbNDx9ibMKn5bXo6q3dgDiI+dkhB1YXmISZvmSYliBlfGTOYw9LSAsDS3D9Fg=; AWSALB=QsFFi65q4/F1K38YaARYYiqCdT4xC3vS75LdqvdupJNSjeq//9ctJ/dwOvNQxEYl7TA4EfrXQ8A5sNun92yD5Vur0ot0PWPlb/v0vkGRDhNVk+0Lt3DySZGLn/Gk; AWSALBCORS=QsFFi65q4/F1K38YaARYYiqCdT4xC3vS75LdqvdupJNSjeq//9ctJ/dwOvNQxEYl7TA4EfrXQ8A5sNun92yD5Vur0ot0PWPlb/v0vkGRDhNVk+0Lt3DySZGLn/Gk")
	resp, err := task.Client.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	fmt.Println(resp.Request.URL)

}

func SubmitReservation(task *UserTask) {

	var data = strings.NewReader(`attendeeConsentAgreements%5B0%5D.consentGiven=1&_attendeeConsentAgreements%5B0%5D.consentGiven=on&attendeeConsentAgreements%5B0%5D.consentGiven=0&_csrf=` + task.Client.CSRF)
	req, err := http.NewRequest("POST", task.Hotel.StemURL+"/reservation/save", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", task.Hotel.StemURL+"/summary")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=IQ_YbbQXqjchvJKrf8xKEdyE8svlIN3OkXzZQeAS.10-96-56-87; AWSALBTG=qkW5D1umRELC1hEAjGhcP4SKgy713TaBklkVaBgFMCsGOMB1XuskbdjLr7EZz2MOeo2eqFvR6whExqBDmlg3Zm81Mv3KZx84g/CYgWNxCRs/vzxRk1sA6+wBMy8YybUtxXOv+g5aP1ZSykFb/hqbHedjGTf2Hm6+OwlJzJkx8gMgnokWDU4=; AWSALBTGCORS=qkW5D1umRELC1hEAjGhcP4SKgy713TaBklkVaBgFMCsGOMB1XuskbdjLr7EZz2MOeo2eqFvR6whExqBDmlg3Zm81Mv3KZx84g/CYgWNxCRs/vzxRk1sA6+wBMy8YybUtxXOv+g5aP1ZSykFb/hqbHedjGTf2Hm6+OwlJzJkx8gMgnokWDU4=; AWSALB=SHH76bGLi/qWB8KCnhlQl5jISFRDt0YLaxXrIPw6ZvMUzvspYXkpb9TQwoiZiw/VbU8q5OoLRs9dEIMCARm6b+hR0Aq1e9HanWlwlH7MUGscruB6WWr7g14dSlLt; AWSALBCORS=SHH76bGLi/qWB8KCnhlQl5jISFRDt0YLaxXrIPw6ZvMUzvspYXkpb9TQwoiZiw/VbU8q5OoLRs9dEIMCARm6b+hR0Aq1e9HanWlwlH7MUGscruB6WWr7g14dSlLt")
	resp, err := task.Client.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	fmt.Println(resp.Request.URL)

}
