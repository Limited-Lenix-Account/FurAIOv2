package pa

import (
	task "FurAIOIgnited/cmd/taskengine"
	"FurAIOIgnited/internal/browser"
	"FurAIOIgnited/internal/discord"

	// "FurAIOIgnited/internal/server"
	"FurAIOIgnited/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/playwright-community/playwright-go"
)

type PaTask struct {
	task.Task
	Serverside  bool
	homeURL     *url.URL
	Cookies     []*http.Cookie
	Method      string
	Fulfillment string

	TaskStartTime     time.Time
	TaskEndTime       time.Time
	CheckoutStartTime time.Time
	CheckoutEndTime   time.Time

	ShippingID string
	LocationID string // For Checkout

	LocationIDGroup []string // For Monitor
	MonitorSKUs     []string

	Keywords      string
	ProductID     string
	ProductName   string
	ProductPrice  float64
	ProductImage  string
	ProductVol    string
	QuantityStr   string
	QuantityInt   int
	ShippingGroup string
	Username      string
	Password      string

	UseAccount    bool
	AccessToken   string
	EncryptedCard string
	OrderTotal    float64
	OrderUUID     string

	PreloadSku string // 00000XXXX
	PreloadId  string // ciXXXXXXXXX
	Precarted  bool
}

var messageChan = make(chan []byte)

func (e *PaTask) Ignite(bPool *rod.Browser) {

	e.homeURL, _ = url.Parse("https://www.finewineandgoodspirits.com/")
	e.UpdateStatus(Start)
	e.PrintStatus("Starting....")

	if e.UseAccount {
		e.Login()
		e.CheckProfile()
	}

	switch e.Mode {
	case "Normal":

		if e.Serverside {
			e.UpdateStatus(MakeSession)

		} else {
			e.ProductName = e.Profile.LiquorKeywords
			e.UpdateStatus(Monitor)

		}

	case "Preload":
		e.UpdateStatus(AddPreloadItem)
		e.Precarted = false
	}

	e.TaskStartTime = time.Now()
	for e.Task.Stage != "Stop" {

		switch e.Mode {
		case "Normal":

			switch e.Task.Stage {
			case Listen:
				// e.ListenToWebsocket()

			case Monitor:
				e.StartMonitor()

			case CheckStock:
				e.CheckB2CStock()

			case MakeSession:
				e.PrintStatus(CreateBrowser)
				e.Cookies = browser.StartRod(bPool)

				e.PrintStatus(SetCookies)
				e.Client.Jar.SetCookies(e.homeURL, e.Cookies)
				e.UpdateStatus(AddToCart)

				// os.Exit(1)

				e.Stage = AddToCart

			case AddToCart:
				e.AddToCart()

			case CheckItems:
				e.CheckItems()

			case SubmitShipping:
				e.SubmitShipping()

			case GetShipRates:
				e.GetShippingRates()

			case SubmitShippingRate:
				e.SubmitShippingRate()

			case EncryptCard:
				e.EncryptCard()

			case SubmitOrder:
				e.SubmitOrder()

			}
		case "Preload": // TODO - add new browser logic to open first to pre-cart stuff (not default anymore)

			switch e.Task.Stage {

			case AddPreloadItem:
				e.ProductID = "000081920"
				e.AddToCart()

			case SubmitShipping:
				e.SubmitShipping()

			case CheckItems:
				e.CheckItems()

			case CheckOrder:
				e.CheckShippingGroup()

			case RemoveItem:
				e.DeleteItem()

			case Monitor:
				e.StartMonitor()

			case AddToCart:
				e.AddToCart()

			case SubmitShippingRate:
				e.SubmitShippingRate()

			case EncryptCard:
				e.EncryptCard()

			case SubmitOrder:
				e.SubmitOrder()
			}

		}

	}
	// e.Login()
}

// func (e *PaTask) ListenToWebsocket() {

// 	fmt.Println("Listening to Server...")
// 	var skuData server.SkuData
// 	go server.ListenToSocket(e.Websocket, messageChan)

// 	for {
// 		message := <-messageChan

// 		err := json.Unmarshal(message, &skuData)
// 		if err != nil {
// 			log.Fatal("error unmarshalling socket json")
// 		}
// 		if string(message) == "null" {
// 			continue
// 		}

// 		e.ProductID = skuData.Sku
// 		fmt.Println(e.ProductID, skuData.ProtectedStock)

// 		if skuData.Available { // should be always true, i just don't want it yelling at me lmfao
// 			e.UpdateStatus(MakeSession)
// 			break
// 		}
// 	}

// 	p := GetSkuDetails(e.ProductID)

// 	e.ProductID = p.ProductID
// 	e.ProductName = p.Name
// 	e.ProductImage = p.Image
// 	e.ProductVol = p.Volume

// 	e.MonitorSKUs = append(e.MonitorSKUs, e.ProductID)
// 	e.UpdateStatus(CheckStock)
// 	e.CheckoutStartTime = time.Now()

// 	// interrupt := make(chan os.Signal, 1)
// 	// signal.Notify(interrupt, os.Interrupt)
// 	// <-interrupt
// 	// log.Println("Interrupt received, shutting down...")

// }

func (e *PaTask) AddToCart() {

	data := strings.NewReader(e.EncodeCartJson())
	// var data = strings.NewReader(`{"items":[{"quantity":1,"productId":"000004405","catRefId":"000004405","locationId":null,"shippingGroupId":""sg99514577""}]}`)
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/items/add?exclude=embedded.order.shippingGroup%2Cembedded.order.shippingMethod%2Cembedded.order.shippingAddress", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/principato-pinot-grigio-chardonnay/product/000007646")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/principato-pinot-grigio-chardonnay/product/000007646")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")

	for e.Task.Stage != CheckItems {

		resp, err := e.Client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			fmt.Println(string(body))
		}
		// fmt.Printf("%s\n", body)

		if resp.StatusCode == 200 {

			var atcResp AddToCartResp

			err = json.Unmarshal(body, &atcResp)
			if err != nil {
				fmt.Println("Error umarhslaling atc response")
			}
			// fmt.Println("ATC Success")
			e.ShippingGroup = atcResp.Embedded.Order.ShippingGroups[0].ShippingGroupID
			e.ProductPrice = atcResp.Items[0].UnitPrice
			// e.Task.Stage = CheckItems
			e.UpdateStatus(CheckItems)

		} else {

			var errr PaError
			err = json.Unmarshal(body, &errr)
			if err != nil {
				fmt.Println("Error umarhslaling atc response")
			}

			switch errr.Errors[0].ErrorCode {
			case "28104":
				// fmt.Println(e.ProductName, "Out of Stock!, retrying")
				e.UpdateStatus(OutOfStock)
			case "28102":
				e.UpdateStatus(ProductNotLoaded)
			}
			time.Sleep(3 * time.Second)
			// fmt.Println(errr)
			e.UpdateStatus(Retrying)
		}

	}

}

