package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetOwnerHotel(raw string) (string, string, error) {

	regexPattern := `https://book\.passkey\.com/event/(\d+)/owner/(\d+)/home`
	re := regexp.MustCompile(regexPattern)

	matches := re.FindStringSubmatch(raw)

	if len(matches) < 3 {
		return "", "", fmt.Errorf("cannot arse url %s", raw)
	}

	EventID := matches[1]
	OwnerID := matches[2]

	return EventID, OwnerID, nil
}

func GetEventTitle(doc *goquery.Document) string {

	str := doc.Find("#info_eventName")
	return str.Text()

}

func GetEventFromSplash(doc *goquery.Document) string {
	str := doc.Find("#header-msg").Text()

	strClip := strings.TrimSpace(str)

	return strClip
}

func ParseIDfromJS(doc *goquery.Document) (string, string) {

	pattern := regexp.MustCompile(`/event/(\d+)/owner/(\d+)/rooms/select`)
	tag := doc.Find("#search")
	v, e := (tag.Attr("action"))

	if e {
		matches := pattern.FindStringSubmatch(v)

		e := matches[1]
		o := matches[2]
		return e, o
	}

	return "", ""

}

func GetCSRFToken(jar http.CookieJar) string {

	u, _ := url.Parse("https://book.passkey.com")
	cookies := jar.Cookies(u)
	var token string

	for _, i := range cookies {
		if i.Name == "XSRF-TOKEN" {
			token = i.Value
		}
	}
	return token

}

func GetAttendeeOptions(doc *goquery.Document) []AttendeeOptions {
	var opts []AttendeeOptions

	dropdown := doc.Find("#groupTypeId")

	dropdown.Find("option").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if !exists {
			fmt.Println("Value attribute not found for option", i)
		}
		text := s.Text()
		a := AttendeeOptions{
			Title: text,
			Value: value,
		}

		opts = append(opts, a)
	})

	if len(opts) == 0 {
		groupID := doc.Find("#groupTypeId")
		v, e := groupID.Attr("value")

		a := AttendeeOptions{
			Title: "",
			Value: v,
		}

		if e {
			opts = append(opts, a)
		}

	}

	return opts
}

func ParseHotels(doc *goquery.Document) HotelSearchResult {

	hotelJson := doc.Find("#last-search-results").Text()
	// fmt.Println(hotelJson)
	var results HotelSearchResult

	//Errors are from a mismatch of int / float types from demo test, will fix if error is still around w/ real stuff
	err := json.Unmarshal([]byte(hotelJson), &results)
	if err != nil {
		fmt.Println("Cannot parse hotels into struct", err)
		// fmt.Println(hotelJson)
	}

	return results
}

func CreateUpdateJson(hotelID int) UpdateStruct {

	b := Block{
		HotelID:          0,
		BlockID:          0,
		CheckIn:          nil,
		CheckOut:         nil,
		NumberOfGuests:   1,
		NumberOfChildren: 0,
		NumberOfRooms:    1,
	}

	bs := []Block{b}

	bm := blockMap{
		Blocks:      bs,
		TotalRooms:  1,
		TotalGuests: 1,
	}

	var str []string

	t := UpdateStruct{
		MultiHotelRoom: false,
		HotelID:        strconv.Itoa(hotelID),
		DistanceEnd:    0,
		MaxGuests:      5,
		HotelIds:       str,
		BlockMap:       bm,
		MinSlideRate:   0,
		MaxSlideRate:   0,
		WlSearch:       true,
		ShowAll:        true,
		Mod:            false,
	}

	return t
}

func ParseQueueURL(raw string) string {
	regex := regexp.MustCompile(`^(https?://[^/]+)`)
	match := regex.FindString(raw)
	fmt.Println(match)

	return match

}

func URLType(raw string) string {

	url, err := url.Parse(raw)
	if err != nil {
		fmt.Println("Not A URL!")
	}

	domain := url.Hostname()
	return domain
}

func GetUserID(doc *goquery.Document) string {

	metaTag := doc.Find("#queue-it_log")
	userID, exists := metaTag.Attr("data-userid")
	if !exists {
		fmt.Println("data-userid attribute not found")
		return ""
	}

	return userID

}

func InList(list []string, key string) bool {

	for _, v := range list {
		if key == v {
			return true
		}
	}
	return false
}

func ListInList(targetList []string, keyList []string) bool {

	if len(keyList) == 0 {
		return false
	}

	hit := 0
	for _, v := range targetList {
		for _, b := range keyList {
			if b == v {
				hit++
			}
		}
	}

	// if hit != len(keyList) {
	// 	return false
	// } else {
	// 	return true
	// }

	if hit != 0 {
		return true
	} else {
		return false
	}
}

func GetAcklowedgement(doc *goquery.Document, csrf string) string {

	var ackStr string
	_csrf := "_csrf=" + csrf

	confSec := doc.Find("#resAggForm")
	divs := confSec.Find("div")
	divs.Each(func(i int, div *goquery.Selection) {
		id, e := div.Attr("id")
		if e {
			split := strings.Split(id, "-")
			if split[0] == "confirmation" {
				input := div.Find("input")
				input.Each(func(j int, input *goquery.Selection) {
					name, _ := input.Attr("name")
					val, e := input.Attr("value")
					if e {
						ackStr += name + "=" + val + "&"
						// fmt.Println(name)
						// fmt.Println(val)
					}
				})

			}

		}
	})

	ackStr = strings.ReplaceAll(ackStr, "[", "%5B")
	ackStr = strings.ReplaceAll(ackStr, "]", "%5D")
	ackStr = strings.ReplaceAll(ackStr, "-", "%2F")
	ackStr = strings.ReplaceAll(ackStr, "@", "%40")

	ackStr += _csrf

	// fmt.Println(ackStr)

	return ackStr
}

func GetPaymentVal(doc *goquery.Document) string {

	attr := doc.Find("#splitFolio0")
	val, e := attr.Attr("value")

	if e {
		return val
	} else {
		fmt.Println("Cannot Parse Payment Value")
		return ""
	}
}

func GetKws(kwString string) ([]string, []string) {

	var posKws []string
	var negKws []string

	s := strings.Split(kwString, ",")

	for _, v := range s {
		if string(v[0]) == "+" {
			posKws = append(posKws, v[1:])
		} else {
			negKws = append(negKws, v[1:])
		}

	}
	return posKws, negKws
}

func ParseConfNumber(doc *goquery.Document) string {

	var span string
	doc.Find(".ack-details").Each(func(i int, s *goquery.Selection) {
		span = s.ChildrenFiltered("span").Text()
	})

	split := strings.Split(span, ":")
	ack := strings.ReplaceAll(split[1], " ", "")

	return ack

}

func ContainsKeywords(s string, positiveKeywords []string, negativeKeywords []string) bool {
	// Check if string contains all positive keywords
	for _, keyword := range positiveKeywords {
		if !strings.Contains(s, keyword) {
			return false
		}
	}

	// Check if string contains any negative keywords
	for _, keyword := range negativeKeywords {
		if strings.Contains(s, keyword) {
			return false
		}
	}
	return true
}
