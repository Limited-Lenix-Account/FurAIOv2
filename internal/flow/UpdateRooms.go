package flow

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func UpdateTotal(updtjson UpdateStruct, task *UserTask) {

	data, err := json.Marshal(updtjson)
	if err != nil {
		fmt.Println("Update Json could not be updated")
	}

	fmt.Println(string(data))

	var postData = strings.NewReader(string(data))
	url := fmt.Sprintf("%s/rooms/select/update?updateTotals=false", task.Hotel.StemURL)
	req, err := http.NewRequest("POST", url, postData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-XSRF-TOKEN", task.Client.CSRF)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", task.Hotel.StemURL+"/rooms/list?sort=default")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=lgNfJIPy_6BfNf20eJ_UQjZr8bmND-NRFJovb3NN.10-96-56-120; AWSALBTG=Wy0YZcBBeuZmp17ewaRUckuMHkJ30F3Yqf1GpcRqizDtiZmCioYk3NcGsnV4W8yBn1iMNoYvw1JiEctAIwAs0BKRS0qckTwTZhGXz0NJLeX1kQ6dElDOwEAdTQIgH55WKMoWcJMht3vvHDE2f2VKz1mT1mSC4PFg9+bgr7wItor5WKqT4W4=; AWSALBTGCORS=Wy0YZcBBeuZmp17ewaRUckuMHkJ30F3Yqf1GpcRqizDtiZmCioYk3NcGsnV4W8yBn1iMNoYvw1JiEctAIwAs0BKRS0qckTwTZhGXz0NJLeX1kQ6dElDOwEAdTQIgH55WKMoWcJMht3vvHDE2f2VKz1mT1mSC4PFg9+bgr7wItor5WKqT4W4=; AWSALB=/krNdQTBX3rQJKJGr9YZ/gKog57Zogdh7FW+gfReAWDUkWD4Jgi0pWDi+4zDiHxYF3aUV3m32d1UjX76rhEO/+3N18CXQwctqyroYShKzF/vRRc8I9Xdyi3HRmrb; AWSALBCORS=/krNdQTBX3rQJKJGr9YZ/gKog57Zogdh7FW+gfReAWDUkWD4Jgi0pWDi+4zDiHxYF3aUV3m32d1UjX76rhEO/+3N18CXQwctqyroYShKzF/vRRc8I9Xdyi3HRmrb")
	resp, err := task.Client.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}