func (e *PaTask) StartBrowser() {
	fmt.Println("launching playwright")
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	fmt.Println("launching chrome")

	opts := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	}

	browser, err := pw.Chromium.Launch(opts)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	fmt.Println("new tab")
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	fmt.Println("launching going to site")
	if _, err = page.Goto("https://www.finewineandgoodspirits.com/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	time.Sleep(3 * time.Second)

	page.Locator(`//*[@id="root"]/header/section/div[3]/div[3]/div/div/div/div/div[3]/button`).Click()

	time.Sleep(60 * time.Second)

	fmt.Println(page.URL())
	cookies, err := page.Context().Cookies()
	if err != nil {
		log.Fatalf("could not get cookies: %v", err)
	}
	for _, v := range cookies {
		fmt.Println(v.Name, v.Value)
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

// Protected by Akamai
func (e *PaTask) CheckStock() {

	// for _, v := range e.Client.Jar.Cookies(e.homeURL) {
	// 	fmt.Println(v.Value)
	// }

	req, err := http.NewRequest("GET", "https://www.finewineandgoodspirits.com/ccstorex/v1/stockStatus?actualStockStatus=true&expandStockDetails=true&products=000006516%3A000006516%2C000079696%3A000079696%2C000003400%3A000003400%2C000007634%3A000007634%2C000079660%3A000079660%2C000007652%3A000007652%2C000006278%3A000006278%2C000003206%3A000003206%2C000007856%3A000007856%2C000006728%3A000006728%2C000007725%3A000007725%2C000001128%3A000001128&locationIds=null%2C211", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-type", "application/json")
	// req.Header.Set("cookie", "osfliveuiroute=a45ccbc8d96c6a0bad83d7c08f4fb64e|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup; ccstoreroute=5100bfefae9f09498c21e9d7b291a95d|cc2d3787e8999f4a634b9d592c9e2268; AGEVERIFY=Over21; ak_bmsc=4C5B8016DE6959FF8904FDC8815275F0~000000000000000000000000000000~YAAQIkjIF5sLy6WOAQAAu73ZyBcmtOzAyQTLdDqLeoPnStMASdPa6+wCqjlVtTTOANFIwVWz5ifui5/tmESo/SpoPIUbKBoRM03zmCiOmaXVa5vpxTiULiX9wNPuYa+pfEf29i8H8CWCwbg6eovcL27/WxHl6/Ev5qh1oZCU3LDohsyhv4ZGsK9Tbu9wE/9LK4wcnBZTQYJ8MfdlcOlmqa9YPjKYT1UzzKo+T4bsdTNhcoPQkXRVvKIM2h9P1Kof8tiqo+ae77LABuDr7gLIzRBfMJSzEPahKUMg8jnotW2Cdniw2RXsu9Wvm/3esE/5iI+zI1Fjyf7YJtEPog0NDqE4r8XNPXFgcX1jgxXxeFllVMuJpHlZXI9IJf/sLuK6vsB7eGcYTrSeHxzHs/CL2yVCABuy+mpfKk8=; bm_sz=D8B5F36C7E45DAA771E163A5D3F63070~YAAQIkjIF50Ly6WOAQAAvL3ZyBdCnZ2nAZcFwwxHCPL8eWrGcL8w7G9xxzzMKuOanPzbThZZ18N8tEDpDpjvoifb8la3hZHxJmur0eg/QuNQG+Hli8SkpfsFkXYa86cOE9dIGvDXGK0PK7p6xNAlEr7uLovEPMsVSu45CdslykNCn7ykaGo2ArgUlGQ5/Hx5xZfPCaIYKfEOwqXlUWhCAf7sC/ReFjDy6R7FpHIZ3POT09Qr0HlrVO+mxAE6JkIlcggE+8kVnCXSUGONE6yiy6ZiOj9WAoYGzXwfdJ+oLKOOnKsn5I4b2647RrqLp4dMVk9cRY59aVib/A89p7yX19gOWW8Qy1Km0qw5RJY8HF/kQVZGavGXkzgJRztn1iz2tC/Pz6IMrLpAk0tsid+cIb1wI804yYE1wOaUPWbx~3290929~3748152; __ZEHIC5581=1712766696; JSESSIONID=1GXI2cSna9Mv9EqzBvOFGZxOwzoJOdt3412heJ846iA5QF-66JHs!-2050283380; __z_a=1379342855158513056815851; _abck=6C9C8AA579E16F4E466EF8E4C502D426~0~YAAQIkjIF0QNy6WOAQAAn+/ZyAuCeJJuKdbiA71Jl/PgZyoc/GAXPCvGAYpmsL2G3QCm87IaKx7i6SpriSnb1sxwaJlZ+jKJJdrEE+tDTQtgvGgI4Gk/Q3owrKr4TjLlumFkTCzeL6wKPWmfzDr6P/w8ra2ldc7+reJE+iC+ICDYMEy2IgMZFm0jirKyIeAXefbbfmMsHhWErxTRSEvNaQ0Ra9UCYFJvuS85+3yrn1xRrSZxN+OLv4dAa8QYv2YSEOcaHFHPvBeaqm4ONQWEGiKTYgRhjTQg7fjAwPRsX3bxwaDdloyiCFJbFaOjXFVVTeSB+ZBlV4ILjhF7Y7MSRse+Cs9u+SS/7B5l2XoivwIOUM1cUUycALDm/nyXf6Zh246PF3i9tyfTsxrdac5LQjlEW+AiWk6FlZ8r9o0N9ltpDjqrzn16rg==~-1~-1~-1; bm_sv=23F91F80DFCB6BB286C8EE973D4A9815~YAAQIkjIF54Ny6WOAQAA9ADayBeB9pPVywpJ80sN0K4s7v7bn/kqbVE0Em5d6j42O2+qywQH57jnTPUmhOfdEl0fKQsJ2JD+rkXNHHHNLeurZjwEdcf806YRfcuwIT4EAf+OI91JZmN2Zk7y2wD8iWxTTHyHN9jWGX2ELJijXACFyH9eYnVwOjmmVG1TiORFdr4wFwTIfZfueyoiniocMwufXrNTqotjNI9H649z2HGok5as1pDlDZFmTrV3F3Y9Y20x2VTK3A55y2TENYhqyzk=~1")
	req.Header.Set("dnt", "1")
	req.Header.Set("referer", "https://www.finewineandgoodspirits.com/red/002?type=bopis&store=211")
	req.Header.Set("sec-ch-ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.3")
	req.Header.Set("x-cc-frontend-forwarded-url", "www.finewineandgoodspirits.com/red/002?type=bopis&store=211")
	req.Header.Set("x-cc-meteringmode", "CC-NonMetered")
	req.Header.Set("x-ccasset-language", "en")
	req.Header.Set("x-ccpricelistgroup", "defaultPriceGroup")
	req.Header.Set("x-ccprofiletype", "storefrontUI")
	req.Header.Set("x-ccsite", "siteUS")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)

	fmt.Println(resp.StatusCode)
}

// uses b2c get-inventory which is not akamai protected
func (e *PaTask) CheckB2CStock() {

	//shipping:
	// method - b2cShip
	// location - null
	// var data = strings.NewReader(`{
	// 	"method": "pickup",
	// 	"location": "9101",
	// 	"items": [
	// 	  "000007646",
	// 	  "000004405",
	// 	  "000004497"
	// 	]
	//   }`)

	c := &http.Client{}

	s := e.EncodeStockJson()

	data := strings.NewReader(s)
	//
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstorex/custom/v1/b2b/get-inventory", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("sec-ch-ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/principato-pinot-grigio-chardonnay/product/000007646")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/principato-pinot-grigio-chardonnay/product/000007646")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/json")

	var result map[string]interface{}
	for e.Task.Stage == (CheckStock) {
		var resp *http.Response
		for {
			resp, err = c.Do(req)
			if err != nil {
				log.Println("Error Sending Request: ", err)

				time.Sleep(5 * time.Second)
				continue
			}
			break

		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		if err != nil {
			log.Fatal("Error Closing the Response Body: ", err)
		}

		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.Fatal("Error unmashalling b2b stock response", err)
		}

		if len(result["errors"].([]interface{})) == 0 {
			// If there are no errors

			if result[e.ProductID].(string) != "0" {
				// If there is more than 0 stock
				e.UpdateStatus(MakeSession)
				return
			} else {
				e.PrintStatus("Out of Stock")
			}
		} else {
			e.PrintStatus(result["errors"].([]interface{})[0].(string))
		}
		e.PrintStatus("Monitoring...")
		time.Sleep(3500 * time.Millisecond)

	}

}

func (e *PaTask) SearchCategory(category int) {
	client := &http.Client{}
	url := fmt.Sprintf("https://www.finewineandgoodspirits.com/ccstore/v1/products?categoryId=%d", category)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prods CatalogResp

	err = json.Unmarshal(body, &prods)
	if err != nil {
		log.Fatal("failed to unmarshal catalog", err)
	}
	posKws, negKws := util.GetKws(e.Keywords)
	var validItem []Item
	var titleList []string
	fmt.Println("Positive KWs:", posKws)
	fmt.Println("Negative KWs:", negKws)
	// names := prods.Items

	for i, name := range prods.Items {
		splitTitle := strings.Split(name.DisplayName, " ")
		for _, title := range splitTitle {
			titleList = append(titleList, strings.ToLower(title))
		}

		if util.ListInList(titleList, posKws) && !util.ListInList(titleList, negKws) {
			validItem = append(validItem, prods.Items[i])
		}
	}

	fmt.Println("Found Names:")
	for _, v := range validItem {
		fmt.Println(v.DisplayName)
	}

}

func (e *PaTask) GetKeywords(prods CatalogResp) {
	//Baron Pichon Longueville XO Extra Armagnac
	posKws, negKws := util.GetKws(e.Keywords)
	var validTitles []Item
	fmt.Println("Positive KWs:", posKws)
	fmt.Println("Negative KWs:", negKws)

	titles := prods.Items
	for _, item := range titles {
		var titleList []string

		splitTitle := strings.Split(item.DisplayName, " ")
		for _, title := range splitTitle {
			titleList = append(titleList, strings.ToLower(title))
		}

		if util.ListInList(titleList, posKws) && !util.ListInList(titleList, negKws) {
			validTitles = append(validTitles, item)
		}
	}

	for _, v := range validTitles {
		fmt.Println(v.DisplayName)
	}

}

// I think shipping rates are *mostly* constant, May be unnecessary if so
func (e *PaTask) GetShippingRates() {
	req, err := http.NewRequest("GET", "https://www.finewineandgoodspirits.com/ccstore/v1/shippingMethods?shippingGroupIds="+e.ShippingGroup, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	// req.Header.Set("X-CCVisitorId", "135Coh4bBux5UQxpF3cBHhQIeACUm4X9ap6HfB188jHzt-MC120")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-71dec6b4:18ee8a31b94:-3343-4094299650")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "_fbp=fb.1.1713301275438.1881706559; bm_sv=F03A05757CCA26BD47942CF7F4ADF0B1~YAAQJo0hFyyoI82OAQAAZ9i36BdZnpR+UljlyMjqIT6bYWZYgwcPWoceEDHR93KjfDl96sPVjQz0fScNmL45nJxuIUnpYl8pdn1kCa1VUCRjWStUob/MPh0WXqy/xgLcLbh3CKqxeV6amphjsiwIlRWgUxOyH3rkc8TtBnshFyanPMak/axkZcYz7gob2lddEz/p1j6p2yifd7gWyBSZjRa4lAlEb55CbU5AG/yhIOpZ4v8tAtuTgpDbImFH0TQlrlmAZPAkGAEfft+ysGeSIxQ=~1; ltkpopup-suppression-8e9eb133-da66-408a-880e-6f4844cae736=1; _vuid=2a3cac1c-d9d8-4aad-8994-98b45ac92b23; _ga_Y3ZSL11TM0=GS1.1.1713301275.1.1.1713301292.43.0.0; _abck=4AE21EFF104E00BEE9E2BF46FF36EB8B~0~YAAQIkjIF8zSHcuOAQAAWAK36Av44gpfv61m1SuAxRO3YMkXvEZOEM7+aDETlmrCJk9RGDWuMzY532y3ZtohGGXXh33Lk8DdgtKnSoRl7BRVvX/m2wdEVE3HH3zRGS/pszABLIckBk8X2t+8MoY9sR++/Wr9aqYatqr8TZGMV5ERTMW8S9zyHwA4SlJhsN6YOBpF677DBE4ObN+jCTr3IS9f0OOn3gL88hQ8BoUk5VQyw00e5/PXetTbfGsCVZSOWkQdZPBcuPTjaGADbVpO/ypwSDcj/56lM3diENocoP9atV9sFtj/fJKIxK0qXU93OMHHuVS/HlHQ1sl/JoT1WaLfIfU0HKsQBLad3Y7b59dmcrmxB9hTI0x/Xb+6fhh8Qf+7m6SBgy4G+s7NMCas5xh0ndVfA6BkOJ7t02S22AqjE1WptICzmkk=~-1~-1~-1; storePriceListGroupId=defaultPriceGroup; xda8383328c1PRD_siteUS=135Coh4bBux5UQxpF3cBHhQIeACUm4X9ap6HfB188jHzt-MC120%3A1713301278178%3A1713301278178%3A1713301278178%3A1%3A1; xva8383328c1PRD_siteUS=-71dec6b4%3A18ee8a31b94%3A-3343-4094299650; AGEVERIFY=Over21; __z_a=773377662341602386634160; ltkSubscriber-Footer=eyJsdGtDaGFubmVsIjoiZW1haWwiLCJsdGtUcmlnZ2VyIjoibG9hZCIsImx0a0VtYWlsIjoiIn0%3D; JSESSIONID=_Onotsx0c5wTmmjjJy1YqUrdnrlKK3HhP0zkq7mStzGgCrkCbOVm!-1728041536; ccstoreroute=7203db4be79cc4a2c7943f648c3401c3|cc2d3787e8999f4a634b9d592c9e2268; ak_bmsc=097261B5BE09A200A6134F9A5EC14349~000000000000000000000000000000~YAAQIkjIF3nPHcuOAQAAfsi26BdQA/8sByF+w19yp39sHKvzEg4Z0/DxnoXl0f58/g3K0C61enla1Uvyxr+dSIIK3oNV0WUnxGzpuQOvBeB62qPtF6Sk+wJbyi3hw0CFCB5qHD4wT+XCJl71Jppi9DEyAvFvzIV3MU/O3DzDBEm/Ndex5DS+8uJJavTchSWUUVzka4Q9VXsZEtX3T6IQahHShC4RzkT7y7TAY0Ie48aTKVvHVJPQGhwLN90MBejPiOO02M/B0Jx5mfR4ErsggUBsF5FW+9a/7dGUJfXtSAiGOC0zTGvMWtJkQ1Kg0fo4Hi7e11OfU+wghq5oZcO5cJjOw8YIxWiTlMZ34uifeFm0ZftH+zPVuXlCTXyFRn/po0he5r+FSxUQl8e8h0Jhys3VSKJAPPBQ9hsu9E5ZzNHxkB2TWwWSD5u0EJRBfZMv2TKeLkqWbw6auDy01g==; BVBRANDID=7b792510-673d-48cb-813b-b1eae7c47bc5; BVBRANDSID=b038e447-cd72-460d-b988-ea0d244ef2ad; GSIDyk8q9gLSz40L=2b015953-4365-4630-8e5c-e3b676446776; STSID148379=9e9afa3f-00ff-4502-97df-cceaa9b6c520; _ga=GA1.1.802531164.1713301275; _gcl_au=1.1.633540061.1713301275; _pin_unauth=dWlkPVpEVmhNVGN3TkRRdFpEaGxOQzAwWmpBd0xXSmtabU10TmpjNE0yWTBNRGxrTldObA; ltkpopup-session-depth=1-2; __adroll_fpc=7f50c153ee0ab4c48ffb612370ad116f-1713301275806; __ar_v4=%7CYEV7UUOV4NHJNEIGTN5CUO%3A20240416%3A1%7CH422MU2HOBC4HLZGUQMCAG%3A20240416%3A1; __ZEHIC6300=1713298388; __pdst=acd2de3a826e490fb1befa694ce5ea39; bm_sz=AA6C7DD0B4601F39F7C56299C0ACE0E3~YAAQIkjIF9nOHcuOAQAAjMC26BcWHNKRPk/Z8HLLDpG51hqZ4yCaimuESriK5Ycnq8m29Q7syqQZTT7Ap9SGo9CGQQDtUlWOpOa0aH9jppNiN3CptsFsehntBG0OmZJjdBuI3/x/Z6rQqPw+bRnWzP+MjJ0Ao+xg7mNgm1MKKixXGj/XKcgAjmCxLNH5JlH48Xn/wm8Xmrwja/UJLbvykhDOuu/+esASVrpbm4PENZwwf5tt3PnrfzGU3JLHRIxwQQmVYHsZKdj3sQ+0qjh49+2nIkqpD1pUliSckW2uVHS6if7RQQR5nIxMldKvgaKaihfFcLtrHzIi24lBiQL5z/1DZ8IAGzxuFVOe7XjSsgF9lpczuiM50uv/1V0sZjR0owg2FJXlGsz1tG+sFu5z7/nLBGTv~3683890~4404035; osfliveuiroute=7aa2295821121d5d9ba1a190c8adfaf6|a271690a033ecd4ba18a69565eca7c19")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	e.Task.Stage = Stop
}

func (e *PaTask) CreateShippingCart() {
	var data = strings.NewReader(`{"items":[{"type":"hardgoodShippingGroup"}]}`)
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/shippingGroups/add?exclude=embedded.order.shippingGroup%2Cembedded.order.shippingMethod%2Cembedded.order.shippingAddress", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	// req.Header.Set("X-CCVisitorId", "10C7QZZku5nlm-NsqcjIRupXsbiPr4LkQgg87waqVvXNTBQ9F92")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/cognac/124?type=ship&store=5105")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCSite", "siteUS")
	// req.Header.Set("X-CCVisitId", "240b89f4:18ede8e1e7c:2172-4094298633")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	// req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/cognac/124?type=ship&store=5105")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "_gcl_au=1.1.2111076155.1712769182; GSIDyk8q9gLSz40L=3a46ef7f-0269-4992-b81c-aa23a55257fd; STSID148379=e6f84471-0f49-4ba7-bf3f-0d50fcd8f76c; BVBRANDID=a82f17bc-93fa-4d18-94af-62fda45fc91e; _ga=GA1.1.2050355724.1712769182; __pdst=1b26f91811f54e848b091ce7573f7f59; __adroll_fpc=d8f13915972a0872f73533a05cdab84c-1712769183244; _pin_unauth=dWlkPU5HTTNZakV3TUdJdE16UmtaaTAwWmpaaUxUZ3paR1V0WW1Fek9UZGhaREExTTJVMA; _fbp=fb.1.1712769211703.185896143; ltkpopup-suppression-8e9eb133-da66-408a-880e-6f4844cae736=1; ccstoreroute=e377ede55d92f51d0b11dd3801f92e6b|cc2d3787e8999f4a634b9d592c9e2268; JSESSIONID=nRTetWIisDyPjtPSoA7snhdeC1SqNHo-oxeMwMAQ8CBXZIXlLaoD!1108060648; osfliveuiroute=e0ed6d79e84c8faff28ef5c16fa91b41|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup; bm_mi=748B4AB9EA5C360F67BB46F003534DE3~YAAQpTovF2nQLsuOAQAArYC13hdC10t0vNb0SbfT9dAM2CzaFRrpb/nqK+NlreVfjt3HQi48NTGJMB3KPiDp+AYyhc+BsJq5fBa6vK3YFRHOx466urylqty7HvqojfXEjmS4rULUdpvi5/njtmTwn82Pu5JaNEY68t0+wFj1F7SF3gFVb3hN5dZflnFZJRvGWmMw/kzYWM6oULrwImSeJ+/fzWg6GR3IoZz53aM1/g++CiWu4nxwy6KWMuntQRAeD/FLycv9UDYx07U+0PNkiXFuyZzOWmTZO/1ugRQ99OvBb3qL8kOIf+tkk9odtpLugbkcfmZZcDPy2OuEOQ==~1; bm_sz=41C5542309474F327C1EBDE9874BC30C~YAAQpTovF2vQLsuOAQAArYC13hehQea6HEYa+rrCbkMuHbhjz3fhSufA1oF+TlTqwexNOgJXM8qMCEtw9rBStFDhLiTzXHECLGWHrfLrics4HlMI7AVlrlg+rHkyLrF+x2dDyOElS+yZbhbV3QfuLyQdsuCrDCEgz+DudrT+4rHdfV81wg/m2l4QHLr0x8upxb8EugOL6EOzBttVq4W+A58nP5GhTHy/Fo3i4Emc9s+eQnQ1PFPFTdTXFx68OIbPoBuLUtPiRgYzJPsZbx83JPFus8m0OXoXOQ92wmFTNnNRP3dsOuEaIoRxwkati55bG8uEnU//p414bw1z2oO55JtIaj3BBEA1ckvydwiKmLAqy4LEHFqK8TiuSxTUTRt4FK8h65FisNhdTc7+jXls/gOwnNDanv8McZt5slwC7hVXz5Q=~4604466~3354947; BVBRANDSID=5193d2f9-56c7-4b08-9918-27616f25c0a9; ltkpopup-session-depth=1-2; ltkSubscriber-Footer=eyJsdGtDaGFubmVsIjoiZW1haWwiLCJsdGtUcmlnZ2VyIjoibG9hZCIsImx0a0VtYWlsIjoiIn0%3D; __ar_v4=H422MU2HOBC4HLZGUQMCAG%3A20240410%3A34%7CYEV7UUOV4NHJNEIGTN5CUO%3A20240410%3A34; _vuid=5622f301-a4dc-4819-844a-04b2d50d9a95; ak_bmsc=25A98E71181A01CA8C926F961875F58A~000000000000000000000000000000~YAAQpTovFxHRLsuOAQAAoYS13he8EONoIDjgBbZ4C6zmxX7F+wQNjaK9VOYwG7JoAd4LF9wcqCVj9SPIXJp4/6AmX3VwygOyBIgZQGsPOGZwgwL74HuqzV2KcObY8JKjPG/y3Pfww5MmUUmcAXEgLsxxPnBCObH/nitZUFoQ2kcImTH5IkqK954aZVG7/eDLwonvKZhf0/JZOUNzOr36DWb581kux8gyX2zVCtGQMTQFnVuszjDQsft96ZsfvTCLM+j5vyFuAcnQLiUmzGFEpWtLHPy40smk7gCMneyrqgS+Z18oECpfXJcu6nwLaW/iky5UgC/pZlGJjrasy0KpxBhpxaHZUpSlUAekRx+5KM6UZfeFYExJKUuLcoFnTiD65KPOSaEjyG7XD/9xL5QQlytHKV0AHnvgpotxTOhM6mZnmzkmm78WuVlVKRldqL0JzO/OjltxKs1Bvq0OG1jZGHs57Va034RPkt+Gyt5/LKhE8q12XwIoH2f1gSzhRw==; __z_a=3762202327352638300035263; __ZEHIC1810=1713130914; xva8383328c1PRD_siteUS=240b89f4%3A18ede8e1e7c%3A2172-4094298633; xda8383328c1PRD_siteUS=10C7QZZku5nlm-NsqcjIRupXsbiPr4LkQgg87waqVvXNTBQ9F92%3A1712769183705%3A1712981713371%3A1713133422171%3A12%3A12; AGEVERIFY=Over21; _ga_Y3ZSL11TM0=GS1.1.1713133421.9.1.1713133553.14.0.0; _abck=0CFD734DABA18D369949E71FA27D3CC8~-1~YAAQuzovF289y6uOAQAAkNK33gshfaw1XNT1rcvcRICP/gYqd7v30vBwnQSBbA0WfqsUf+aHLG/pCxLLqJryNapMEBTZrVcTs6wlB4tsZtisiWYfIHMDBa4AJEbqehlnUsy0lII7hjS9ziW9X5rKAkdMVcdZmJ112t/HZXZ2xcglKaH+6hCSMR34GSNHioW4lmBt3uSaE1elvcJxlXpLBh64kB+4e4nsOBMQA3YztFt5UUfn3eNnscoEvFoJ9z7Kebqh7P2u9fVo/Gg2IeSSx5hF7bHCe7goA2UvXbP9hvJJ04FIpmO2BVMdc4PM490SlD0drSOnjXjlygmHiRZ9+RISC8jyOEQ7s3hMLmPAcuaZO199Uk0xrG9Z6B/3aPtBKKFgJZFEViHTgFGpOv+c08/uMuB7Mzsrhpu1KDEyXgaYXRgLndBy5g==~0~-1~-1; bm_sv=5FB7F4ED9F98D6DC5E5004BCF11DD632~YAAQuzovF3A9y6uOAQAAkNK33heOVgd1kmKZ5ghUqsEpuTZ+BN8ey3b1m8aPn+RRghYcBBpMexcgg418eNHe26G8JE0mo5ZQry4U6s/WCRbSKahxNNjnsSYCQiFqffzI4BmCxUROuORPxcTIgMJuzsc6t7Zdn9Ee0bap10kFfFnm8rCS8Bl5aoCFNmn/Rt87Xbvlwm2SaemuzM0NtVa96VFsOTNd/5i9bir9vtZlrVLVX1E4YS58BFY5bdVJEa/9DwXNrmREHiRDFsa/IZk3i/0=~1")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cart CreateCartResp

	err = json.Unmarshal(body, &cart)
	if err != nil {
		log.Fatal("error unmarshaling cart creation", err)
	}
	e.ShippingGroup = cart.Items[1].ShippingGroupID

	if len(cart.Items) > 1 {
		for _, v := range cart.Items {
			if v.Type == "hardgoodShippingGroup" {
				e.ShippingGroup = v.ShippingGroupID
				break
			}
		}
	} else {
		e.ShippingGroup = cart.Items[0].ShippingGroupID
	}

	if e.ShippingGroup != "" {
		fmt.Println("cart created!", e.ShippingGroup)
		res := e.CheckShippingGroup()

		if res == 200 {
			e.Task.Stage = SubmitShipping
		} else {
			log.Fatal(res)
		}
	}
}

// Required for checkout, if not included then results in some error when submitting order
func (e *PaTask) CheckItems() {

	data := strings.NewReader(e.EncodeCartCheck())
	// fmt.Println(data)
	// var data = strings.NewReader(`{"items":[{"externalPriceQuantity":-1,"externalPrice":39.99,"productId":"000004405","quantity":1,"catRefId":"000004405","locationId":null,"shippingGroupId":"sg99588661"}]}`)
	req, err := http.NewRequest("PATCH", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/items?exclude=embedded.order.shippingGroup%2Cembedded.order.shippingMethod%2Cembedded.order.shippingAddress", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	// req.Header.Set("X-CCVisitorId", "100FUf5GjWZEfUZmKGFaLq7V0G8aWCt25t2asJjUCfxUHMYAEC8")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-71dec6b4:18ee9237d73:-4e05-4094299650")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "bm_sv=D27DE244CED16F4EE117BFD1FF5F4828~YAAQCo0hF22GEeaOAQAA0Eww6RcZzc8Oshd/I0vQe8OoPQ1AnbxC6A7L5yMqtEN4zlPrp0NXmI/iCPeO5DcyooVh8Kvj7l+7JyBGwvgyDSCYhgpQBAdtIwXwXqrufkEjM9FtbYhOQVkAIx4u8kmPp2rVXlWAJPIpwcGEW0zv4CXyuP9wyUR0aRl8xqJg6D74LCfK6dqHCtW0BHaW8JIC8ZbGEj4K8+umGjzdcDQYqy71RDR/FwQKVKY6oa6LRZlNMrxR3DZAlKwS2+Omzzq4NWs=~1; _abck=85A89C236E17A525CCA96B97540A2451~0~YAAQCo0hFwWGEeaOAQAAmEkw6Qt5v8EW3H7pwtZJ3CePkaqfwjWGd008ZJGLq9ViNbCgjecvj4no1KOpCUqibtqjtTNQk3ImDClwsXrQzwwAu1iDm5/edkRwWf0nSXzHC0gWyMI72JabE6WiP/37U504VkD4xiFaKEwwBP5hOPKOzuX2oBspu+pkJ4LnX//RCr8se3tqY7xqyvvnt1B+BG18sqz1WiOQPBblXqLPOdEpSdSxChOIltLtuaUrDj6ckOldDXaymPsrNK9rKtEl6DaY+86TFUraUAuHgka8GzbDK2Zclu0d2kkUyzasud3t0TSwZPT6BUQok5+xUiKLcAJrrEWPIiIrhhrWbVgx3tCuUR5FFk59C5iR8HXVDCzDuQZw6SMq0+BSxfsNy7UCNPV8tnnSnG6mT25oklgqRDc7PhaITpTOveM=~-1~-1~-1; xda8383328c1PRD_siteUS=100FUf5GjWZEfUZmKGFaLq7V0G8aWCt25t2asJjUCfxUHMYAEC8%3A1713309226183%3A1713309226183%3A1713309226183%3A1%3A1; xva8383328c1PRD_siteUS=-71dec6b4%3A18ee9237d73%3A-4e05-4094299650; AGEVERIFY=Over21; __z_a=988891653718383787718383; ak_bmsc=C7D8E44AF22C1E43330DBBA85D4C126D~000000000000000000000000000000~YAAQCo0hFw5/EeaOAQAAJREw6RfyhLJhfy1D4fZIWvKEKdFhba8k1qbQCNZ3lJscVeEfY6SpG+MCGSL6TVMW1CVT4YhXRBsSmzrFxZotTj0ZYn6pN0x4dxTtem7zmv82uic28Z8oqoev4K47NG3bT1jpvd9u+MPbqvSCsg0Mi6oT4JkbJfBl3kfGCAfRw4DzYM37/GOuUltDjoxiC+tjqK/4qsbqRQKMOZQR+Bpep+KXPZK0d5tv1qKrP5twF/WONbM+tB7qCATb3cxVOO2Q1HRhaDoLphPDVuRGMkvEQp3bH+XNcAyJtar70wIO98am3U6TUulWNaCCJ2sCEAl1Lmkvtzuip+ZbbNOC0EwBLGwNwVGKaMZSkfHJlSkB+ocqLwdqfs78i0pk5OTWyHO1VzPGXeKbEfMtJvll4wgdYcdfKGCagPo3XVnBVAtL/lw75h2kBoyNq/qfF/i2tQ==; JSESSIONID=YQ7pMBPRF_5I6HzANZ33fzzvcjBeF-9pvBQK2T-LPidzS5Kwcgat!142620027; __ZEHIC2113=1713309222; ccstoreroute=dfa3d0ae6ef299a99dfcf14f32a604a2|cc2d3787e8999f4a634b9d592c9e2268; BVBRANDID=e7f414c5-7215-4f0c-a0c6-4cef867c764e; BVBRANDSID=bb404882-56e1-4b76-ae14-d11f4ec895d8; bm_sz=D8D94A9FFE35488433D67DE9D34A3463~YAAQCo0hF+Z9EeaOAQAACwYw6RfsBIt0dAwMxntaS5W8+fTiZwm9OFvwOoY4OIlGdpWs/FIL9S2RwfKny/9Qikz79ioecmWmSH/g+/wkwBVUZEyX09FuMlJsPLZ8wTIowc0C24pvMszdBLCcc0QlcwPmiR5PaW65JmblAFsjvUa/8oUZbcFkc+bPrFbeZFNJ4Wi4//l08PP4Q6DFWutR/dsg89jumvZHKgkF0uSFz9SkRFxgl6GM1Ri6FOKGCIABPvw70U5KGDFjGb7CrsKVxcNFCYWQxvC6c0GE30fUynPMAZjhiT/kiZKTwxIFproL4O+IYHjffBW1/ZIoKipq4F0bex6G1/uUmJSK3hLDQtrGRHloacxd5Qswvla8+pHD5kLoxVNKXBmjTmBmQ+ZxOZJ5pWa0~4474181~4408373; osfliveuiroute=e0ed6d79e84c8faff28ef5c16fa91b41|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup")

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for e.Task.Stage != SubmitShipping && e.Task.Stage != SubmitShippingRate {

		resp, err := e.Client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == 200 {
			// e.Task.Stage = SubmitShipping
			e.UpdateStatus(SubmitShipping)

			if e.Mode == "Preload" {
				if e.Precarted {
					// e.Task.Stage = SubmitShippingRate
					e.UpdateStatus(SubmitShippingRate)
				}
			}
		} else {
			e.UpdateStatus(Retrying)

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", body)
		}

	}

}

