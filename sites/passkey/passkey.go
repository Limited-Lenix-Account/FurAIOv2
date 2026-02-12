package passkey

import (
	task "FurAIOIgnited/cmd/taskengine"
	"FurAIOIgnited/internal/discord"
	"FurAIOIgnited/sites/queueit"
	"FurAIOIgnited/util"

	"encoding/base64"
	"encoding/json"

	"net/http"
	"net/url"

	"fmt"
	"html"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/PuerkitoBio/goquery"
)

// TODO: Add more handling to GetHotel function, and add retrying n stuff
// Right now it just crashes when it tries to make a survey of an empty list lmfao
// Check for list w/ current endpoint -> check for list with room endpoint -> return err and loop if all is OOS

type PasskeyTask struct {
	task.Task
	GivenURL string

	CheckoutDateIn  string
	CheckoutDateOut string

	HotelPosKws []string
	HotelNegKws []string

	Flow string

	EventName  string
	HotelName  string
	HotelImage string
	BlockName  string
	OrderLink  *url.URL
	AckNumber  string

	EventID    string
	OwnerID    string
	HotelID    int
	HotelIndex int
	BlockID    int

	GroupID      string
	GroupTitle   string
	CSRF         string
	StemURL      string
	Consent      string
	PaymentFolio string

	GroupOptions []util.AttendeeOptions
	HotelOptions util.HotelSearchResult
	UpdateBlock  UpdateStruct

	EncodedInfo    string
	EncodedPayment string

	Subtotal  string
	Charge    string
	TaxAmount string

	UQueue       string
	SoftBlockUrl *url.URL

	UserAgent string
	SourceIP  string

	QueueEnv []queueit.QueueSession

	QueueUserID     string
	QueueSession    string
	QueueID         string
	ChalComplexity  int
	ChalInput       string
	ChalZeros       int
	ChalDetails     string
	EncodedSolution string
	Tries           int
	Enqueued        bool
}

// TODO: Add function to change profile dates to the following format
// TODO: Add Handling when hotels are OOS

func (e *PasskeyTask) Ignite() {

	//Parse keywords from user profile
	p, n := util.GetKws(e.Profile.HotelKeywords)
	e.HotelPosKws = p
	e.HotelNegKws = n

	e.UpdateStatus(Start)

	if e.Flow == "Booking" {
		e.ParseBookingURL()
	}

	e.StartTime = time.Now()

	//Ideally have list of user agents to use for queue, but idrk if they'll care
	e.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	for e.Task.Stage != "Stop" {
		switch e.Flow {
		case "Booking":
			switch e.Task.Stage {
			case Start:
				e.GetLanding()
				e.GetHome()

			case GetHotels:
				e.GetAllHotels()

			case Update:
				e.SendUpdate()

			case GetRooms:
				e.GetAllRooms()

			case MakeBlock:
				e.BuildBlock()

			case SubmitInfo:
				e.SubmitInformation()
			case SubmitTravel:
				e.SubmitTravel()

			case SubmitPayment:
				e.SubmitPayment()

			case SubmitReservation:
				e.SubmitReserv()

			}
		case "Queue":
			switch e.Task.Stage { // Once Thru queue sends to the start of Booking
			case Start:
				e.GetQueuePage()

			case PowChallenge:
				e.GetPowChallenge()

			case StartPow:
				e.SolvePow()

			case SubmitSolution:
				e.SubmitSolution()

			case CheckQueue:
				e.InQueue()

			}
		}
	}
}

