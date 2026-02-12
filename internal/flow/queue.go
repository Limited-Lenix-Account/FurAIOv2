package flow

import (
	"FurAIOIgnited/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const (
	Challenge = "Challenge"
	Solve     = "Solve"
)

type QueueTask struct {
	BaseURL      string
	SoftBlockUrl *url.URL

	Client *http.Client
	UQueue string

	E string
	C string

	Key         string
	ImageBase64 string

	SessionID        string
	ChallengeDetails string
	Durationg        string
	Solution         string
	Tries            int
	Solved           bool
	Checksum         string

	UserAgent string
	SourceIP  string
	Timestamp time.Time

	Status string
}

func (q *QueueTask) StartQueue() {
	q.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	q.BaseURL = "https://queue.passkey.com/?c=lanyon&e=furryweekend2024"

	q.Tries = 1
	q.GetQueueID()

	// switch q.Status {
	// case Challenge:
	// 	q.GetChallenge()
	// case Solve:
	// 	q.ReadImage()
	// }

	q.GetChallenge()
	q.SolveChallenge()
	time.Sleep(3 * time.Second)
	q.GetRedirect()
	// q.GetQueueStatus()

}

func (q *QueueTask) GetQueueID() {
	fmt.Println("Getting Queue...")

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error making session!")
	}
	q.Client = &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", q.BaseURL, nil)

	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", q.UserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := q.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// fmt.Println(resp.Request.URL)

	u, _ := url.Parse("https://queue.passkey.com")
	for _, v := range q.Client.Jar.Cookies(u) {
		if v.Name == "Queue-it" {
			// fmt.Println(v.Value)
			q.UQueue = strings.Split(v.Value, "=")[1]
			q.Status = Challenge
		} else {
			fmt.Println("Error getting QueueIt U")
		}
	}

	q.SoftBlockUrl = resp.Request.URL

}