func (e *PaTask) SubmitShipping() {

	// fmt.Println("Submitting Shipping...")
	url := "https://www.finewineandgoodspirits.com:443/ccstore/v1/orders/current/shippingGroups/" + e.ShippingGroup

	data := strings.NewReader(e.EncodeShippingAddress())
	req, err := http.NewRequest("PATCH", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	// req.Header.Set("X-CCVisitorId", "100FUf5GjWZEfUZmKGFaLq7V0G8aWCt25t2asJjUCfxUHMYAEC8")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-71dec6b4:18ee9237d73:-4e05-4094299650")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var order OrderBody

	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Fatal("Cannot unmarhsal shipping json", err)
	}

	if resp.StatusCode == 200 {
		if order.ShippingAddress.Address1 != "" {
			// fmt.Println("Submitted Shipping")
			// e.Task.Stage = SubmitShippingRate

			e.UpdateStatus(SubmitShippingRate)
			if e.Mode == "Preload" {
				if !e.Precarted {
					e.PreloadId = order.Items[0].CommerceItemID
					// e.Task.Stage = RemoveItem
					e.UpdateStatus(RemoveItem)
				}
			}

		}
	} else {
		fmt.Println("Error Submitting Shipping")
		fmt.Println(string(body))
		// e.Task.Stage = Stop
		e.UpdateStatus(Stop)
	}
}

