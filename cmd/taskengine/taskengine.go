package taskengine

import (
	proxy "FurAIOIgnited/internal/proxies"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/google/uuid"
)

type IBotTask interface {
	Get() *Task
	Initalize()
	SetProxy()
	Ignite()
}

func (t *Task) Initalize() {

	t.ID = uuid.NewString()

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error making session!")
	}
	t.Client = &http.Client{
		Jar: jar,
	}

}

func (t *Task) SetProxy(proxyList *[]string) {

	proxy, _ := url.Parse(proxy.GetRandProxy(proxyList))

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}

	t.Client.Transport = transport
	t.Proxy = proxy

}
