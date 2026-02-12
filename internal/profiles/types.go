package profiles

type Profile struct {
	Site            string          `json:"Site"`
	ProfileName     string          `json:"ProfileName"`
	HotelKeywords   string          `json:"HotelKeywords"`
	RoomKeywords    string          `json:"RoomKeywords"`
	LiquorKeywords  string          `json:"LiquorKeywords"`
	UserInformation UserInformation `json:"UserInformation:"`
}

type UserInformation struct {
	FirstName          string             `json:"FirstName"`
	LastName           string             `json:"LastName"`
	AddressLine1       string             `json:"AddressLine1"`
	AddressLine2       string             `json:"AddressLine2"`
	City               string             `json:"City"`
	State              string             `json:"State"`
	Zip                string             `json:"Zip"`
	Email              string             `json:"Email"`
	PhoneNumber        string             `json:"PhoneNumber"`
	PaymentInformation PaymentInformation `json:"PaymentInformation"`
	BookingInformation BookingInformation `json:"BookingInformation"`
}

type PaymentInformation struct {
	CardNumber string `json:"CardNumber"`
	ExpMon     string `json:"ExpMon"`
	ExpYr      string `json:"ExpYr"`
	CardType   string `json:"CardType"`
	CVV        string `json:"CVV"`
}

type BookingInformation struct {
	CheckInDate  string `json:"CheckInDate"`
	CheckOutDate string `json:"CheckOutDate"`
}