func (e *PaTask) Login() {
	login := fmt.Sprintf("grant_type=password&username=%s&password=%s", e.Username, e.Password)
	var data = strings.NewReader(login)
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstore/v1/login", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("referer", "https://www.finewineandgoodspirits.com/search?Ntt=Weller%20full%20proof&")
	req.Header.Set("sec-ch-ua", `"Chromium";v="123", "Not:A-Brand";v="8"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("x-cc-frontend-forwarded-url", "www.finewineandgoodspirits.com/search?Ntt=Weller%20full%20proof&")
	req.Header.Set("x-cc-meteringmode", "CC-NonMetered")
	req.Header.Set("x-ccasset-language", "en")
	req.Header.Set("x-ccpricelistgroup", "defaultPriceGroup")
	req.Header.Set("x-ccprofiletype", "storefrontUI")
	req.Header.Set("x-ccsite", "siteUS")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	var token LoginResp
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Fatal("Error Logging in!")
	}

	if resp.StatusCode == 200 {
		e.AccessToken = token.AccessToken
		fmt.Println(e.AccessToken)
	}
}

func (e *PaTask) CurrentOrder() {

	req, err := http.NewRequest("GET", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current?exclude=shippingMethod", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Content-Type", "application/json")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	e.Task.Stage = Stop
}

func (e *PaTask) CheckProfile() {
	var data = strings.NewReader(`{"cart_save_for_later":""}`)
	req, err := http.NewRequest("PUT", "https://www.finewineandgoodspirits.com/ccstore/v1/profiles/current", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/cart")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/cart")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Authorization", "Bearer "+e.AccessToken)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var acc AccountInformation

	err = json.Unmarshal(body, &acc)
	if err != nil {
		log.Fatal("Erorr unmarshalling account json, ", err)
	}

	for _, v := range acc.DynamicProperties {
		if v.ID == "saved_cards" {
			fmt.Printf("v.Value: %v\n", v.Value)
		}
	}

	if acc.ShippingAddress.Address1 != "" {
		fmt.Println(
			"Shipping Address Found: ",
			acc.ShippingAddress.Address1)
	}
	e.Task.Stage = AddToCart
}

// Adds shipping group to order...? I'm not sure but it breaks if i don 't use it
func (e *PaTask) CheckShippingGroup() int {

	url := "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/shippingGroups/" + e.ShippingGroup
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", "_gcl_au=1.1.2111076155.1712769182; GSIDyk8q9gLSz40L=3a46ef7f-0269-4992-b81c-aa23a55257fd; STSID148379=e6f84471-0f49-4ba7-bf3f-0d50fcd8f76c; BVBRANDID=a82f17bc-93fa-4d18-94af-62fda45fc91e; _ga=GA1.1.2050355724.1712769182; __pdst=1b26f91811f54e848b091ce7573f7f59; __adroll_fpc=d8f13915972a0872f73533a05cdab84c-1712769183244; _pin_unauth=dWlkPU5HTTNZakV3TUdJdE16UmtaaTAwWmpaaUxUZ3paR1V0WW1Fek9UZGhaREExTTJVMA; _fbp=fb.1.1712769211703.185896143; ltkpopup-suppression-8e9eb133-da66-408a-880e-6f4844cae736=1; xda8383328c1PRD_siteUS=10C7QZZku5nlm-NsqcjIRupXsbiPr4LkQgg87waqVvXNTBQ9F92%3A1712769183705%3A1713145493278%3A1713310922110%3A14%3A14; __ar_v4=H422MU2HOBC4HLZGUQMCAG%3A20240410%3A36%7CYEV7UUOV4NHJNEIGTN5CUO%3A20240410%3A36; _ga_Y3ZSL11TM0=GS1.1.1713314438.12.0.1713314438.60.0.0; ak_bmsc=DE436D3C343A15932A5129743D951B50~000000000000000000000000000000~YAAQLqfLF+XE5uWOAQAAaFE+7Reoy9OEKq4J7zp3DivRabVIUeXGOF2OPHuzMcoQCONSqzjbHjJbQD2fgruB0emiwekoeLWG9p/t+tGFfXKRDsX15Lx65kt9fXgYkUf+remX/Xxq9R7j89DHGHyygIyEZJ/+mest1COIhsfbJoMd/tPL4TkNsecLs2ECW40AzzfG3aRTqiK0MtAi0c0l469WWos/xN6ZXcDC4ESRX2Ms3KrybUf0QlbbK2s7VcUJ0c+fIW3APvLs0bSmCGSnaRkvb2avDM/j76HsGMfXyOfOM60/Ogq749J96njflbWYIz024ttMNM7g1Qxyl6OkzbkCcWG+8zbfKpK/my1tJiOs0fiKHM7QOE1e72DmAiuvpHhdH1/NO77hwiAs+TKAMaZx1D/oFlv7yY7C; ccstoreroute=48697f4ea9258fc8ee7877afe26607c6|cc2d3787e8999f4a634b9d592c9e2268; JSESSIONID=_FztPlTVaXExEzzL_oul5m3ICnlkOWW3okmKd1gyKTM_zvjzQACj!668367565; _abck=0CFD734DABA18D369949E71FA27D3CC8~-1~YAAQJo0hF1fzlc2OAQAAhFU+7QseNaVnDMpAZw3rnjSdI3AECYEYpWKPKVAenogi0Ng6OLYhwGcMcHBs3YMEefl5za+kt79YiXXdAkgomKuXKIcMdklMm6bD2A6cEPz/KBabcMmOOAWgxIwUbuXieyT6xSnxx/DL92zI6Xh3OCXx5PfXFPMw/8JTuZzHAcIIJjJ8IlMjPRUZV8V93lk7C6Sn+ZvIatFD3ZfFvp158zPlccpiiQgnsJzREzwBms6c0lYgWJA8I5QBJzgokbKYOQEe4X9ydW/EXANlR538TyJFpo1SpTgJjgC+RgNArRYspNK+RJCt2K0de8rwi5jmuZRvmZtELqszoUOVmYtNrXpm+kTCbzogI8u22A90NwKjTvCVqQSwmotnR9jVkyJH5dRw1e/7+VFbxIW7i34HLY3hdNArjVqE8U4=~-1~-1~-1; bm_sv=7489C1F1D12B58DB840348251D15BE2F~YAAQJo0hF1jzlc2OAQAAhFU+7RfiNuZN2CLlUSSuFty+4Ep/QRmtvck/WQ9E6CCQwTSlx62TbjROncpC0UKkeyJmbxgsNnim18OoqZ6RDL5n8KhO8ogMIssEDuOXG5Ynm/XosNGDd/HMFm0px3gSICeLFjCUqfQsRjRcDz3+GqBhvWksYbD67Knw/jMwwILh/zRjzfHVYNYwYHnJmSBkSFg8CHu3lEytKqMhgqdcIImppS6h5I7MghzsNL00YWPIs541QKLdLIwAsl2ezE1Cxw==~1; bm_sz=C0F07C8FF18E5F55374952E51FC3B340~YAAQJo0hF1nzlc2OAQAAhFU+7RekjomNyq+5B9cHSkOyNvD12KOjHx16YMydNdUy4nNoDLfOSFd2AesroR8rBhJmmipJLRu2fYnV8ElhxJImrMewvnet+If9/ryfS0fKmScuT1A6E0PwVV77bFTTBcFqm7n+Mzu2a0zaRf5kQqjytD+/flFmlC/Dxy8Sw9QjqhOtszZ8ypwwyu76uQP4TtQ0g83wKJrwyw1+tiEbw0DsBJZM6iD8LODcj8QvR/iU/kDN0lAhPXiacnRS1yzO3VMJJC6DKbLD3OanNjVPsCru9FvTI993KtET2uecGfHaTEiuk8rbkKmFwt059nqK/8ZpYU3Mzk5Q5o7HLvrThFyZBrPDjV0zHBX9qp4IkoRRu7EMXnTmdhkB3q6LQAdCvCGBlPA/kiLKU+YpIQZPTsukWeoq~3228215~3360056")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := e.Client.Do(req)
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

	// e.Task.Stage = Stop
}

func (e *PaTask) SubmitShippingRate() {

	// fmt.Println("Submitting Shipping Rate...")
	var data = strings.NewReader(`{"shippingMethod":{"value":"sm50001"}}`)
	req, err := http.NewRequest("PATCH", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/shippingGroups/"+e.ShippingGroup, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	// req.Header.Set("X-CCVisitorId", "100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-57489689:18ef340742f:-587b-4094299143")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "bm_sv=19D8AD9EA900EE07070E545555E8DC37~YAAQJo0hFz7+N86OAQAA5plJ8xcypZrTjmnN53TbQFQl9ZvLemSyk3zXGYgbNOgc6VL8btb0vp5aR3Z73sUy5YgyMLpZVZVXNrrA5pUG/zWX0DsSGLobUdGyQynHj+/1oGYvOPC9t63/Knumm0lt+9dCSSkRxrDM/gimncvGreos+W1oLom9PRrQMfiz6cxZPZjtiWI3yJ/JU+RJNET2xRawjHq6B636BewH9e6Zl+bvKeGCVH1KkqUTn1P51qTWwjtSxMiV22OmNN089YWVzuA=~1; _abck=182E4ECB85F1A83863E7512B441E2AEC~0~YAAQJo0hFwzpN86OAQAAmgJJ8wuI5JgQu4LQ7nPCBmSksqoXlUXJsclynhZwmzqplXLBSrVXo1tOZGCVVayA9SjeK8ebedVxL2Wd4zpGTdASodkz+kjt6OedIpjMi+6Rc+JwBzhMSsoV1/zxtADA+ic2mWHoxpu7E0EMKXlLvPev/C30EMnBy6AC9CTbYbTLd2He7pEAevBMT6aaTbaiV3kuI37xwwJgEzJtmpwqSLufMbABkoZtAlxZWj+Wg5DS6wu3o95XBmTXsrb/JKZtYG0qswzXHXWwOsCClVqDhRncINbFaMQ4tTA/M2Pv7NvV2tbCSkIbCj5TMlOumax3NhUIPGAgKHdZFAAy6wdcmRY8Vn/py98JW3cOdZ7WoscDs8GCEPIZIa40eG7LSwuTnxgfOkjbFnoXZXj2RLo1CqXljcv61ASmY1Q=~-1~-1~-1; AGEVERIFY=Over21; __ZEHIC6741=1713478309; xda8383328c1PRD_siteUS=100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8%3A1713478616127%3A1713478616127%3A1713478616127%3A1%3A1; xva8383328c1PRD_siteUS=-57489689%3A18ef340742f%3A-587b-4094299143; __z_a=227309176116686791411668; JSESSIONID=O7DzSMH3a-jgqNwgVZbA0qjgsPVk2A13koFSMypZj9fVsCPO-yZl!142620027; ccstoreroute=dfa3d0ae6ef299a99dfcf14f32a604a2|cc2d3787e8999f4a634b9d592c9e2268; ak_bmsc=D412863F7D676C87B610C746234B923D~000000000000000000000000000000~YAAQJo0hF17gN86OAQAAyr9I8xeqjiHLw8Xy27/NFFV7jBu3ZSNiJWRVxrTn9qKIOtEBTrNdfHRlhWiQ8Do+u+oOWU5sM7lgM1xXuhUz5G8kWwJHstjzVbeX7WMKEM1j4ZWS89gfoXPLPN/0gfjHwkksCOqm00wJ5c7IS3ba3gVoKcWJAvcKGHWk4LKH0ZkZgkyYS3QQXWY4gIfWFs3S9r62pykcy7cq9vN+2wrr8Yl0egdC04dIacATC8tO66yOrie/judlRM6+jJSViK6iWx9xLeQ+4EURCo7q56qB7TUjj4reIHO5+hlUv1w/uIwzY8Oema9E/o5yzm84GABwVCuAVeuoog9ikxeIlZUejL3S3mR8/hb/dSwjtVsBLEdRicTIS42Fwq0/lL0WNQyZTwjCIUtwHOjKiOCwkRJFSgnR2UEjcbCvthHf32o5keXIM5Gf+oFanr64ImMTkw0=; BVBRANDID=6455d91a-4707-4ab9-8e4b-852da571249a; BVBRANDSID=78597cee-9058-41b0-83eb-784025667542; bm_sz=9DEEE650DA275694A269D8BD44118C5D~YAAQJo0hF4LfN86OAQAAK7lI8xePpWlfhBNoHjE5Ve3PrXRRoiFnDPMESLt5UHY3NQvc+VImWev8nk7N2+QD4uuUW8mZAT/Q4icOxE+SSldOxvpBXQt8KxRofvEGYupXczpgANaRKNOVSszp8Rz0GN7JvNNdjfKaC3uWyGbmhOW4j+ubaGaKkzmNOQd8xfQxYYGvVpkbMpeeJ2W6LlS0d6drEzNuM/RBMnbNkAhe0TYMaLOosOOymlI29aEH2hkE4aNmJEJw9UeIg/O129BPXDxkeBx6kqr497fk2wJXXALrtSvBIskUqoe9Bo1ySDSNUmZkipHlTpeY7rvOCRu6ebRjAdEEPJgIUtTAjSilCq8qX1eegNtjKqdYAgSgKGG79nZK43CFXgUKhrOstP5df+ysU4Gpaw==~3228985~3618097; osfliveuiroute=89df68cb9702ba9bbc692b60aeba1cf2|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var order OrderBody

	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Fatal("Error unmarshalling shipping rate POST", err)
	}

	if resp.StatusCode == 200 {
		// log.Println("Submitted Shipping Rate")
		e.OrderTotal = order.PriceInfo.Total
		// fmt.Println(e.OrderTotal)
		// e.Task.Stage = EncryptCard
		e.UpdateStatus(EncryptCard)
		// e.Task.Stage = Stop

	} else {
		log.Println("Error Submitting shipping rate")
		log.Fatal(resp.Status)
		// e.Task.Stage = Stop
		e.UpdateStatus(Stop)
	}

}

// Just makes a request to a 3rd party service to tokenize the card number to submit to the order.
// The encrypted number is always the same so once it's used it can be skipped
func (e *PaTask) EncryptCard() {
	// fmt.Println("Encrypting Card...")
	// var data = strings.NewReader(`{"account":"4444555555555555","source":"iToke","encryptionhandler":null,"unique":false,"expiry":null,"cvv":"998"}`)
	data := strings.NewReader(e.EncodeCardStruct())
	req, err := http.NewRequest("POST", "https://palcb.cardconnect.com/cardsecure/api/v1/ccn/tokenize", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Host", "palcb.cardconnect.com")
	req.Header.Set("Origin", "https://palcb.cardconnect.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "113")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var card EncryptedCard
	err = json.Unmarshal(body, &card)
	if err != nil {
		fmt.Println(card.Message)
		log.Fatal("Error Encrypting card", err)
	}

	if card.Errorcode == 0 {
		// fmt.Println("Card Encrypted...")
		e.EncryptedCard = card.Token
		// e.Task.Stage = SubmitOrder
		e.UpdateStatus(SubmitOrder)
	} else {
		// e.Task.Stage = Stop
		e.UpdateStatus(Stop)
	}

}

// Adds payment to the order - also not requried because it submits the exact same information as when the order is submitted lmfao
func (e *PaTask) AddPayment() {
	var data = strings.NewReader(`{"items":[{"billingAddress":{"lastName":"GORW","country":"US","address3":"","address2":"","city":"PHILADELPHIA","prefix":"","address1":"123 MAIN ST","jobTitle":"","companyName":"","postalCode":"19127-2148","county":"","suffix":"","firstName":"Bhiwgriuherg","phoneNumber":"3938402900","faxNumber":"","middleName":"","state":"PA","email":"qgwroiuherg@gmail.com","company":null},"cardNumber":"9443550680705555","cardType":"visa","expiryMonth":"09","expiryYear":"29","nameOnCard":"Lenix Woof","customProperties":{"token":"9443550680705555","type":"visa","exp":"0929"},"amount":53.99,"type":"generic"}]}`)
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/payments/add?exclude=embedded.order.shippingGroup%2Cembedded.order.shippingMethod%2Cembedded.order.shippingAddress", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("Cookie", "bm_sv=19D8AD9EA900EE07070E545555E8DC37~YAAQJo0hF4jxQs6OAQAADf6m8xcHU1BHz1VmeIZNvYfUSJdpQQsXv2d2s0YFBG5sYSrVJedfL2Jck0GZ2GjRPh+vCOWuK6DD2S3R/WKKLBAcdgHmA1/xYPlU7ENp/SOmvFurBE1Ozs/maZnwJQlhXIniTB7Nw9lfe3rsIkrRPXKFVch1Kyfze2wpbtOycfiWWuFdFWayUiTklet71qUYR53EZt2EmhaYJyKCld/6PIReth+moUuLJ2+1mpD0O+x4qU15toB+hcv/ld9syV0+QIQ=~1; _abck=182E4ECB85F1A83863E7512B441E2AEC~0~YAAQJo0hF8+3Qs6OAQAAL4il8wvyk4QhAxIBlJYtCtEvUKB1g8x4DPtUcUWMzjVmMGKYuwCb7kczQ1L660hwyMIMKmV2uGP5sHvVxnq0SNAZ5OEkK6eA9J3ZgvaWiVb/xqg4XWJFQnGShCnAxZyyQ5XJn4VkYxCvURPmpKO51ghHNpYyE/ZIuYwSwBDwQuAX8yB3gFTKMHoBtZPoashTGKb8Pkxl4lizh4+q/7rMYdeeWnitoxGTKqZ7OQCQCRDy0UQvoNn+u0FZUL5tpfgNfFZ+00f/G+O+wBLv9pE3sWhtlzIfDvm5k3b1xZV8ineQeO+zmXVTWlzKFvW6sGq9bd+mEIXWhMfrU9zyqO9hJq1OA/LGSYB3KHj95pSQh8lklFrOLFmiBfmytF3YilUChvKdYEffMi5vktaVfh0GJHHj+ZkuatLH2Ow=~-1~-1~-1; xda8383328c1PRD_siteUS=100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8%3A1713478616127%3A1713478616127%3A1713478616127%3A1%3A1; xva8383328c1PRD_siteUS=-48eff850%3A18ef3632fbd%3A-41d7-4094298630; JSESSIONID=CcTzpOUU3R9oIrOqlL4kuD3qu4wDtXwfSMJUN7vXdli9KhfJYkVr!142620027; AGEVERIFY=Over21; ccstoreroute=dfa3d0ae6ef299a99dfcf14f32a604a2|cc2d3787e8999f4a634b9d592c9e2268; ak_bmsc=D412863F7D676C87B610C746234B923D~000000000000000000000000000000~YAAQJo0hF17gN86OAQAAyr9I8xeqjiHLw8Xy27/NFFV7jBu3ZSNiJWRVxrTn9qKIOtEBTrNdfHRlhWiQ8Do+u+oOWU5sM7lgM1xXuhUz5G8kWwJHstjzVbeX7WMKEM1j4ZWS89gfoXPLPN/0gfjHwkksCOqm00wJ5c7IS3ba3gVoKcWJAvcKGHWk4LKH0ZkZgkyYS3QQXWY4gIfWFs3S9r62pykcy7cq9vN+2wrr8Yl0egdC04dIacATC8tO66yOrie/judlRM6+jJSViK6iWx9xLeQ+4EURCo7q56qB7TUjj4reIHO5+hlUv1w/uIwzY8Oema9E/o5yzm84GABwVCuAVeuoog9ikxeIlZUejL3S3mR8/hb/dSwjtVsBLEdRicTIS42Fwq0/lL0WNQyZTwjCIUtwHOjKiOCwkRJFSgnR2UEjcbCvthHf32o5keXIM5Gf+oFanr64ImMTkw0=; BVBRANDID=6455d91a-4707-4ab9-8e4b-852da571249a; bm_sz=9DEEE650DA275694A269D8BD44118C5D~YAAQJo0hF4LfN86OAQAAK7lI8xePpWlfhBNoHjE5Ve3PrXRRoiFnDPMESLt5UHY3NQvc+VImWev8nk7N2+QD4uuUW8mZAT/Q4icOxE+SSldOxvpBXQt8KxRofvEGYupXczpgANaRKNOVSszp8Rz0GN7JvNNdjfKaC3uWyGbmhOW4j+ubaGaKkzmNOQd8xfQxYYGvVpkbMpeeJ2W6LlS0d6drEzNuM/RBMnbNkAhe0TYMaLOosOOymlI29aEH2hkE4aNmJEJw9UeIg/O129BPXDxkeBx6kqr497fk2wJXXALrtSvBIskUqoe9Bo1ySDSNUmZkipHlTpeY7rvOCRu6ebRjAdEEPJgIUtTAjSilCq8qX1eegNtjKqdYAgSgKGG79nZK43CFXgUKhrOstP5df+ysU4Gpaw==~3228985~3618097; osfliveuiroute=89df68cb9702ba9bbc692b60aeba1cf2|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Content-Length", "598")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	// req.Header.Set("X-CCVisitorId", "100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-48eff850:18ef3632fbd:-41d7-4094298630")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	resp, err := e.Client.Do(req)
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

func (e *PaTask) SubmitOrder() {
	// fmt.Println("Submitting Order")

	// var data = strings.NewReader(`{"payments":[{"billingAddress":{"lastName":"GORW","country":"US","address3":"","address2":"","city":"PHILADELPHIA","prefix":"","address1":"123 MAIN ST","jobTitle":"","companyName":"","postalCode":"19127-2108","county":"","suffix":"","firstName":"Bhiwgriuherg","phoneNumber":"3938402900","faxNumber":"","middleName":"","state":"PA","email":"qgwroiuherg@gmail.com","company":null},"cardNumber":"9443550680705555","cardType":"visa","expiryMonth":"09","expiryYear":"29","nameOnCard":"Lenix Woof","customProperties":{"token":"9443550680705555","type":"visa","exp":"0929"},"amount":53.99,"type":"generic"}]}`)
	data := strings.NewReader(e.EncodeBillingInformation())
	req, err := http.NewRequest("POST", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/submit", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2.1 Safari/605.1.15")
	// req.Header.Set("X-CCVisitorId", "100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/checkout")
	req.Header.Set("X-CCSite", "siteUS")
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/checkout")
	// req.Header.Set("X-CCVisitId", "-57489689:18ef340742f:-587b-4094299143")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "JSESSIONID=86zzVJLEhDIsPV4c3dU_iy6_uTqz_IECrS6wHKWgbZsb5eAo1gR2!142620027; bm_sv=19D8AD9EA900EE07070E545555E8DC37~YAAQIkjIF5kUdMuOAQAADpJU8xd+W2jx6fwnOTLkcYYG/3h79+j9MbgzrBbazqlg2xoMYIJq4/FBoEsPGZDYkP1h25Gh0SDulhTzGa6TR2W4nkXsOvva0ojl3SYmZaMGzO0shYkqI6CYMjUVGFNmkfeOFg3hjUtVLGFNx4bVAHhl+R5iD6BOKYbHNk7ahQOHQN7DGRuFH6E5Qei/liIfxyJE27KfOGbGUrxzVLlhBzgYMfGD8zCvBisZ4ivxVnl7QdKYjd7zW6/RbNvciSpl/tc=~1; _abck=182E4ECB85F1A83863E7512B441E2AEC~0~YAAQJo0hFwzpN86OAQAAmgJJ8wuI5JgQu4LQ7nPCBmSksqoXlUXJsclynhZwmzqplXLBSrVXo1tOZGCVVayA9SjeK8ebedVxL2Wd4zpGTdASodkz+kjt6OedIpjMi+6Rc+JwBzhMSsoV1/zxtADA+ic2mWHoxpu7E0EMKXlLvPev/C30EMnBy6AC9CTbYbTLd2He7pEAevBMT6aaTbaiV3kuI37xwwJgEzJtmpwqSLufMbABkoZtAlxZWj+Wg5DS6wu3o95XBmTXsrb/JKZtYG0qswzXHXWwOsCClVqDhRncINbFaMQ4tTA/M2Pv7NvV2tbCSkIbCj5TMlOumax3NhUIPGAgKHdZFAAy6wdcmRY8Vn/py98JW3cOdZ7WoscDs8GCEPIZIa40eG7LSwuTnxgfOkjbFnoXZXj2RLo1CqXljcv61ASmY1Q=~-1~-1~-1; AGEVERIFY=Over21; __ZEHIC6741=1713478309; xda8383328c1PRD_siteUS=100FtByoZ60K23lwFE0kvPcxIV7hQlU3SDV_G1r25uxP5vE6FF8%3A1713478616127%3A1713478616127%3A1713478616127%3A1%3A1; xva8383328c1PRD_siteUS=-57489689%3A18ef340742f%3A-587b-4094299143; __z_a=227309176116686791411668; ccstoreroute=dfa3d0ae6ef299a99dfcf14f32a604a2|cc2d3787e8999f4a634b9d592c9e2268; ak_bmsc=D412863F7D676C87B610C746234B923D~000000000000000000000000000000~YAAQJo0hF17gN86OAQAAyr9I8xeqjiHLw8Xy27/NFFV7jBu3ZSNiJWRVxrTn9qKIOtEBTrNdfHRlhWiQ8Do+u+oOWU5sM7lgM1xXuhUz5G8kWwJHstjzVbeX7WMKEM1j4ZWS89gfoXPLPN/0gfjHwkksCOqm00wJ5c7IS3ba3gVoKcWJAvcKGHWk4LKH0ZkZgkyYS3QQXWY4gIfWFs3S9r62pykcy7cq9vN+2wrr8Yl0egdC04dIacATC8tO66yOrie/judlRM6+jJSViK6iWx9xLeQ+4EURCo7q56qB7TUjj4reIHO5+hlUv1w/uIwzY8Oema9E/o5yzm84GABwVCuAVeuoog9ikxeIlZUejL3S3mR8/hb/dSwjtVsBLEdRicTIS42Fwq0/lL0WNQyZTwjCIUtwHOjKiOCwkRJFSgnR2UEjcbCvthHf32o5keXIM5Gf+oFanr64ImMTkw0=; BVBRANDID=6455d91a-4707-4ab9-8e4b-852da571249a; BVBRANDSID=78597cee-9058-41b0-83eb-784025667542; bm_sz=9DEEE650DA275694A269D8BD44118C5D~YAAQJo0hF4LfN86OAQAAK7lI8xePpWlfhBNoHjE5Ve3PrXRRoiFnDPMESLt5UHY3NQvc+VImWev8nk7N2+QD4uuUW8mZAT/Q4icOxE+SSldOxvpBXQt8KxRofvEGYupXczpgANaRKNOVSszp8Rz0GN7JvNNdjfKaC3uWyGbmhOW4j+ubaGaKkzmNOQd8xfQxYYGvVpkbMpeeJ2W6LlS0d6drEzNuM/RBMnbNkAhe0TYMaLOosOOymlI29aEH2hkE4aNmJEJw9UeIg/O129BPXDxkeBx6kqr497fk2wJXXALrtSvBIskUqoe9Bo1ySDSNUmZkipHlTpeY7rvOCRu6ebRjAdEEPJgIUtTAjSilCq8qX1eegNtjKqdYAgSgKGG79nZK43CFXgUKhrOstP5df+ysU4Gpaw==~3228985~3618097; osfliveuiroute=89df68cb9702ba9bbc692b60aeba1cf2|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup")
	e.CheckoutEndTime = time.Now()
	e.TaskEndTime = time.Now()
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var conf OrderResponse
	err = json.Unmarshal(body, &conf)
	if err != nil {
		fmt.Printf("%s\n", body)
		log.Fatal("Error unmarshalling order submission json", err)
	}

	d := discord.PaDiscordData{
		ProductName:   e.ProductName,
		ProductImage:  e.ProductImage,
		ProductVol:    e.ProductVol,
		Subtotal:      fmt.Sprintf("%2f", e.OrderTotal),
		DeclineReason: conf.Payments[0].Message,
		Mode:          e.Mode,
		Fulfillment:   e.Fulfillment,
		ProfileEmail:  e.Profile.UserInformation.Email,
		TaskTime:      e.TaskEndTime.Sub(e.TaskStartTime).Seconds(),
		CheckoutTime:  e.CheckoutEndTime.Sub(e.CheckoutStartTime).Seconds(),
		ProfileName:   e.Profile.ProfileName,
		OrderNumber:   conf.ID,
	}

	if resp.StatusCode == 200 {
		// fmt.Println(conf.Payments[0].Message)
		if conf.State == "INCOMPLETE" {
			discord.PaLiquorDeclineWebhook(d)
			e.UpdateStatus(PaymentDeclined)
			fmt.Printf("%s\n", body)
			// fmt.Println(d)
			// fmt.Println(e)

		} else if conf.State == "SUBMITTED" {
			e.UpdateStatus(Success)
			discord.PaLiquorCheckoutWebhook(d)
			fmt.Printf("%s\n", body)

		}
	} else {
		fmt.Printf("%s\n", body)
	}
	// e.Task.Stage = Stop
	e.UpdateStatus(Stop)

}

func (e *PaTask) DeleteItem() {
	req, err := http.NewRequest("DELETE", "https://www.finewineandgoodspirits.com/ccstore/v1/orders/current/items/"+e.PreloadId, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="99", "Chromium";v="124"`)
	req.Header.Set("DNT", "1")
	req.Header.Set("X-CC-MeteringMode", "CC-NonMetered")
	req.Header.Set("X-CCPriceListGroup", "defaultPriceGroup")
	// req.Header.Set("X-CCVisitorId", "1216RwFSsmukH482YHIM9h0OedAxEadp1qEbuySGBk4vMb06995")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("X-CC-Frontend-Forwarded-Url", "www.finewineandgoodspirits.com/cognac/124")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CCSite", "siteUS")
	// req.Header.Set("X-CCVisitId", "-71dec6b4:18f2b1051d8:-73a0-4094299650")
	req.Header.Set("X-CCProfileType", "storefrontUI")
	req.Header.Set("X-CCAsset-Language", "en")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Origin", "https://www.finewineandgoodspirits.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.finewineandgoodspirits.com/cognac/124")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Cookie", "osfliveuiroute=2b53c2803d1a269a3b8748149ee6aff3|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup; BVBRANDID=4bebb429-a543-4fd8-b8e9-4c9aaeb027da; BVBRANDSID=b4757513-3030-40f3-9e23-2830f9e15649; _gcl_au=1.1.1325948082.1714414583; _ga=GA1.1.706909971.1714414584; __pdst=0c6ad1422b724907b6c593a75b3d4dd2; GSIDyk8q9gLSz40L=90ff9cdf-b0bc-4a09-951d-16d76cdafeb5; STSID148379=6083d321-95ca-4529-8625-1ff412fd4057; __ZEHIC4573=1714411220; _fbp=fb.1.1714414583953.584478619; _pin_unauth=dWlkPU5qYzNOemt5WVRjdE5UVmlZUzAwTnpVMExUZ3hNRFV0TjJaaE9XUm1ORGN5WVRrMA; __adroll_fpc=81c83151d4a525fe48d9d863a5cd8a1a-1714414584077; _vuid=3698737b-ec7b-4caa-a3e2-8743de5c8a23; ltkSubscriber-Footer=eyJsdGtDaGFubmVsIjoiZW1haWwiLCJsdGtUcmlnZ2VyIjoibG9hZCIsImx0a0VtYWlsIjoiIn0%3D; xda8383328c1PRD_siteUS=1216RwFSsmukH482YHIM9h0OedAxEadp1qEbuySGBk4vMb06995%3A1714414586252%3A1714414586252%3A1714414586252%3A1%3A1; xva8383328c1PRD_siteUS=-71dec6b4%3A18f2b1051d8%3A-73a0-4094299650; AGEVERIFY=Over21; __z_a=3204214719100996681010099; ccstoreroute=9bdf8b85f993f33582f4c9359f5fbbc1|cc2d3787e8999f4a634b9d592c9e2268; JSESSIONID=dj0rEpYhOQK0MudvgsIOnTK22n2ea7luru62HnU9LSPDmMBFoqn1!1062285647; ltkpopup-suppression-8e9eb133-da66-408a-880e-6f4844cae736=1; ltkSubscriber-Checkout=eyJsdGtDaGFubmVsIjoiZW1haWwiLCJsdGtUcmlnZ2VyIjoibG9hZCIsImx0a0VtYWlsIjoiZWxpamFzc2xpdmEwQGdtYWlsLmNvbSIsImx0a09wdEluIjoib2ZmIiwiZmlyc3RuYW1lIjoiRWxpamFzIiwibGFzdG5hbWUiOiJTbGl2YSJ9; __ZEHIC6265=1714414780; bm_mi=07D18BAE35B76C11EB3D7A0785C7888C~YAAQIkjIFz83tSGPAQAAjeoVKxfpxBXVhDSoZoHrJyhpQ/5++CyLHos9cN2fiviehnnhiM/mT97jy5CXHQB4cZ0/qQ6SCJbik8Qvi1S2G2NxgFHZsvCLm9uPu10aJjCDDLErbgH9QmCmEl6oBNuw3uQ+relCHZIWr7A4+XXTx98VqdKcfZzkDM5z4UgZaONbyxP8TaoiEQ1FfLMdPODN1PHa1wAKb5ZzeRLGIxRC+OYIZ2WhiROT46X7Ga9VmHLz6Cu4eWzpmLbzbPxEpqO7/cjY3aua0PwNJo0u51jcJgl3m/IJuer1PiK14O2bwHc4rO6cXNAysX61zk4/EzeH8VVDIdwz~1; ak_bmsc=230CA6E7676B184C22599D611AC2E75A~000000000000000000000000000000~YAAQIkjIF143tSGPAQAAJ+0VKxfE1uhswrhIoTioT3QL6UhvUCj6LZFAlrd35jLic2rtJSiPqlA0B4T1fhRd7jXCke1EiRPFm+SWcDOxSVOCZSDeObHKi7YzLp0/sZpAXFXy0j00+wQ31daybdzz75CRWVAxtAm2lIqYu2ZrYgYJlTx1LoJT3bCd9Lq4U0X8KO28oTQvt2a5jehkot8O87DX5V/OaearLPdrmkinsFoPAM7dieOK6vIEB9CSWve9ExQpVNqEzAuAaalg/GnLtjSz0E6AZCaCh5JmlAFqQ4SxAbrdVkpDO6A2RugQE+vDsalGOtFD9v5uJtTICg4gxQaWGqxpxQhj9f/P2G1XJyc0pmdQ7dhwFyhk7/m8ib5edj/y7CWKQHef2c0ATbyy0gw74yvP33XtW6OKgLZPpscCzZAj3iMb/0R2ZjCDQmPQcShStadIYl874bvUju7mFGT0Df9daUmxq2YkFGYQ39Jug7HWqVtm; __ZEHIC6657=1714414807; bm_sz=6506D01F5461019F29E972C83CFC0133~YAAQIkjIF39itSGPAQAAHxwZKxfxDaHMWHNi7dFRw09urSQk94/XXyLlPz+68N8UdxSMm4eSHv7VkQm5XMGKj5g4zF9tR6YrW8HvlhZavUO80GU9nlmKrAaHgfzydxhEI1A7Ka/mDotPDhpKtNbQEsDdn6CxNH7JHDdvu9wsrU+t+eK2l9tEQkYGRLtAjhkrUcL8clFnNXXHuWbmV0mTlboH7sa7mcyUDnpDCyMdPKMWkHOCPPUQLDmikSmL0BXPWx2pBlp+URsy02mG1ZmhyaRM/eu8W4O8zAjOk6Kb4Omk99etqz7Wwse8HrGzE3c9afrki1z7yWpbAnuYlAwaK9oC/FrVNQIPYgt2bKFulPJOjBFd9J3l9fg3g4EBhhokk1Tw+7Zwau/YazMyyx3wzISNe7B65jHnOA27hU5eA8i3N9ZzWurG6FLJbyo=~3683632~3289397; _ga_Y3ZSL11TM0=GS1.1.1714414583.1.1.1714415017.59.0.0; ltkpopup-session-depth=3-7; __ar_v4=H422MU2HOBC4HLZGUQMCAG%3A20240429%3A4%7CYEV7UUOV4NHJNEIGTN5CUO%3A20240429%3A4; _abck=B9A839AC42EB2E299566ED6D2551A3DD~-1~YAAQIkjIF8tjtSGPAQAAZTMZKwuo2Rc3wfSzXYPyMRJKqJSiPaY6lZyv5JxLr7v8Z2+EtwYx+3oPdKzrNkoplSUs7KAyUwK00Ojla/ljvCPnw4FruucuFfrRnZHOb8ot76G3r7hjBxNbHpSklguHma2r2IgWFwLMkPpKFsT/FBH2N70wsDV73uBDfFtD2uNpNIfD2hlbB2wVvNUKZmARnDxD9PDuSzNyQeEigZlulGpLQPAzEKp1Hs8EmS1QNsO2wsKAQbfnsB0Bb8A6aLNLnBRM4VZDZlqI5SLTy2zZrfY2guiy3p0Jo6wO1Gi86YFoprlBHk15CZO5hdY4TKx+wUVr5SbxFTU+7Nd+pKMPBAHgtP/L7D1RqlR8HcVkUmj1Vcid24BJZUeujv9rUid/2FugrEibrDrg+z92zfF/3Ux3f0egpThXzA==~0~-1~-1; bm_sv=1817989E96CD35A788268BC7B312B653~YAAQIkjIF8xjtSGPAQAAZTMZKxfdGnGyaCwFIQxCW4gaUDjSIQayZ2XaeND3StDwbMReG39Btaj0B+qh1RTNz0TMaAKHa/AXpQImhUglLKNIufRhDBARzUw87meL/ANekSehh00FCFijbbjLnij+9V6OFcmJlGs4Jnbnztm4Hpbdo3FV3qeurwCTJmQR/09a3sZ3loZDGJVEE1aWBoWdxHxqxKD0w0r0odogfUjhCE93Qs7juo4BYC11iLWSBEYNILf0WBXXpE3rTgYxCNopSXg=~1")
	resp, err := e.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var order AddToCartResp

	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Fatal("error unmarshalling delete response", err)
	}

	if len(order.Items) == 0 && order.Embedded.Order.ShippingGroups[0].ShippingAddress.Address1 != "" {
		e.Precarted = true
		fmt.Println("cart is empty")
		e.Task.Stage = Monitor
	} else {
		e.Task.Stage = Stop
	}

}
