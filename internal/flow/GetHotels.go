package flow

import (
	"FurAIOIgnited/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetAllHotels(task *UserTask) (int, util.HotelSearchResult) {

	url := fmt.Sprintf("%s/list/hotels/all", task.Hotel.StemURL)
	// fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=721wR7JRODw6L9vorY7RzEfiOSZI42AUB1OJaREI.10-96-56-148; AWSALBTG=WCPUtPIgQPHwposv8j2cISQI/WDSkgAod4Who/bOzEQ9z0ZtZeWIIsgVyI/pjQsWDv4qRvlsUb5gfeZ7TbPavaRKMXkdfP76A6oZboI0HQb5f8JJMK6ScT7RMaodk0t+Zeyjho2VSitMhufjix3qlj8QeTHpA6SW/a5v4IB+x9FdvShcZno=; AWSALBTGCORS=WCPUtPIgQPHwposv8j2cISQI/WDSkgAod4Who/bOzEQ9z0ZtZeWIIsgVyI/pjQsWDv4qRvlsUb5gfeZ7TbPavaRKMXkdfP76A6oZboI0HQb5f8JJMK6ScT7RMaodk0t+Zeyjho2VSitMhufjix3qlj8QeTHpA6SW/a5v4IB+x9FdvShcZno=; AWSALB=dHiRExjjRLEcFDB5bhrHE9A4EAx91JcP6B+zBFRtprhd4Eyex7Rl/RayUbUMYS/qz+1Fny2ezwL3zPFeAVQVdLpE2civhXWo3H5Y7pXy5qX9EkJ7vUYQ8OB2IFex; AWSALBCORS=dHiRExjjRLEcFDB5bhrHE9A4EAx91JcP6B+zBFRtprhd4Eyex7Rl/RayUbUMYS/qz+1Fny2ezwL3zPFeAVQVdLpE2civhXWo3H5Y7pXy5qX9EkJ7vUYQ8OB2IFex")
	req.Header.Set("DNT", "1")
	req.Header.Set("Referer", task.Hotel.StemURL+"/landing")
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

	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error getting Hotel body into doc")
	}

	hotels := util.ParseHotels(doc)
	return resp.StatusCode, hotels

}

func GetAllRooms(task *UserTask) (int, util.HotelSearchResult) {

	url := fmt.Sprintf("%s/rooms/list?sort=default", task.Hotel.StemURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=721wR7JRODw6L9vorY7RzEfiOSZI42AUB1OJaREI.10-96-56-148; AWSALBTG=WCPUtPIgQPHwposv8j2cISQI/WDSkgAod4Who/bOzEQ9z0ZtZeWIIsgVyI/pjQsWDv4qRvlsUb5gfeZ7TbPavaRKMXkdfP76A6oZboI0HQb5f8JJMK6ScT7RMaodk0t+Zeyjho2VSitMhufjix3qlj8QeTHpA6SW/a5v4IB+x9FdvShcZno=; AWSALBTGCORS=WCPUtPIgQPHwposv8j2cISQI/WDSkgAod4Who/bOzEQ9z0ZtZeWIIsgVyI/pjQsWDv4qRvlsUb5gfeZ7TbPavaRKMXkdfP76A6oZboI0HQb5f8JJMK6ScT7RMaodk0t+Zeyjho2VSitMhufjix3qlj8QeTHpA6SW/a5v4IB+x9FdvShcZno=; AWSALB=dHiRExjjRLEcFDB5bhrHE9A4EAx91JcP6B+zBFRtprhd4Eyex7Rl/RayUbUMYS/qz+1Fny2ezwL3zPFeAVQVdLpE2civhXWo3H5Y7pXy5qX9EkJ7vUYQ8OB2IFex; AWSALBCORS=dHiRExjjRLEcFDB5bhrHE9A4EAx91JcP6B+zBFRtprhd4Eyex7Rl/RayUbUMYS/qz+1Fny2ezwL3zPFeAVQVdLpE2civhXWo3H5Y7pXy5qX9EkJ7vUYQ8OB2IFex")
	req.Header.Set("DNT", "1")
	req.Header.Set("Referer", task.Hotel.StemURL+"/list/hotels/all")
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

	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error getting Hotel body into doc")
	}

	hotels := util.ParseHotels(doc)

	return resp.StatusCode, hotels

}

func SendUpdate(task *UserTask) int {

	url := fmt.Sprintf("%s/rooms/select/update", task.Hotel.StemURL)

	dataJson := util.CreateUpdateJson(task.UserHotel.HotelID)
	data, err := json.Marshal(dataJson)
	if err != nil {
		fmt.Println("Update Struct Coudld not marshall")
	}

	var postData = strings.NewReader(string(data))
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
	req.Header.Set("Referer", task.Hotel.StemURL+"/list/hotels/all")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=08QW2J9NEVieKbzzk25aT8bYwUsORR69zkMnOwfp.10-96-57-30; AWSALBTG=MuhKhovLm5X2QP+1Qg2jvye5oEA+UddF2rV+TFU18YDpnVKpDdn9rQmmQ5HAQA6lR7fQbG9BFl75DePPm5jsZVQUntamzQnmBwImdPzaCmY7v/HGO/1axBhvZeg3ulb0M+bNB5Nq4xfotqDW605fr7f3czWyml6HbHTJvwNq5mcNBnxcRDU=; AWSALBTGCORS=MuhKhovLm5X2QP+1Qg2jvye5oEA+UddF2rV+TFU18YDpnVKpDdn9rQmmQ5HAQA6lR7fQbG9BFl75DePPm5jsZVQUntamzQnmBwImdPzaCmY7v/HGO/1axBhvZeg3ulb0M+bNB5Nq4xfotqDW605fr7f3czWyml6HbHTJvwNq5mcNBnxcRDU=; AWSALB=r5+W6/D6Lx3uTN/rxGQ9ciXMUgiUKq251rCU/NmYMdSNeALNLdOOMk2ABt6fYvCJmZTjt/VFsJxMjk1CVt83KDGnat6Jnz600h829Wo9103Rj0lsKCe8zXfnuyqo; AWSALBCORS=r5+W6/D6Lx3uTN/rxGQ9ciXMUgiUKq251rCU/NmYMdSNeALNLdOOMk2ABt6fYvCJmZTjt/VFsJxMjk1CVt83KDGnat6Jnz600h829Wo9103Rj0lsKCe8zXfnuyqo")
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

	return resp.StatusCode

}
