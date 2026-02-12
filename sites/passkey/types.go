package passkey

const (
	Start             = "Start"
	GetHotels         = "GetHotels"
	Stop              = "Stop"
	Update            = "Update"
	GetRooms          = "GetRooms"
	MakeBlock         = "MakeBlock"
	SubmitInfo        = "SubmitInfo"
	SubmitTravel      = "SubmitTravel"
	SubmitPayment     = "SubmitPayment"
	SubmitReservation = "SubmitReservation"
	CheckOrder        = "CheckOrder"
	Successful        = "Successful"
	Challenge         = "GetChallenge"
	PowChallenge      = "PowChallenge"
	StartPow          = "StartPow"
	SubmitSolution    = "SubmitSolution"
	CheckQueue        = "CheckQueue"
)

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
