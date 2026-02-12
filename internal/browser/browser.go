package browser

import (
	"net/http"
	"time"

	"github.com/go-rod/rod"
)

const layout = "2006-01-02 15:04:05 -0700 MST"

func StartRod(r *rod.Browser) []*http.Cookie {

	// page := rod.New().ControlURL(u).MustConnect().MustIncognito().NoDefaultDevice().MustPage("https://www.finewineandgoodspirits.com/")
	page := r.MustPage("https://www.finewineandgoodspirits.com/")
	page.MustElement("#root > header > section > div.header-desktop-wrapper > div.modal.fade.in.modal--active > div > div > div > div > div.age-gate__cta > button").MustClick()
	page.Mouse.Scroll(0, 500, 20)
	time.Sleep(3 * time.Second)

	// utils.Pause()
	var httpCookies []*http.Cookie

	str := "1969-12-31 16:59:59 -0700 MST"
	t, _ := time.Parse(layout, str)

	c := http.Cookie{
		Name:    "AGEVERIFY",
		Value:   "Over21",
		Domain:  "www.finewineandgoodspirits.com",
		Path:    "/",
		Expires: t,
	}

	httpCookies = append(httpCookies, &c)
	cookies, _ := page.Cookies([]string{})
	for _, v := range cookies {
		// fmt.Println(v.Name, v.Value, v.Domain, v.Path, v.Expires.Time())
		c := http.Cookie{
			Name:    v.Name,
			Value:   v.Value,
			Domain:  v.Domain,
			Path:    v.Path,
			Expires: v.Expires.Time(),
		}
		httpCookies = append(httpCookies, &c)
	}

	r.Close()

	return httpCookies
}
