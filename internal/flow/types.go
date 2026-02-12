package flow

import (
	"FurAIOIgnited/util"
	"net/http"
	"time"
)

type UserTask struct {
	Hotel        PasskeyID
	Client       HttpClient
	AttendeeInfo util.AttendeeOptions
	UserHotel    UserHotel
	User         UserInfo

	Personal PersonalInfo

	Checkout Checkout
}

type HttpClient struct {
	client *http.Client
	CSRF   string
}

type PasskeyID struct {
	OwnerID string
	EventID string
	GroupID string

	StemURL string
}

type UserHotel struct {
	HotelID int
	BlockID int
	RoomID  string

	Charge   string
	Subtotal string
}

type UserInfo struct {
	CheckIn        string
	CheckOut       string
	NumberOfGuests int
	NumberOfRooms  int
}

type UpdateStruct struct {
	HotelID  int      `json:"hotelId"`
	BlockMap BlockMap `json:"blockMap"`
}

type BlockMap struct {
	Blocks      []Block `json:"blocks"`
	TotalRooms  int     `json:"totalRooms"`
	TotalGuests int     `json:"totalGuests"`
}

type Block struct {
	HotelID          int `json:"hotelId"`
	BlockID          int `json:"blockId"`
	CheckIn          any `json:"checkIn"`
	CheckOut         any `json:"checkOut"`
	NumberOfGuests   int `json:"numberOfGuests"`
	NumberOfRooms    int `json:"numberOfRooms"`
	NumberOfChildren int `json:"numberOfChildren"`
}

type PersonalInfo struct {
	FirstName string
	LastName  string

	Address1 string
	Address2 string
	City     string
	State    string
	Zip      string

	PhoneNumber  string
	EmailAddress string

	Payment PaymentInfo
}

type PaymentInfo struct {
	CreditNumber string
	ExpMonth     string
	ExpYear      string
}

type Checkout struct {
	CheckoutDateIn  string
	CheckoutDateOut string
}

type QueueItChallenge struct {
	Key              string `json:"key"`
	ImageBase64      string `json:"imageBase64"`
	SoundBase64      string `json:"soundBase64"`
	Meta             string `json:"meta"`
	SessionID        string `json:"sessionId"`
	ChallengeDetails string `json:"challengeDetails"`
}

type QueueItVerify struct {
	ChallengeType    string `json:"challengeType"`
	SessionID        string `json:"sessionId"`
	ChallengeDetails string `json:"challengeDetails"`
	Solution         string `json:"solution"`
	Stats            Stats  `json:"stats"`
	CustomerID       string `json:"customerId"`
	EventID          string `json:"eventId"`
	Version          int    `json:"version"`
}

type Stats struct {
	UserAgent      string `json:"userAgent"`
	Screen         string `json:"screen"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browserVersion"`
	IsMobile       bool   `json:"isMobile"`
	Os             string `json:"os"`
	OsVersion      string `json:"osVersion"`
	CookiesEnabled bool   `json:"cookiesEnabled"`
	Tries          int    `json:"tries"`
	Duration       int    `json:"duration"`
}

type InQueue struct {
	IsVerified  bool            `json:"isVerified"`
	Timestamp   time.Time       `json:"timestamp"`
	SessionInfo VerifiedSession `json:"sessionInfo"`
}

type VerifiedSession struct {
	SessionID     string    `json:"sessionId"`
	Timestamp     time.Time `json:"timestamp"`
	Checksum      string    `json:"checksum"`
	SourceIP      string    `json:"sourceIp"`
	ChallengeType string    `json:"challengeType"`
	Version       int       `json:"version"`
	CustomerID    string    `json:"customerId"`
	WaitingRoomID string    `json:"waitingRoomId"`
}

type QueryJson struct {
	SessionID     string    `json:"sessionId"`
	Timestamp     time.Time `json:"timestamp"`
	Checksum      string    `json:"checksum"`
	SourceIP      string    `json:"sourceIp"`
	ChallengeType string    `json:"challengeType"`
	Version       int       `json:"version"`
	CustomerID    string    `json:"customerId"`
	WaitingRoomID string    `json:"waitingRoomId"`
}

type PostStatus struct {
	TargetURL               string `json:"targetUrl"`
	CustomURLParams         string `json:"customUrlParams"`
	LayoutVersion           int64  `json:"layoutVersion"`
	LayoutName              string `json:"layoutName"`
	IsClientRedayToRedirect bool   `json:"isClientRedayToRedirect"`
	IsBeforeOrIdle          bool   `json:"isBeforeOrIdle"`
}