func (e *PasskeyTask) ParseBookingURL() {

	e.EventID, e.OwnerID, _ = util.GetOwnerHotel(e.GivenURL)

	if e.EventID == "" || e.OwnerID == "" {

		req, err := http.NewRequest("GET", e.GivenURL, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Connection", "keep-alive")
		// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; cookieConsent_50648914=ESSENTIAL; JSESSIONID=mBRalTAYc5qgT0lC52LrfD-4y1FeAm_dK-E1Ra3t.10-96-56-134; AWSALBTG=Nt0N/lc59ujTuO6tHho539HzU/24F2s1WRSgQepJq9N4mc0XJ3m9rDcVjMm+HalrRFmWsEys5E7RHM4DVxIjomaRW8Vz4B+Szyu+6fnd13lVduqSEEOJGww5mBDzL845ERUWGjIRWxpPUK5sFVdJQicjCz3S+nukZaR70YZtyR3ycxbH3R8=; AWSALBTGCORS=Nt0N/lc59ujTuO6tHho539HzU/24F2s1WRSgQepJq9N4mc0XJ3m9rDcVjMm+HalrRFmWsEys5E7RHM4DVxIjomaRW8Vz4B+Szyu+6fnd13lVduqSEEOJGww5mBDzL845ERUWGjIRWxpPUK5sFVdJQicjCz3S+nukZaR70YZtyR3ycxbH3R8=; AWSALB=NX1uMC7vczN1gg/mBo8bHdWb1D/FgNbIuX7rMES5zkGZIWjQwnQ8iQmHYaD2USvaXVHfdBI64xx11h8ep+vBDaWVVpjzV9MISUrmBHtEL+hq6WCfLShIARp4C9um; AWSALBCORS=NX1uMC7vczN1gg/mBo8bHdWb1D/FgNbIuX7rMES5zkGZIWjQwnQ8iQmHYaD2USvaXVHfdBI64xx11h8ep+vBDaWVVpjzV9MISUrmBHtEL+hq6WCfLShIARp4C9um")
		req.Header.Set("DNT", "1")
		// req.Header.Set("Referer", "https://www.furryfiesta.org/")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "cross-site")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
		req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"macOS"`)
		resp, err := e.Task.Client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("Error Parsing HTML")
		}

		// fmt.Println(doc)

		e.EventID, e.OwnerID = util.ParseIDfromJS(doc)

	} else {
		// fmt.Println("IDs Found, starting checkout...")
		// e.Task.Stage = Start
		e.UpdateStatus(Start)
	}

	if e.EventID != "" && e.OwnerID != "" {
		// fmt.Println("IDs Found, starting checkout...")
		// e.Task.Stage = Start
	} else {
		fmt.Println("No IDs Found")
		e.Task.Stage = Stop
	}

}

func (e *PasskeyTask) GetLanding() {

	// fmt.Printf("Event ID: %s, Owner ID: %s \n", e.EventID, e.OwnerID)

	e.StemURL = fmt.Sprintf("https://book.passkey.com/event/%s/owner/%s", e.EventID, e.OwnerID)

	url := fmt.Sprintf("https://book.passkey.com/event/%s/owner/%s/home", e.EventID, e.OwnerID)

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
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	e.CSRF = util.GetCSRFToken(e.Task.Client.Jar)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Cannot Read Resp into Body")
	}

	opts := util.GetAttendeeOptions(doc)

	// See if there's only One GroupID
	if len(opts) == 1 {
		e.GroupID = opts[0].Value // Jump to get home here
	} else if len(opts) == 3 { // If theres a 0 or -1
		var validOpts []util.AttendeeOptions
		for _, v := range opts {
			if v.Value != "0" && v.Value != "-1" {
				validOpts = append(validOpts, v) // Append only valid GroupIDs
			}
		}
		if len(validOpts) == 1 {
			e.GroupID = validOpts[0].Value // If there is only one valid one, assign it go the task GroupID

		} else {
			e.GroupOptions = validOpts // Otherwise only assign valid objects to the task
		}
	} else {
		e.GroupOptions = opts
		e.SurveyGroup() // Ask Survey
	}

	// Add Survey here and then jump to get home...

}

func (e *PasskeyTask) GetHome() {

	d := fmt.Sprintf(`groupTypeId=%s&accessCode=&_csrf=%s`, e.GroupID, e.CSRF)
	data := strings.NewReader(d)
	url := fmt.Sprintf("%s/home/group", e.StemURL)

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
	req.Header.Set("Referer", e.StemURL+"/home")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	e.EventName = util.GetEventTitle(doc)

	e.UpdateStatus(GetHotels)

}

// If there is only one hotel, the /list/hotels/all and /rooms/list/ URLs give the same information
func (e *PasskeyTask) GetAllHotels() {

	url := fmt.Sprintf("%s/list/hotels/all", e.StemURL)
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
	req.Header.Set("Referer", e.StemURL+"/landing")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error getting Hotel body into doc")
	}

	hotels := util.ParseHotels(doc)
	e.HotelOptions = hotels

	if len(hotels) == 0 {
		fmt.Println("All Hotels Out of Stock, monitoring...")
		e.Task.Stage = GetHotels
		time.Sleep(3 * time.Second)
	} else if e.Task.Mode == "Auto" {
		for i := range e.HotelOptions {
			// fmt.Println(e.HotelOptions[i].Name)
			if util.ContainsKeywords(e.HotelOptions[i].Name, e.HotelPosKws, e.HotelNegKws) {
				e.HotelID = e.HotelOptions[i].ID
				e.HotelName = e.HotelOptions[i].Name
				e.HotelImage = e.HotelOptions[i].ImageURL
				e.HotelIndex = i

				e.UpdateStatus(Update)
				return
			}
		}

	} else {
		e.SurveyHotel()
	}

}

func (e *PasskeyTask) SendUpdate() {

	url := fmt.Sprintf("%s/rooms/select/update", e.StemURL)

	dataJson := util.CreateUpdateJson(e.HotelID)
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
	req.Header.Set("X-XSRF-TOKEN", e.CSRF)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.StemURL+"/list/hotels/all")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=08QW2J9NEVieKbzzk25aT8bYwUsORR69zkMnOwfp.10-96-57-30; AWSALBTG=MuhKhovLm5X2QP+1Qg2jvye5oEA+UddF2rV+TFU18YDpnVKpDdn9rQmmQ5HAQA6lR7fQbG9BFl75DePPm5jsZVQUntamzQnmBwImdPzaCmY7v/HGO/1axBhvZeg3ulb0M+bNB5Nq4xfotqDW605fr7f3czWyml6HbHTJvwNq5mcNBnxcRDU=; AWSALBTGCORS=MuhKhovLm5X2QP+1Qg2jvye5oEA+UddF2rV+TFU18YDpnVKpDdn9rQmmQ5HAQA6lR7fQbG9BFl75DePPm5jsZVQUntamzQnmBwImdPzaCmY7v/HGO/1axBhvZeg3ulb0M+bNB5Nq4xfotqDW605fr7f3czWyml6HbHTJvwNq5mcNBnxcRDU=; AWSALB=r5+W6/D6Lx3uTN/rxGQ9ciXMUgiUKq251rCU/NmYMdSNeALNLdOOMk2ABt6fYvCJmZTjt/VFsJxMjk1CVt83KDGnat6Jnz600h829Wo9103Rj0lsKCe8zXfnuyqo; AWSALBCORS=r5+W6/D6Lx3uTN/rxGQ9ciXMUgiUKq251rCU/NmYMdSNeALNLdOOMk2ABt6fYvCJmZTjt/VFsJxMjk1CVt83KDGnat6Jnz600h829Wo9103Rj0lsKCe8zXfnuyqo")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	var UpResp util.UpdateResp

	err = json.Unmarshal(bodyText, &resp)
	if err != nil {
		fmt.Println("Cannot unmarshall update")
	}

	e.Subtotal = fmt.Sprintf("%.2f", UpResp.TotalWithTaxes)
	e.TaxAmount = fmt.Sprintf("%.2f", UpResp.TaxAmount)

	if resp.StatusCode == 200 {
		e.Task.Stage = GetRooms
	}

}

func (e *PasskeyTask) GetAllRooms() {

	url := fmt.Sprintf("%s/rooms/list?sort=default", e.StemURL)

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
	req.Header.Set("Referer", e.StemURL+"/list/hotels/all")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error getting Hotel body into doc")
	}

	// https://book.passkey.com/gt/219470785?gtid=6668d41b8820983da16d5720ab74dc9b
	e.HotelOptions = util.ParseHotels(doc)

	if e.Task.Mode == "Manual" {
		e.SurveyRooms()
		var AvailBlock [][]int
		var FormattedBlock []string

		var DateSelection []int

		for _, v := range e.HotelOptions {
			if v.ID == e.HotelID {
				for _, b := range v.Blocks {
					if b.ID == e.BlockID {

						for _, i := range b.Inventory {
							if i.Available != 0 {
								FormattedBlock = append(FormattedBlock, util.ConvertDate(i.Date))
								AvailBlock = append(AvailBlock, i.Date)
							}
						}
					}
				}
			}
		}

		prompt := &survey.MultiSelect{
			Message: "Select Your Days...",
			Options: FormattedBlock,
		}
		survey.AskOne(prompt, &DateSelection)
		e.Task.Profile.UserInformation.BookingInformation.CheckInDate, e.Task.Profile.UserInformation.BookingInformation.CheckOutDate, e.CheckoutDateIn, e.CheckoutDateOut = util.GetCheckoutDates(AvailBlock, DateSelection)
		e.Task.Stage = MakeBlock

	} else if e.Task.Mode == "Auto" {
		// Auto Mode will automatically search for rooms that match the kws AND date, instead of just picking the VERY first one
		// If the requested one isn't available just pick a random one then.

		// Check to Match KWs
		posKws, negKws := util.GetKws(e.Profile.RoomKeywords)
		allRooms := e.HotelOptions[e.HotelIndex].Blocks
		var validRooms util.Blocks

		//Gets All Valid Hotel Blocks
		for _, block := range allRooms {
			titleList := []string{} // Creates Empty List

			if block.AverageRate > 0.00 { // Only Rooms that are not Completely Sold out

				//Creates list of block title words
				splitTitle := strings.Split(html.UnescapeString(block.Name), " ")
				for _, title := range splitTitle {
					titleList = append(titleList, strings.ToLower(title))
				}

				// Checks to see if every pos kw is in the name, and if not a single neg one is in as well
				if util.ListInList(titleList, posKws) && !util.ListInList(titleList, negKws) {
					validRooms = append(validRooms, block)
				}
			}
		}

		e.BlockID = validRooms[0].ID
		e.BlockName = validRooms[0].Name

		if e.BlockName != "" {
			e.CheckoutDateIn, e.CheckoutDateOut = util.ConvertProfileDates(e.Profile.UserInformation.BookingInformation.CheckInDate, e.Profile.UserInformation.BookingInformation.CheckOutDate)
			e.UpdateStatus(MakeBlock)
		} else {
			fmt.Println("No Blocks Found")
		}

	} else if e.Mode == "Monitor" {
		validRooms := util.Blocks{}

		for _, v := range e.HotelOptions[e.HotelIndex].Blocks {

			if v.Charge > 0 {
				validRooms = append(validRooms, v)
			}
		}
		if len(validRooms) == 0 {
			fmt.Println("All Rooms OOS!")
		} else {
			for i := range validRooms {
				fmt.Println(validRooms[i].Name, validRooms[i].AverageRate)

				for _, d := range validRooms[i].Inventory {
					if d.Available != 0 {
						fmt.Println(d.Date)
					}

				}
			}
		}
		time.Sleep(3 * time.Second)
	}

}

func (e *PasskeyTask) BuildBlock() {

	b := []Block{{
		HotelID:          e.HotelID,
		BlockID:          e.BlockID,
		CheckIn:          e.Task.Profile.UserInformation.BookingInformation.CheckInDate,
		CheckOut:         e.Task.Profile.UserInformation.BookingInformation.CheckOutDate,
		NumberOfRooms:    1,
		NumberOfGuests:   1,
		NumberOfChildren: 0,
	}}

	bm := BlockMap{
		Blocks:      b,
		TotalGuests: 1,
		TotalRooms:  1,
	}

	up := UpdateStruct{
		HotelID:  e.HotelID,
		BlockMap: bm,
	}

	data, err := json.Marshal(up)
	if err != nil {
		fmt.Println("Update Json could not be updated")
	}

	var postData = strings.NewReader(string(data))
	url := fmt.Sprintf("%s/rooms/select/update?updateTotals=false", e.StemURL)
	req, err := http.NewRequest("POST", url, postData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-XSRF-TOKEN", e.CSRF)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.StemURL+"/rooms/list?sort=default")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	var UpdateResp util.UpdateResp

	err = json.Unmarshal(bodyText, &UpdateResp)
	if err != nil {
		fmt.Println("errer unmarshalling hotel block update resp", err)
	}

	e.Charge = fmt.Sprintf("%.2f", UpdateResp.Charge)
	e.TaxAmount = fmt.Sprintf("%.2f", UpdateResp.TaxAmount)
	e.Subtotal = fmt.Sprintf("%.2f", UpdateResp.TotalWithTaxes)

	if resp.StatusCode == 200 {

		e.UpdateStatus(SubmitInfo)
	} else {
		fmt.Println(resp.StatusCode)
		fmt.Println(string(data))
		fmt.Println(url)
		// bodyText, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%s\n", bodyText)

		// e.Task.Stage = Stop
		e.UpdateStatus(Stop)

	}

}

func (e *PasskeyTask) SubmitInformation() {

	e.EncodePersonalInformation()

	var data = strings.NewReader(e.EncodedInfo)
	req, err := http.NewRequest("POST", e.StemURL+"/guest/info", data)
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
	req.Header.Set("Referer", e.StemURL+"/guest/info")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// https://book.passkey.com/event/50383999/owner/49830583/guest/payment

	split := strings.Split(resp.Request.URL.String(), "/")

	fmt.Println(split)

	switch split[8] {
	case "payment":
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		e.PaymentFolio = util.GetPaymentVal(doc)

		if e.PaymentFolio != "" {
			e.UpdateStatus(SubmitPayment)
		} else {
			e.UpdateStatus(Stop)
		}
	case "travel":
		e.UpdateStatus(SubmitTravel)
	}

	// if split[8] == "payment" {

	// } else {
	// 	fmt.Println("Error Submitting User Info")
	// 	e.Task.Stage = Stop
	// }

}

func (e *PasskeyTask) SubmitTravel() {
	var data = strings.NewReader(`reservations%5B0%5D.guests%5B0%5D.id=0&reservations%5B0%5D.guests%5B0%5D.travelInfo.arrDateTime=&reservations%5B0%5D.guests%5B0%5D.travelInfo.depDateTime=&reservations%5B0%5D.guests%5B0%5D.travelInfo.arrAirline.id=0&reservations%5B0%5D.guests%5B0%5D.travelInfo.arrAirline.name=&reservations%5B0%5D.guests%5B0%5D.travelInfo.depAirline.id=0&reservations%5B0%5D.guests%5B0%5D.travelInfo.depAirline.name=&reservations%5B0%5D.guests%5B0%5D.travelInfo.arrFlightNum=&reservations%5B0%5D.guests%5B0%5D.travelInfo.depFlightNum=&reservations%5B0%5D.guests%5B0%5D.travelInfo.otherDetails=&_csrf=` + e.CSRF)
	req, err := http.NewRequest("POST", e.StemURL+"/guest/travel", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	// req.Header.Set("Referer", "https://book.passkey.com/event/50914548/owner/14285164/guest/travel")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "XSRF-TOKEN=4b2fbe45-5af9-4102-8bc4-b7c13fbff096; cookieConsent_50908645=ALL; cookieConsent_50909588=ALL; locale=en_US; cookieConsent_50910425=ESSENTIAL; cookieConsent_50911679=ALL; JSESSIONID=_Dqjvpi8OCXlktJB3o3vWUNhYKWMPQDakoTD6WOd.10-96-56-154; cookieConsent_50914548=ALL; AWSALBTG=jnf0Y4s04pKcORY6LXnjWmLeHg/m0jrTn44RB3V8uJ9/WynlJ9qhCb0CayIQTWorTu3TLsWHAa7sOUC0B2GMDc+lzmSrd3HHI5gNvXPVGEbOBYctvx5hgXqfF0+nnaE2o5RtRum+zcSj8Gl0UOUvMB5JJecy+As0ewfFai0kIzdVemJw4dY=; AWSALBTGCORS=jnf0Y4s04pKcORY6LXnjWmLeHg/m0jrTn44RB3V8uJ9/WynlJ9qhCb0CayIQTWorTu3TLsWHAa7sOUC0B2GMDc+lzmSrd3HHI5gNvXPVGEbOBYctvx5hgXqfF0+nnaE2o5RtRum+zcSj8Gl0UOUvMB5JJecy+As0ewfFai0kIzdVemJw4dY=; AWSALB=EfcLvvc9l7lnd6CFi6VDEnh2ZyYHlwEmjfWYnyqsIQxPWQyvO/dR8CHSlD6pfNgQswrv7PfqHln3OytEas2jQegioZY5+JkedOV8fejhR924UlGWRFxiPgx/JgdM; AWSALBCORS=EfcLvvc9l7lnd6CFi6VDEnh2ZyYHlwEmjfWYnyqsIQxPWQyvO/dR8CHSlD6pfNgQswrv7PfqHln3OytEas2jQegioZY5+JkedOV8fejhR924UlGWRFxiPgx/JgdM")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", body)
	split := strings.Split(resp.Request.URL.String(), "/")

	if split[8] == "payment" {
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		e.PaymentFolio = util.GetPaymentVal(doc)

		if e.PaymentFolio != "" {
			e.UpdateStatus(SubmitPayment)
		} else {
			e.UpdateStatus(Stop)
		}

	} else {
		fmt.Println("Error Submitting User Info")
		e.Task.Stage = Stop
	}

	e.UpdateStatus(SubmitPayment)
}

func (e *PasskeyTask) SubmitPayment() {

	e.EncodePayment()

	var data = strings.NewReader(e.EncodedPayment)
	req, err := http.NewRequest("POST", e.StemURL+"/guest/payment", data)
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
	req.Header.Set("Referer", e.StemURL+"/guest/payment")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "cookieConsent_50711677=ALL; XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; JSESSIONID=rXwacYZ0D-ydhxWIaZNcLiMMLk51mGU2K1Dj9NPv.10-96-56-87; AWSALBTG=ghKRwNI9MBpe+Ouzb5rI/s0+0KcnbO3CTYMG+5y/EZfRtPXLhxyCb7lkazMC2albY0u7Q4xnOAFOODO2/thsPiPAF4EomWLaofkrc8slvTCSFC0ksqKiiBqbNDx9ibMKn5bXo6q3dgDiI+dkhB1YXmISZvmSYliBlfGTOYw9LSAsDS3D9Fg=; AWSALBTGCORS=ghKRwNI9MBpe+Ouzb5rI/s0+0KcnbO3CTYMG+5y/EZfRtPXLhxyCb7lkazMC2albY0u7Q4xnOAFOODO2/thsPiPAF4EomWLaofkrc8slvTCSFC0ksqKiiBqbNDx9ibMKn5bXo6q3dgDiI+dkhB1YXmISZvmSYliBlfGTOYw9LSAsDS3D9Fg=; AWSALB=QsFFi65q4/F1K38YaARYYiqCdT4xC3vS75LdqvdupJNSjeq//9ctJ/dwOvNQxEYl7TA4EfrXQ8A5sNun92yD5Vur0ot0PWPlb/v0vkGRDhNVk+0Lt3DySZGLn/Gk; AWSALBCORS=QsFFi65q4/F1K38YaARYYiqCdT4xC3vS75LdqvdupJNSjeq//9ctJ/dwOvNQxEYl7TA4EfrXQ8A5sNun92yD5Vur0ot0PWPlb/v0vkGRDhNVk+0Lt3DySZGLn/Gk")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Request.URL)

	split := strings.Split(resp.Request.URL.String(), "/")

	///https://book.passkey.com/event/50383999/owner/49830583/summary

	if split[7] == "summary" {
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		e.Consent = util.GetAcklowedgement(doc, e.CSRF)
		e.UpdateStatus(SubmitReservation)
	} else {
		e.UpdateStatus(Stop)
	}

}

func (e *PasskeyTask) ResAggregate() {

	req, err := http.NewRequest("GET", e.StemURL+"/summary/resAggregate", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("X-XSRF-TOKEN", "fd4eb5be-1604-418c-a125-1178a91a61fd")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.StemURL+"/summary")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Cookie", "JSESSIONID=BC7cnhAgqZtkszotaukM3zQLKML9eD0jt70MaACX.10-96-56-30; XSRF-TOKEN=fd4eb5be-1604-418c-a125-1178a91a61fd; AWSALBTG=kcq9Cbrd+lAZK4WpCoh4DvlmTrxnvRD/W0z/0tSlf/wxrypiMAaRfgBFY21OKbM5Q6VU6X21sEyao328jHOf2uvp/fOXWq1Aj3oN/+AlWsIWskI/6VZMHzlnOaFKmxOyZk/kD30k00LhzauvBV+0zxcRnd//08Y8R//GxpZP9hcykpsfAgo=; AWSALBTGCORS=kcq9Cbrd+lAZK4WpCoh4DvlmTrxnvRD/W0z/0tSlf/wxrypiMAaRfgBFY21OKbM5Q6VU6X21sEyao328jHOf2uvp/fOXWq1Aj3oN/+AlWsIWskI/6VZMHzlnOaFKmxOyZk/kD30k00LhzauvBV+0zxcRnd//08Y8R//GxpZP9hcykpsfAgo=; AWSALB=2aWdfvgJpzrVeLX6CJRRFsOqZgxQKJGPBfsFDUOlumxmUbVIhhnI/SRdDzJTxtsTsI7sRLGb/SpPVcC0MYycUeeTBZvU3XP1kVTuTDD7hqwyHuxLsduz7ul8LQRN; AWSALBCORS=2aWdfvgJpzrVeLX6CJRRFsOqZgxQKJGPBfsFDUOlumxmUbVIhhnI/SRdDzJTxtsTsI7sRLGb/SpPVcC0MYycUeeTBZvU3XP1kVTuTDD7hqwyHuxLsduz7ul8LQRN")
	resp, err := e.Task.Client.Do(req)
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

func (e *PasskeyTask) SubmitReserv() {

	var data = strings.NewReader(e.Consent)

	req, err := http.NewRequest("POST", e.StemURL+"/reservation/save", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://book.passkey.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", e.StemURL+"/summary")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	split := strings.Split(resp.Request.URL.String(), "/")
	entry := split[3][:5]

	if entry == "entry" {

		e.OrderLink = resp.Request.URL

		tok := time.Since(e.StartTime)
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		e.AckNumber = util.ParseConfNumber(doc)

		if e.AckNumber != "" {
			util.SaveHTML(resp.Body)
			e.UpdateStatus(Successful)
			d := discord.SuccessfulHotelData{

				EventName:    e.EventName,
				HotelName:    e.HotelName,
				HotelImage:   e.HotelImage,
				BlockName:    e.BlockName,
				TotalCost:    e.Subtotal,
				ProfileEmail: e.Task.Profile.UserInformation.Email,

				StartDate:    e.Profile.UserInformation.BookingInformation.CheckInDate,
				EndDate:      e.Profile.UserInformation.BookingInformation.CheckOutDate,
				CheckoutTime: tok.Seconds(),
				OrderLink:    *resp.Request.URL,
				AckNumber:    e.AckNumber,
			}

			discord.HotelCheckoutWebhook(d)

			e.UpdateStatus(Stop)
		} else {
			fmt.Println("No Order Number Found!")
		}

	} else {
		// Takes you back to https://book.passkey.com/event/50770887/owner/50168607/list/hotels if dates are sold out
		fmt.Println("Error Placing Order, retrying...")
		util.SaveHTML(resp.Body)

		fmt.Println(resp.Request.URL)
		e.UpdateStatus(Stop)
	}

	// https://book.passkey.com/entry?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwYXlsb2FkIjp7ImVudHJ5IjoiUkVTRVJWQVRJT04iLCJwYXJhbXMiOlt7Im5hbWUiOiJjb25maXJtTnVtYmVyIiwidmFsdWUiOiJDWUZIU1RMQSJ9XX19.O3k0VrfyPWcGHd5EIk7A14aQpPhue_I3Pj98iJHcF6A

}

// Can't check order in the same session that you checkout in, need to go back and solve ReCap
func (e *PasskeyTask) CheckOrder() {

	req, err := http.NewRequest("GET", e.StemURL+"/summary/resAggregate", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "book.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-XSRF-TOKEN", e.CSRF)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.OrderLink.String())
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Cookie", "XSRF-TOKEN=f2e31237-dc49-40ca-963a-d39a364333e2; cookieConsent_50383999=ESSENTIAL; locale=en_US; cookieConsent_50648914=ESSENTIAL; cookieConsent_50480332=ESSENTIAL; cookieConsent_50756657=ESSENTIAL; cookieConsent_50770887=ESSENTIAL; cookieConsent_50770908=ESSENTIAL; cookieConsent_50749991=ESSENTIAL; JSESSIONID=JkyrrLHr4GqrAbqBVF5qmgJOTgzbblMkNiE9nrG5.10-96-56-122; AWSALBTG=efWc3Qukurj7Cs9UUgOgFWC+q6ugq6ou9+vRqWJTRnOXBTrHD9t8sNFFcfm11DJ46xHrDh5Beo+L6VUewFYShyzXd4kLD2GNqnrNUvTJ04h1Y79+TRz7o2vb8BxhD3W5VNsUnYsGifQaxPcQcfvV2UqDwuwejnzv+VHi4yQIC2DozHycn6Q=; AWSALBTGCORS=efWc3Qukurj7Cs9UUgOgFWC+q6ugq6ou9+vRqWJTRnOXBTrHD9t8sNFFcfm11DJ46xHrDh5Beo+L6VUewFYShyzXd4kLD2GNqnrNUvTJ04h1Y79+TRz7o2vb8BxhD3W5VNsUnYsGifQaxPcQcfvV2UqDwuwejnzv+VHi4yQIC2DozHycn6Q=; AWSALB=Q7BmTQGAtb7F4+Y6lFE+93xsMJKIb88GpBGIJLTHGyH5BdTil0sXe+xoopIFANHRre6ONKhSSUY/pfozHHs7qewKCVfwUP+l5JhHmGINJ6W2YMc0lfNdDlFIygei; AWSALBCORS=Q7BmTQGAtb7F4+Y6lFE+93xsMJKIb88GpBGIJLTHGyH5BdTil0sXe+xoopIFANHRre6ONKhSSUY/pfozHHs7qewKCVfwUP+l5JhHmGINJ6W2YMc0lfNdDlFIygei")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	var CheckOrder util.PasskeySuccessfulOrder
	err = json.Unmarshal(body, &CheckOrder)
	if err != nil {
		fmt.Println("Error Parsing Check Order ", err)
	}
	e.AckNumber = CheckOrder.Reservations[0].AckNumber

}

// Prompts the user to select a respective group, is not called when there is only one option to enter the site
func (e *PasskeyTask) SurveyGroup() {

	var optsTitle []string
	var index int
	for _, v := range e.GroupOptions {
		optsTitle = append(optsTitle, v.Title)
	}

	prompt := &survey.Select{
		Message: "Select Your Group:",
		Options: optsTitle,
	}

	survey.AskOne(prompt, &index)

	e.GroupID = e.GroupOptions[index].Value

}

// Prompts the user and sets just the HotelID and index (along with details for the webhook)
func (e *PasskeyTask) SurveyHotel() {

	var hotelTitle []string
	var index int

	for _, v := range e.HotelOptions {
		hotelTitle = append(hotelTitle, html.UnescapeString(v.Name))
	}

	prompt := &survey.Select{
		Message: "Select A Hotel: ",
		Options: hotelTitle,
	}
	survey.AskOne(prompt, &index)
	e.HotelIndex = index
	e.HotelID = e.HotelOptions[e.HotelIndex].ID

	e.HotelName = html.UnescapeString(e.HotelOptions[e.HotelIndex].Name)
	e.HotelImage = e.HotelOptions[e.HotelIndex].ImageURL

	fmt.Println("Sending Update...")
	e.Task.Stage = Update
}

// Uses the updated HotelOptions to set the blockID and blockName, as well sets (and overrides) the dates for the profile.
// Only used in a Manual Task
func (e *PasskeyTask) SurveyRooms() {
	var roomTitle []string
	var index int

	for _, v := range e.HotelOptions[e.HotelIndex].Blocks {
		roomTitle = append(roomTitle, html.UnescapeString(v.Name)+" - $"+strconv.Itoa(int(v.AverageBasicRate)))
	}

	prompt := &survey.Select{
		Message: "Select Which Room: ",
		Options: roomTitle,
	}

	survey.AskOne(prompt, &index)

	h := e.HotelOptions[e.HotelIndex]

	e.BlockID = h.Blocks[index].ID
	e.BlockName = h.Blocks[index].Name
	fmt.Printf("Selected: %s @ %s\n", h.Blocks[index].Name, h.Name)

	fmt.Println("Making Block...")
	e.Task.Stage = MakeBlock

}

func (e *PasskeyTask) GetQueuePage() {
	e.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	e.Tries = 1

	req, _ := http.NewRequest("GET", e.GivenURL, nil)

	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", e.UserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// fmt.Println(resp.Request.URL)

	u, _ := url.Parse("https://queue.passkey.com")

	for _, v := range e.Task.Client.Jar.Cookies(u) {
		if v.Name == "Queue-it" {
			// fmt.Println(v.Value)
			e.UQueue = strings.Split(v.Value, "=")[1]
			e.Task.Stage = Challenge
		}
	}

	e.SoftBlockUrl = resp.Request.URL
	if e.SoftBlockUrl.Host == "book.passkey.com" { // Check for automatic redirect
		fmt.Println("No Queue Found...")

		e.EventID, e.OwnerID, err = util.GetOwnerHotel(e.SoftBlockUrl.String())
		if err != nil {
			fmt.Printf("Error Parsing URL: %s", e.SoftBlockUrl.String())
		}
		fmt.Printf("Event: %s, Owner: %s\n", e.EventID, e.OwnerID)

		fmt.Println("Starting Booking")
		e.Flow = "Booking"
		e.Task.Stage = Start
		return
	} else {
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		e.QueueUserID = util.GetUserID(doc)
		fmt.Println(e.QueueUserID)
		e.Task.Stage = PowChallenge
	}
}

func (e *PasskeyTask) GetPowChallenge() {

	url := fmt.Sprintf("https://queue.passkey.com/challengeapi/pow/challenge/%s", e.QueueUserID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("X-Queueit-Challange-EventId", "anthrocontest2028")
	req.Header.Set("powTag-UserId", e.QueueUserID)
	req.Header.Set("X-Queueit-Challange-CustomerId", "lanyon")
	req.Header.Set("powTag-EventId", "anthrocontest2028")
	req.Header.Set("powTag-CustomerId", "lanyon")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", e.UserAgent)
	req.Header.Set("X-Queueit-Challange-Hash", "2qYfON4WHD2MggY3XsUpcjNZVlxlIM1gCcoZkXf7/8w=")
	req.Header.Set("powTag-Hash", "2qYfON4WHD2MggY3XsUpcjNZVlxlIM1gCcoZkXf7/8w=")
	req.Header.Set("X-Queueit-Challange-UserId", e.QueueUserID)
	req.Header.Set("X-Queueit-Challange-reason", "1")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.GivenURL)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// req.Header.Set("Cookie", "Queue-it-lanyon______________anthrocontest2028=Cid=en-US&f=0; Queue-it=u=d1e21a33-b2ba-4d73-bc40-ffdaefe7fea0")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var qChal queueit.ChallengeResponse

	err = json.Unmarshal(body, &qChal)
	if err != nil {
		fmt.Println("Error unmarshalling queueit pow challenge", err)
	}

	e.ChalComplexity = qChal.Parameters.Complexity
	e.ChalInput = qChal.Parameters.Input
	e.ChalZeros = qChal.Parameters.ZeroCount
	e.ChalDetails = qChal.ChallengeDetails
	e.QueueSession = qChal.SessionID

	if e.ChalComplexity == 0 || e.ChalZeros == 0 {
		fmt.Println("Zero Complexity, Zero Zeros")
		e.Task.Stage = Stop
	} else {
		e.Task.Stage = StartPow
	}

}

func (e *PasskeyTask) SolvePow() {

	sol, err := queueit.SolvePoW(e.ChalInput, e.ChalComplexity, e.ChalZeros)
	if err != nil {
		fmt.Println("Error Solving PoW")
	}

	solution := queueit.QueueItSolution{
		Hash: sol,
		Type: "HashChallenge",
	}

	solJson, _ := json.Marshal(solution)
	// fmt.Println(string(solJson))

	solStr := (solJson)
	e.EncodedSolution = base64.StdEncoding.EncodeToString(solStr)

	// e.Task.Stage = SubmitSolution

	e.SubmitSolution()

}

func (e *PasskeyTask) SubmitSolution() {

	s := queueit.Stats{
		UserAgent:      e.UserAgent,
		Screen:         "1440 x 900",
		Browser:        "Chrome",
		BrowserVersion: "121.0.0.0",
		IsMobile:       false,
		Os:             "Mac OS X",
		OsVersion:      "10_15_7",
		CookiesEnabled: true,
		Tries:          e.Tries,
		Duration:       23246,
	}

	v := queueit.QueueItVerify{
		ChallengeType:    "proofofwork",
		SessionID:        e.QueueSession,
		ChallengeDetails: e.ChalDetails,
		Solution:         e.EncodedSolution,
		Stats:            s,
		CustomerID:       "lanyon",
		EventID:          "anthrocontest2028",
		Version:          6,
	}

	postData, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error Marshalling QueueIt Verification")
	}

	var data = strings.NewReader(string(postData))
	req, err := http.NewRequest("POST", "https://queue.passkey.com/challengeapi/verify", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", e.UserAgent)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", e.SoftBlockUrl.String())
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var QueueSession queueit.VerifyResponse

	err = json.Unmarshal(body, &QueueSession)

	if err != nil {
		fmt.Println("Failed to unmarshal Solution Response", err)
		// fmt.Printf("%s\n", body)
	}

	if QueueSession.IsVerified {
		fmt.Println("Queue Verified!")
		e.QueueEnv = append(e.QueueEnv, QueueSession.SessionInfo)
		e.Task.Stage = CheckQueue
	} else {
		e.Task.Stage = Stop

	}

}

func (e *PasskeyTask) InQueue() {

	e.Enqueue()

	for e.Enqueued {
		e.PingQueue()
		time.Sleep(2000 * time.Millisecond)

		if e.Task.Stage == Stop {
			break
		}
	}

}

func (e *PasskeyTask) Enqueue() {

	c := queueit.Enqueue{
		ChallengeSessions: e.QueueEnv,
		LayoutName:        "Cvent Styles - Messages - Passkey",
		CustomURLParams:   "",
		TargetURL:         "",
		Referrer:          "",
	}

	d, _ := json.Marshal(c)

	var data = strings.NewReader(string(d))
	req, err := http.NewRequest("POST", "https://queue.passkey.com/spa-api/queue/lanyon/anthrocontest2028/enqueue?cid=en-US", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://queue.passkey.com/?c=lanyon&e=anthrocontest2028")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "Queue-it=u=b0639742-9f2a-458f-8c8d-c0e63627fa66")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	var EnqResp queueit.EnqueueResp

	err = json.Unmarshal(body, &EnqResp)
	if err != nil {
		fmt.Println("Error Unmarshalling Enqueue Response")
	}
	if !EnqResp.ChallengeFailed {
		e.Enqueued = true
		e.QueueID = EnqResp.QueueID

	} else {
		e.Enqueued = false
	}
}

func (e *PasskeyTask) PingQueue() {

	url := fmt.Sprintf("https://queue.passkey.com/spa-api/queue/lanyon/anthrocontest2028/%s/status?", e.QueueID)
	// url := fmt.Sprintf("https://queue.passkey.com/spa-api/queue/lanyon/anthrocontest2028/c99d9731-a438-4f88-aeed-facba374235b/status?cid=en-US&l=Cvent%20Styles%20-%20Messages%20-%20Passkey&seid=5e5af792-58bb-9583-2efd-781d2c267fde&sets=1707286424999")

	var data = strings.NewReader(`{"targetUrl":"","customUrlParams":"","layoutVersion":172163682823,"layoutName":"Cvent Styles - Messages - Passkey","isClientRedayToRedirect":true,"isBeforeOrIdle":false}`)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Queueit-Queueitem-V1", "ro48nDpSpeCyZzgHDkASfqLUyY_AIiFlawVTLnDJtynsuATQ4iEYMjOMEKqjNVLfjcLdozcWsERZPvv8ypLvzYOCjxF0aONVl5Xqz-zxCof0lu0at1gWzh7ZikBRdx-d9MrWKIFMTmIXc2JD61wtFkyuQF6ztcSwLNp6fnDCUbocVUbRuYWH6p8Dk8ud6Tdm_TEydsrmdlsL2u-p3esqya3jXBE2OpPuKJS0uwaLGUTRqOX9SB9pVTByzgKTXA7ZN3M0p3s7Vnp4q2C5791NICpGmwupIcZnU8bGX-c7TXiNaoTXfqqVXGg6hV1kQJ-FnaudWO2_eliiPcS2yW4GOYFZxaApZ96CIEIxsG9adVkfX0gIrgyuVf0vrK2JE1WUKUhRg9Dd9rUNGqjbScAFgN5FjrTInNkpJsRyvHTJax31UNf1JuqGrx27YUtmpY4b9WbJ4MhoDkDtKJgEwMJK1-ai7h9UqmRzuE8oC1GkL8iK2Nn6nd-yGBeWytQRaPsvucpCK01kRWodM90lulnX7uthHfP0oEkbUjOhC5MfsG7s3OVJAEDLt7pMLiJDfUuqWYZo21fmGUWzRaxPra0FL8DOk0ut7XBcL6oik-cFq4AvdBed-RruBF1B81A15ezu2l1286qc85Ia77sXL57iiB86YubQ6q0UA0CUbNkfM-_8EOvUJsZOuunTltzeattuZFQdrwjyqlwnILriPbhuvNQ3b9ctiIGYv-DSVzgqEK8XYr43ItXfpLVUeqMG7UcgOrIoQqEn3mSVLDZ5ZtMH9851KwrGueokgbgLA0CvBF2ECEd_zZT6qLJus8_QMsE5Of3bv691-K3Yp3ax5FMWOF5vcpwEDybaZtCyNLTiJkmP8Ns18xVeTG6QnZcRXAapbpHWXrDyBGCk0fMlvfizAx05lh2Nxdp_T4mc-Mc5dxFmVAHAj5EudhV3srU7LLQ-Hjk1bpqziDyHQbY7ByZDMds8TwttEbMNJpuOY_IEpgT-uZZgj8usCRStcnV9Y1cyD8g805eKPbKqj_fjmqg-7Mq17GVz64r4iUmQ9z-vbFIww3cDIwouGTsBWasNoZQW0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://queue.passkey.com/?c=lanyon&e=anthrocontest2028")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "Queue-it=u=b0639742-9f2a-458f-8c8d-c0e63627fa66; Queue-it-lanyon______________anthrocontest2028=Qid=c99d9731-a438-4f88-aeed-facba374235b&Cid=en-US&f=0; Queue-it-c99d9731-a438-4f88-aeed-facba374235b=uifh=O7Y5LI5D1Op3tSyblb3-q6LUyY_AIiFlawVTLnDJtym0jwdEClKq7Gsr5-oFJ1eU0&WasRedirected=false&i=638428832266472041")
	resp, err := e.Task.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var QueueStatus queueit.InQueueResp
	var RedirectURL queueit.LinkFound

	err = json.Unmarshal(body, &QueueStatus)
	if err != nil {
		fmt.Println("Error unmarshalling Queue Status")
	}

	err = json.Unmarshal(body, &RedirectURL)
	if err != nil {
		fmt.Println("Error unmarshalling RedirectURL")
		fmt.Println(body)
	}

	if RedirectURL.RedirectURL == "" {
		fmt.Println("Current Users AHead of You: ", QueueStatus.Ticket.UsersInLineAheadOfYou)
	} else {
		fmt.Println("Redirect URL Found")
		fmt.Println(RedirectURL.RedirectURL)
		e.Task.Stage = Stop
		e.Flow = "Booking"
	}

}
