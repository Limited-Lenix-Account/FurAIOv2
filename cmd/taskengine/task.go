package taskengine

import (
	"FurAIOIgnited/internal/profiles"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Task struct {
	Websocket *websocket.Conn
	Stage     string
	Proxy     *url.URL
	ProxyList []string
	Mode      string
	ID        string
	StartTime time.Time
	Log       *logrus.Logger

	// Flow    string
	Profile *profiles.Profile
	Client  *http.Client
}