// This will Get a base64 encoded image and set it to the session.
func (q *QueueTask) GetChallenge() {
	fmt.Println("Getting Challenge")

	req, err := http.NewRequest("POST", "https://queue.passkey.com/challengeapi/queueitcaptcha/challenge/en-us", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("X-Queueit-Challange-EventId", "furryweekend2024")
	req.Header.Set("X-Queueit-Challange-CustomerId", "lanyon")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("X-Queueit-Challange-Hash", "Jlm2GzG8cAwO/N4HZVypHJ8QscBdaJaRYkGkofCa+Ig=")
	req.Header.Set("X-Queueit-Challange-reason", "0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", q.SoftBlockUrl.String())
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	resp, err := q.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s\n", body)

	var Challenge QueueItChallenge

	if resp.StatusCode != 200 {
		fmt.Println("Error getting Challenge, ", resp.StatusCode)
		q.Client.Do(req)
	} else {

		err = json.Unmarshal(body, &Challenge)

		if err != nil {
			fmt.Println("Error Unmarshalling Challenge Response")
		}

		q.Key = Challenge.Key
		q.ImageBase64 = Challenge.ImageBase64
		q.SessionID = Challenge.SessionID
		q.ChallengeDetails = Challenge.ChallengeDetails
		q.Status = Solve

		return
	}

}

// Read base64 image, and loop until a solution is found and verified.
func (q *QueueTask) SolveChallenge() {
	fmt.Println("Solving Challenge...")
	pattern := regexp.MustCompile("[a-zA-Z0-9]+")

	for !q.Solved {
		fmt.Println("In Loop")
		solveStr, err := ReadImage(q.ImageBase64)
		fmt.Println("Text found in image", solveStr, ".")
		matches := pattern.FindAllString(solveStr, -1)
		fmt.Println(matches)

		if err != nil {
			fmt.Println("Error Reading Base64 Image")
		}

		if len(matches) < 1 || len(matches[0]) < 4 {
			fmt.Println("Retrying Challenge")

			if len(matches) != 0 {
				SaveUnsuccessfulCaptcha(q.ImageBase64, matches[0])
			}

			q.Solution = "uerqerg"
			q.Tries++
			q.SubmitChallenge()
			q.GetChallenge()

		} else {
			fmt.Println("Regex Match", matches[0])
			fmt.Println("Solution Found")
			SaveSuccessfulCaptcha(q.ImageBase64, matches[0])
			q.Solution = matches[0]
			q.SubmitChallenge()

			if q.Solved {
				break
			}

			q.GetChallenge()

		}

	}

}

func (q *QueueTask) SubmitChallenge() {

	s := Stats{
		UserAgent:      q.UserAgent,
		Screen:         "1440 x 900",
		Browser:        "Chrome",
		BrowserVersion: "121.0.0.0",
		IsMobile:       false,
		Os:             "Mac OS X",
		OsVersion:      "10_15_7",
		CookiesEnabled: true,
		Tries:          q.Tries,
		Duration:       23246,
	}

	v := QueueItVerify{
		ChallengeType:    "botdetect",
		SessionID:        q.SessionID,
		ChallengeDetails: q.ChallengeDetails,
		Solution:         q.Solution,
		Stats:            s,
		CustomerID:       "lanyon",
		EventID:          "furryweekend2024",
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
	req.Header.Set("User-Agent", q.UserAgent)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", q.SoftBlockUrl.String())
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	resp, err := q.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	var QueueSuccess InQueue

	err = json.Unmarshal(body, &QueueSuccess)
	if err != nil {
		fmt.Println("Error Unmarshalling Successful Queueit Json")
	}

	q.Checksum = QueueSuccess.SessionInfo.Checksum
	q.SourceIP = QueueSuccess.SessionInfo.SourceIP
	q.Timestamp = QueueSuccess.SessionInfo.Timestamp
	q.Solved = QueueSuccess.IsVerified

}

func (q *QueueTask) GetRedirect() {

	j := QueryJson{
		SessionID:     q.SessionID,
		Timestamp:     q.Timestamp, // not the timestamp from the verification response, timestamp from when the first request was made
		Checksum:      q.Checksum,
		SourceIP:      q.SourceIP,
		ChallengeType: "botdetect",
		Version:       6,
		CustomerID:    "lanyon",
		WaitingRoomID: "furryweekend2024",
	}

	queryJson, err := json.Marshal(j)
	if err != nil {
		fmt.Println("Error marshalling Query JSON")
	}

	_url := "https://queue.passkey.com/?c=lanyon&e=furryweekend2024&cid=en-US&scv=" + string(queryJson)

	fmt.Println(_url)

	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="121", "Not A(Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("DNT", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://queue.passkey.com/softblock/?c=lanyon&e=furryweekend2024&cid=en-US&rticr=0")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// req.Header.Set("Cookie", "Queue-it=u=d069000d-5451-4b50-ad9d-e4e683c4c755; queueitsoftblock_ecc1a7d5-fad9-4856-a2ed-c1dc77734784=J/wYmHcuWUNL9zHvT3lyKbK6pr5PvY9sCkdz1qOj2IE=; queueitsoftblock_ec6c5b57-a78f-48bd-8110-3e7c753bcc04=u8ySdNw6Z/Rw0/6jLvEYVFRmt8ZwWghmvZaQMU+qNvc=; Queue-it-lanyon______________furryweekend2024=Cid=en-US&f=0")
	resp, err := q.Client.Do(req)
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

	event, owner, _ := util.GetOwnerHotel(resp.Request.URL.String())

	fmt.Printf("Event: %s, Owner: %s\n", event, owner)

}

func (q *QueueTask) GetQueueStatus() {

	p := PostStatus{
		TargetURL:               "",
		CustomURLParams:         "",
		LayoutVersion:           171445457923,
		LayoutName:              "Cvent Styles - Messages - Passkey",
		IsClientRedayToRedirect: true,
		IsBeforeOrIdle:          false,
	}

	postJson, _ := json.Marshal(p)

	// url := fmt.Sprintf("https://queue.passkey.com/spa-api/queue/lanyon/furryweekend2024/50e5ed2b-9f1c-475f-a4e5-afd15ae9117d/status?cid=en-US&l=Cvent%20Styles%20-%20Messages%20-%20Passkey&seid=229caaae-0333-68e3-c908-edc2e3bd0be0&sets=1706837548528")

	var data = strings.NewReader(string(postJson))
	req, err := http.NewRequest("POST", "https://queue.passkey.com/spa-api/queue/lanyon/furryweekend2024/50e5ed2b-9f1c-475f-a4e5-afd15ae9117d/status?cid=en-US&l=Cvent%20Styles%20-%20Messages%20-%20Passkey&seid=229caaae-0333-68e3-c908-edc2e3bd0be0&sets=1706837548528", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "queue.passkey.com")
	req.Header.Set("sec-ch-ua", `"Not A(Brand";v="99", "Google Chrome";v="121", "Chromium";v="121"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Queueit-Queueitem-V1", "GN24Tg9mYME4sXO1N6HjaqLUyY_AIiFlawVTLnDJtynsuATQ4iEYMjOMEKqjNVLfjcLdozcWsERZPvv8ypLvzYOCjxF0aONVl5Xqz-zxCof0lu0at1gWzh7ZikBRdx-d9MrWKIFMTmIXc2JD61wtFkyuQF6ztcSwLNp6fnDCUbocVUbRuYWH6p8Dk8ud6Tdm_TEydsrmdlsL2u-p3esqya3jXBE2OpPuKJS0uwaLGUTRqOX9SB9pVTByzgKTXA7ZN3M0p3s7Vnp4q2C5791NICpGmwupIcZnU8bGX-c7TXiNaoTXfqqVXGg6hV1kQJ-FnaudWO2_eliiPcS2yW4GOYFZxaApZ96CIEIxsG9adVkfX0gIrgyuVf0vrK2JE1WUKUhRg9Dd9rUNGqjbScAFgN5FjrTInNkpJsRyvHTJax31UNf1JuqGrx27YUtmpY4b9WbJ4MhoDkDtKJgEwMJK1-ai7h9UqmRzuE8oC1GkL8iK2Nn6nd-yGBeWytQRaPsvucpCK01kRWodM90lulnX7uthHfP0oEkbUjOhC5MfsG7s3OVJAEDLt7pMLiJDfUuqWYZo21fmGUWzRaxPra0FL8DOk0ut7XBcL6oik-cFq4AvdBed-RruBF1B81A15ezu2l1286qc85Ia77sXL57iiB86YubQ6q0UA0CUbNkfM-_8EOvUJsZOuunTltzeattuZFQdrwjyqlwnILriPbhuvNQ3b9ctiIGYv-DSVzgqEK8XYr43ItXfpLVUeqMG7UcgOrIoQqEn3mSVLDZ5ZtMH9851KwrGueokgbgLA0CvBF2ECEd_zZT6qLJus8_QMsE5Z4yHIGFVhLbFm7IAfoAafdKeb-q9APKBv-_UDhLixM8CBIIIT2Kg0XDcfl8tZAuFK2qOo_XRt0s2IossecXSlqvs8gsK876MwZvTp_sgJatQYdhPBLF5pAmQrIeWlNeiHcNWlOA3_8o4kr3-BTFpZMwkSdfxpZhH27Y4Zoq-YksoU4IOpTorMJ26P-lb5TdciNuKpGwqf-d-MtDsejdovxn2SkclWsa59B2Dtbg66NYaKViO8y569q1uU5ON23eh0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://queue.passkey.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://queue.passkey.com/?c=lanyon&e=furryweekend2024&cid=en-US&scv=%7B%22sessionId%22%3A%22dc1133b8-17dd-4e89-b288-c4f04ca5c41b%22%2C%22timestamp%22%3A%222024-02-02T01%3A32%3A28.280377Z%22%2C%22checksum%22%3A%22RsR%2Bolk3cNR%2FZf%2FoGPNv6A4wVBfkBknwgrLxLbbIong%3D%22%2C%22sourceIp%22%3A%22154.6.90.137%22%2C%22challengeType%22%3A%22botdetect%22%2C%22version%22%3A6%2C%22customerId%22%3A%22lanyon%22%2C%22waitingRoomId%22%3A%22furryweekend2024%22%7D")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "Queue-it=u=76846597-2273-4848-b5f1-b30e83683c63; queueitsoftblock_50e5ed2b-9f1c-475f-a4e5-afd15ae9117d=5QglZhO7cpxKjhak/AOuE81n+eoM5I5YLffQA+7OxEU=; Queue-it-lanyon______________furryweekend2024=Qid=50e5ed2b-9f1c-475f-a4e5-afd15ae9117d&Cid=en-US&f=0; Queue-it-50e5ed2b-9f1c-475f-a4e5-afd15ae9117d=uifh=O7Y5LI5D1Op3tSyblb3-q6LUyY_AIiFlawVTLnDJtym0jwdEClKq7Gsr5-oFJ1eU0&WasRedirected=false&i=638424343487669329")
	resp, err := q.Client.Do(req)
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

func ReadImage(base64 string) (string, error) {

	scriptPath := "/Users/elijah/Documents/coding/FurAIOIgnited/scripts/solve.sh"
	arg := base64
	cmd := exec.Command(scriptPath, arg)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(output), nil

}
