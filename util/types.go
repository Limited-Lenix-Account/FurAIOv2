package util

type AttendeeOptions struct {
	Title string
	Value string
}

type OldHotelSearchResult []struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	MarketingMessage string `json:"marketingMessage"`
	Blocks           []struct {
		ID               int     `json:"id"`
		Name             string  `json:"name"`
		HotelID          int     `json:"hotelId"`
		RoomTypeID       int     `json:"roomTypeId"`
		GroupTypeID      int     `json:"groupTypeId"`
		Description      any     `json:"description"`
		AverageRate      float64 `json:"averageRate"`
		AverageBasicRate float64 `json:"averageBasicRate"`
		Charge           float64 `json:"charge"`
		TaxesIncluded    bool    `json:"taxesIncluded"`
		MaxPersons       int     `json:"maxPersons"`
		RoomRate2Nd      float64 `json:"roomRate2nd"`
		RoomRate3Rd      float64 `json:"roomRate3rd"`
		RoomRate4Th      float64 `json:"roomRate4th"`
		RoomRate5ThPlus  float64 `json:"roomRate5thPlus"`
		Inventory        []struct {
			Date            []float64 `json:"date"`
			Rate            float64   `json:"rate"`
			SingleRate      float64   `json:"singleRate"`
			HideRate        bool      `json:"hideRate"`
			Available       int       `json:"available"`
			MinLengthOfStay int       `json:"minLengthOfStay"`
			WlAvailable     int       `json:"wlAvailable"`
		} `json:"inventory"`
		Rating             float64 `json:"rating"`
		CancelPolicy       any     `json:"cancelPolicy"`
		TaxPolicy          any     `json:"taxPolicy"`
		UpsellAmount       int     `json:"upsellAmount"`
		MarketingMessage   any     `json:"marketingMessage"`
		ChildrenAffectRate bool    `json:"childrenAffectRate"`
		UserDefinedRank    any     `json:"userDefinedRank"`
		ShowSmokingPref    bool    `json:"showSmokingPref"`
		ImageURL           any     `json:"imageUrl"`
		ImageUrls          any     `json:"imageUrls"`
		Available          bool    `json:"available"`
		MinNumberOfRooms   int     `json:"minNumberOfRooms"`
	} `json:"blocks"`
	MultiPayment      int     `json:"multiPayment"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	DistanceFromEvent float64 `json:"distanceFromEvent"`
	Address           struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		State    string `json:"state"`
		Zip      string `json:"zip"`
		Country  struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Alpha2Code string `json:"alpha2Code"`
		} `json:"country"`
	} `json:"address"`
	PhoneNumber    string `json:"phoneNumber"`
	PhoneNumber2   string `json:"phoneNumber2"`
	TollFreeNumber any    `json:"tollFreeNumber"`
	FaxNumber      string `json:"faxNumber"`
	FaxNumber2     any    `json:"faxNumber2"`
	URL            any    `json:"url"`
	Email          any    `json:"email"`
	Amenities      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"amenities"`
	ChildPolicy       any     `json:"childPolicy"`
	MinAvgRate        float64 `json:"minAvgRate"`
	MaxAvgRate        int     `json:"maxAvgRate"`
	MaxGuests         int     `json:"maxGuests"`
	AccommodationType struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Enable bool   `json:"enable"`
	} `json:"accommodationType"`
	StarRating struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Enable bool   `json:"enable"`
	} `json:"starRating"`
	ShowAccessible  bool `json:"showAccessible"`
	RewardsPrograms []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		IsDefault int    `json:"isDefault"`
		URL       string `json:"url"`
		Text      string `json:"text"`
		ImageURL  string `json:"imageUrl"`
	} `json:"rewardsPrograms"`
	DistanceUnit           int    `json:"distanceUnit"`
	MessageMap             any    `json:"messageMap"`
	Marker                 bool   `json:"marker"`
	ResAccessDate          []int  `json:"resAccessDate"`
	HotelCloseDate         []int  `json:"hotelCloseDate"`
	SplitFolio             int    `json:"splitFolio"`
	ChildrenAffectRate     bool   `json:"childrenAffectRate"`
	FacebookPage           any    `json:"facebookPage"`
	UserDefinedRank        int    `json:"userDefinedRank"`
	SuppressChildrenCount  bool   `json:"suppressChildrenCount"`
	ShowHotelOnClosedEvent bool   `json:"showHotelOnClosedEvent"`
	RewardPlacement        string `json:"rewardPlacement"`
	HotelLevelAddValFlag   int    `json:"hotelLevelAddValFlag"`
	AttendeeAgeHotelInfo   []any  `json:"attendeeAgeHotelInfo"`
	Taxes                  []struct {
		PrimaryInterface             string  `json:"primaryInterface"`
		HotelTaxStructureID          int     `json:"hotelTaxStructureId"`
		HotelTaxName                 string  `json:"hotelTaxName"`
		AmountValue                  float64 `json:"amountValue"`
		HotelTaxAmountTypeID         int     `json:"hotelTaxAmountTypeID"`
		CalculatedFromTypeID         int     `json:"calculatedFromTypeID"`
		CollectionsScheduleTypeID    int     `json:"collectionsScheduleTypeId"`
		CalculatedFromTypeLabel      any     `json:"calculatedFromTypeLabel"`
		HotelTaxAmountTypeLabel      any     `json:"hotelTaxAmountTypeLabel"`
		HotelTaxID                   int     `json:"hotelTaxID"`
		CollectionsScheduleTypeLabel any     `json:"collectionsScheduleTypeLabel"`
	} `json:"taxes"`
	LogoURL            string   `json:"logoUrl"`
	ImageURL           string   `json:"imageUrl"`
	ImageUrls          []string `json:"imageUrls"`
	Charge             any      `json:"charge"`
	TaxAmount          any      `json:"taxAmount"`
	TotalWithTaxes     any      `json:"totalWithTaxes"`
	QuickbookTaxes     any      `json:"quickbookTaxes"`
	CheckInDate        any      `json:"checkInDate"`
	CheckOutDate       any      `json:"checkOutDate"`
	Hqhotel            bool     `json:"hqhotel"`
	HotelChildAgeRange int      `json:"hotelChildAgeRange"`
}

type UpdateStruct struct {
	MultiHotelRoom bool     `json:"multiHotelRoom"`
	HotelID        string   `json:"hotelId"`
	DistanceEnd    int      `json:"distanceEnd"`
	MaxGuests      int      `json:"maxGuests"`
	HotelIds       []string `json:"hotelIds"`
	BlockMap       blockMap `json:"blockMap"`
	MinSlideRate   int      `json:"minSlideRate"`
	MaxSlideRate   int      `json:"maxSlideRate"`
	WlSearch       bool     `json:"wlSearch"`
	ShowAll        bool     `json:"showAll"`
	Mod            bool     `json:"mod"`
}

type blockMap struct {
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

type UpdateResp struct {
	Charge         float64 `json:"charge"`
	TaxAmount      float64 `json:"taxAmount"`
	TotalWithTaxes float64 `json:"totalWithTaxes"`
	QuickbookTaxes []struct {
		ID                   int     `json:"id"`
		Name                 any     `json:"name"`
		Charge               float64 `json:"charge"`
		TaxName              string  `json:"taxName"`
		TaxAmount            float64 `json:"taxAmount"`
		HotelTaxAmountTypeID int     `json:"hotelTaxAmountTypeID"`
	} `json:"quickbookTaxes"`
	BlockInfo []struct {
		BlockID          int     `json:"blockId"`
		Charge           float64 `json:"charge"`
		ActualRate       float64 `json:"actualRate"`
		BlockIDMRT       string  `json:"blockIdMRT"`
		NumberOfRooms    int     `json:"numberOfRooms"`
		NumberOfGuests   int     `json:"numberOfGuests"`
		NumberOfChildren int     `json:"numberOfChildren"`
	} `json:"blockInfo"`
	MultiHotelInfo    any `json:"multiHotelInfo"`
	ProcessingFeeList any `json:"processingFeeList"`
}

type HotelSearchResult []struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	MarketingMessage  any     `json:"marketingMessage"`
	Blocks            Blocks  `json:"blocks"`
	MultiPayment      int     `json:"multiPayment"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	DistanceFromEvent float64 `json:"distanceFromEvent"`
	Address           struct {
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		State    string `json:"state"`
		Zip      string `json:"zip"`
		Country  struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Alpha2Code string `json:"alpha2Code"`
		} `json:"country"`
	} `json:"address"`
	PhoneNumber    string `json:"phoneNumber"`
	PhoneNumber2   string `json:"phoneNumber2"`
	TollFreeNumber string `json:"tollFreeNumber"`
	FaxNumber      string `json:"faxNumber"`
	FaxNumber2     any    `json:"faxNumber2"`
	URL            string `json:"url"`
	Email          any    `json:"email"`
	Amenities      []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"amenities"`
	ChildPolicy       string  `json:"childPolicy"`
	MinAvgRate        float64 `json:"minAvgRate"`
	MaxAvgRate        float64 `json:"maxAvgRate"`
	MaxGuests         int     `json:"maxGuests"`
	AccommodationType struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Enable bool   `json:"enable"`
	} `json:"accommodationType"`
	StarRating struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Enable bool   `json:"enable"`
	} `json:"starRating"`
	ShowAccessible         bool   `json:"showAccessible"`
	RewardsPrograms        []any  `json:"rewardsPrograms"`
	DistanceUnit           int    `json:"distanceUnit"`
	MessageMap             any    `json:"messageMap"`
	Marker                 bool   `json:"marker"`
	ResAccessDate          []int  `json:"resAccessDate"`
	HotelCloseDate         []int  `json:"hotelCloseDate"`
	SplitFolio             int    `json:"splitFolio"`
	ChildrenAffectRate     bool   `json:"childrenAffectRate"`
	FacebookPage           string `json:"facebookPage"`
	UserDefinedRank        int    `json:"userDefinedRank"`
	SuppressChildrenCount  bool   `json:"suppressChildrenCount"`
	ShowHotelOnClosedEvent bool   `json:"showHotelOnClosedEvent"`
	RewardPlacement        string `json:"rewardPlacement"`
	HotelLevelAddValFlag   int    `json:"hotelLevelAddValFlag"`
	AttendeeAgeHotelInfo   []any  `json:"attendeeAgeHotelInfo"`
	Taxes                  []struct {
		PrimaryInterface             string  `json:"primaryInterface"`
		HotelTaxStructureID          int     `json:"hotelTaxStructureId"`
		HotelTaxName                 string  `json:"hotelTaxName"`
		AmountValue                  float64 `json:"amountValue"`
		HotelTaxAmountTypeID         int     `json:"hotelTaxAmountTypeID"`
		CalculatedFromTypeLabel      any     `json:"calculatedFromTypeLabel"`
		HotelTaxAmountTypeLabel      any     `json:"hotelTaxAmountTypeLabel"`
		HotelTaxID                   int     `json:"hotelTaxID"`
		CalculatedFromTypeID         int     `json:"calculatedFromTypeID"`
		CollectionsScheduleTypeLabel any     `json:"collectionsScheduleTypeLabel"`
		CollectionsScheduleTypeID    int     `json:"collectionsScheduleTypeId"`
	} `json:"taxes"`
	LogoURL            string   `json:"logoUrl"`
	ImageURL           string   `json:"imageUrl"`
	ImageUrls          []string `json:"imageUrls"`
	Charge             any      `json:"charge"`
	TaxAmount          any      `json:"taxAmount"`
	TotalWithTaxes     any      `json:"totalWithTaxes"`
	QuickbookTaxes     any      `json:"quickbookTaxes"`
	CheckInDate        any      `json:"checkInDate"`
	CheckOutDate       any      `json:"checkOutDate"`
	Hqhotel            bool     `json:"hqhotel"`
	HotelChildAgeRange int      `json:"hotelChildAgeRange"`
}

type Blocks []struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	HotelID            int       `json:"hotelId"`
	RoomTypeID         int       `json:"roomTypeId"`
	GroupTypeID        int       `json:"groupTypeId"`
	Description        string    `json:"description"`
	AverageRate        float64   `json:"averageRate"`
	AverageBasicRate   float64   `json:"averageBasicRate"`
	Charge             float64   `json:"charge"`
	TaxesIncluded      bool      `json:"taxesIncluded"`
	MaxPersons         int       `json:"maxPersons"`
	RoomRate2Nd        float64   `json:"roomRate2nd"`
	RoomRate3Rd        float64   `json:"roomRate3rd"`
	RoomRate4Th        float64   `json:"roomRate4th"`
	RoomRate5ThPlus    float64   `json:"roomRate5thPlus"`
	Inventory          Inventory `json:"inventory"`
	Rating             float64   `json:"rating"`
	CancelPolicy       string    `json:"cancelPolicy"`
	TaxPolicy          string    `json:"taxPolicy"`
	UpsellAmount       float64   `json:"upsellAmount"`
	MarketingMessage   any       `json:"marketingMessage"`
	ChildrenAffectRate bool      `json:"childrenAffectRate"`
	UserDefinedRank    int       `json:"userDefinedRank"`
	ShowSmokingPref    bool      `json:"showSmokingPref"`
	ImageURL           string    `json:"imageUrl"`
	ImageUrls          []string  `json:"imageUrls"`
	Available          bool      `json:"available"`
	MinNumberOfRooms   int       `json:"minNumberOfRooms"`
}

type Inventory []struct {
	Date            []int   `json:"date"`
	Rate            float64 `json:"rate"`
	SingleRate      float64 `json:"singleRate"`
	HideRate        bool    `json:"hideRate"`
	Available       int     `json:"available"`
	MinLengthOfStay int     `json:"minLengthOfStay"`
	WlAvailable     int     `json:"wlAvailable"`
}

type PasskeySuccessfulOrder struct {
	MasterAck         any     `json:"masterAck"`
	Charge            float64 `json:"charge"`
	TaxAmount         float64 `json:"taxAmount"`
	TotalWithTaxes    float64 `json:"totalWithTaxes"`
	Nor1Upsell        int     `json:"nor1Upsell"`
	GuestWithSameName bool    `json:"guestWithSameName"`
	Reservations      []struct {
		ID                     int    `json:"id"`
		Name                   any    `json:"name"`
		EventID                int    `json:"eventId"`
		HotelID                int    `json:"hotelId"`
		BlockID                int    `json:"blockId"`
		GroupTypeID            int    `json:"groupTypeId"`
		StatusID               int    `json:"statusId"`
		AccessibilityRequested bool   `json:"accessibilityRequested"`
		SpecialRequests        string `json:"specialRequests"`
		OptIn                  bool   `json:"optIn"`
		CheckInDate            []int  `json:"checkInDate"`
		CheckOutDate           []int  `json:"checkOutDate"`
		AckNumber              string `json:"ackNumber"`
		MasterAckNumber        any    `json:"masterAckNumber"`
		Guests                 []struct {
			ID              int64  `json:"id"`
			Name            any    `json:"name"`
			Prefix          any    `json:"prefix"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			MiddleName      any    `json:"middleName"`
			Suffix          any    `json:"suffix"`
			SpecialRequests string `json:"specialRequests"`
			PhoneNumber     string `json:"phoneNumber"`
			Email           string `json:"email"`
			ConfirmEmail    string `json:"confirmEmail"`
			PrevEmail       any    `json:"prevEmail"`
			Organization    any    `json:"organization"`
			Position        any    `json:"position"`
			Address         struct {
				Address1 string `json:"address1"`
				Address2 string `json:"address2"`
				City     string `json:"city"`
				State    string `json:"state"`
				Zip      string `json:"zip"`
				Country  struct {
					ID         int    `json:"id"`
					Name       any    `json:"name"`
					Alpha2Code string `json:"alpha2Code"`
				} `json:"country"`
			} `json:"address"`
			BillingInfo struct {
				ID      int `json:"id"`
				Name    any `json:"name"`
				Payment struct {
					CreditCard struct {
						ID          int    `json:"id"`
						Name        string `json:"name"`
						CardNumber  string `json:"cardNumber"`
						ExpiryMonth int    `json:"expiryMonth"`
						ExpiryYear  int    `json:"expiryYear"`
						CardTypeID  int    `json:"cardTypeId"`
						DepDate     any    `json:"depDate"`
						ValidMasked bool   `json:"validMasked"`
						Cvvcode     any    `json:"cvvcode"`
					} `json:"creditCard"`
					PaymentType  string `json:"paymentType"`
					OtherPayment any    `json:"otherPayment"`
				} `json:"payment"`
				Payer struct {
					HolderName  string `json:"holderName"`
					PhoneNumber string `json:"phoneNumber"`
					Address     struct {
						Address1 string `json:"address1"`
						Address2 any    `json:"address2"`
						City     string `json:"city"`
						State    string `json:"state"`
						Zip      string `json:"zip"`
						Country  struct {
							ID         int    `json:"id"`
							Name       any    `json:"name"`
							Alpha2Code string `json:"alpha2Code"`
						} `json:"country"`
					} `json:"address"`
				} `json:"payer"`
			} `json:"billingInfo"`
			TravelInfo struct {
				ArrDateTime any `json:"arrDateTime"`
				DepDateTime any `json:"depDateTime"`
				ArrAirline  struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"arrAirline"`
				DepAirline struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"depAirline"`
				ArrFlightNum   any    `json:"arrFlightNum"`
				DepFlightNum   any    `json:"depFlightNum"`
				ArrDate        any    `json:"arrDate"`
				ArrTime        any    `json:"arrTime"`
				DepDate        any    `json:"depDate"`
				DepTime        any    `json:"depTime"`
				OtherDetails   any    `json:"otherDetails"`
				Transportation string `json:"transportation"`
			} `json:"travelInfo"`
			Primary           bool   `json:"primary"`
			CopyContact       bool   `json:"copyContact"`
			ArrDateTime       []int  `json:"arrDateTime"`
			DepDateTime       []int  `json:"depDateTime"`
			ArrDate           []int  `json:"arrDate"`
			ArrTime           []int  `json:"arrTime"`
			DepDate           []int  `json:"depDate"`
			DepTime           []int  `json:"depTime"`
			AckSuffix         string `json:"ackSuffix"`
			FamilyName        any    `json:"familyName"`
			GivenName         any    `json:"givenName"`
			LocaleID          string `json:"localeId"`
			PartialNameString string `json:"partialNameString"`
			FullNameString    string `json:"fullNameString"`
		} `json:"guests"`
		InfoList []struct {
			Date             []int   `json:"date"`
			Rate             float64 `json:"rate"`
			RateWithFee      float64 `json:"rateWithFee"`
			NumberOfGuests   int     `json:"numberOfGuests"`
			HideRate         bool    `json:"hideRate"`
			NumberOfChildren int     `json:"numberOfChildren"`
			Confirmed        bool    `json:"confirmed"`
		} `json:"infoList"`
		NumberOfGuests      int     `json:"numberOfGuests"`
		NumberOfChildren    int     `json:"numberOfChildren"`
		Charge              float64 `json:"charge"`
		TaxAmount           float64 `json:"taxAmount"`
		Subtotal            float64 `json:"subtotal"`
		UpsellAmount        float64 `json:"upsellAmount"`
		RewardsProgram      any     `json:"rewardsProgram"`
		RewardMemberAccount any     `json:"rewardMemberAccount"`
		SmokingPreference   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"smokingPreference"`
		Addons                []any   `json:"addons"`
		SuppressedOnAWSAddons []any   `json:"suppressedOnAWSAddons"`
		PrevBlockID           int     `json:"prevBlockId"`
		SplitFolio            int     `json:"splitFolio"`
		ChildrenAffectRate    bool    `json:"childrenAffectRate"`
		BridgeID              any     `json:"bridgeId"`
		Mobile                bool    `json:"mobile"`
		CustomField1          string  `json:"customField1"`
		CustomField2          string  `json:"customField2"`
		CustomField3          string  `json:"customField3"`
		CustomField4          string  `json:"customField4"`
		CustomField5          string  `json:"customField5"`
		CustomField6          string  `json:"customField6"`
		ChildProfiles         []any   `json:"childProfiles"`
		Cancelled             bool    `json:"cancelled"`
		UsdCharge             float64 `json:"usdCharge"`
		ActualRate            float64 `json:"actualRate"`
		UsdActualRate         float64 `json:"usdActualRate"`
		UsdTaxAmount          float64 `json:"usdTaxAmount"`
		UsdSubtotal           float64 `json:"usdSubtotal"`
		Roc                   float64 `json:"roc"`
		Waitlisted            bool    `json:"waitlisted"`
		CancelledWaitlisted   bool    `json:"cancelledWaitlisted"`
		NewGuestAckSuffix     string  `json:"newGuestAckSuffix"`
		PrimaryGuest          struct {
			ID              int64  `json:"id"`
			Name            any    `json:"name"`
			Prefix          any    `json:"prefix"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			MiddleName      any    `json:"middleName"`
			Suffix          any    `json:"suffix"`
			SpecialRequests string `json:"specialRequests"`
			PhoneNumber     string `json:"phoneNumber"`
			Email           string `json:"email"`
			ConfirmEmail    string `json:"confirmEmail"`
			PrevEmail       any    `json:"prevEmail"`
			Organization    any    `json:"organization"`
			Position        any    `json:"position"`
			Address         struct {
				Address1 string `json:"address1"`
				Address2 string `json:"address2"`
				City     string `json:"city"`
				State    string `json:"state"`
				Zip      string `json:"zip"`
				Country  struct {
					ID         int    `json:"id"`
					Name       any    `json:"name"`
					Alpha2Code string `json:"alpha2Code"`
				} `json:"country"`
			} `json:"address"`
			BillingInfo struct {
				ID      int `json:"id"`
				Name    any `json:"name"`
				Payment struct {
					CreditCard struct {
						ID          int    `json:"id"`
						Name        string `json:"name"`
						CardNumber  string `json:"cardNumber"`
						ExpiryMonth int    `json:"expiryMonth"`
						ExpiryYear  int    `json:"expiryYear"`
						CardTypeID  int    `json:"cardTypeId"`
						DepDate     any    `json:"depDate"`
						ValidMasked bool   `json:"validMasked"`
						Cvvcode     any    `json:"cvvcode"`
					} `json:"creditCard"`
					PaymentType  string `json:"paymentType"`
					OtherPayment any    `json:"otherPayment"`
				} `json:"payment"`
				Payer struct {
					HolderName  string `json:"holderName"`
					PhoneNumber string `json:"phoneNumber"`
					Address     struct {
						Address1 string `json:"address1"`
						Address2 any    `json:"address2"`
						City     string `json:"city"`
						State    string `json:"state"`
						Zip      string `json:"zip"`
						Country  struct {
							ID         int    `json:"id"`
							Name       any    `json:"name"`
							Alpha2Code string `json:"alpha2Code"`
						} `json:"country"`
					} `json:"address"`
				} `json:"payer"`
			} `json:"billingInfo"`
			TravelInfo struct {
				ArrDateTime any `json:"arrDateTime"`
				DepDateTime any `json:"depDateTime"`
				ArrAirline  struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"arrAirline"`
				DepAirline struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"depAirline"`
				ArrFlightNum   any    `json:"arrFlightNum"`
				DepFlightNum   any    `json:"depFlightNum"`
				ArrDate        any    `json:"arrDate"`
				ArrTime        any    `json:"arrTime"`
				DepDate        any    `json:"depDate"`
				DepTime        any    `json:"depTime"`
				OtherDetails   any    `json:"otherDetails"`
				Transportation string `json:"transportation"`
			} `json:"travelInfo"`
			Primary           bool   `json:"primary"`
			CopyContact       bool   `json:"copyContact"`
			ArrDateTime       []int  `json:"arrDateTime"`
			DepDateTime       []int  `json:"depDateTime"`
			ArrDate           []int  `json:"arrDate"`
			ArrTime           []int  `json:"arrTime"`
			DepDate           []int  `json:"depDate"`
			DepTime           []int  `json:"depTime"`
			AckSuffix         string `json:"ackSuffix"`
			FamilyName        any    `json:"familyName"`
			GivenName         any    `json:"givenName"`
			LocaleID          string `json:"localeId"`
			PartialNameString string `json:"partialNameString"`
			FullNameString    string `json:"fullNameString"`
		} `json:"primaryGuest"`
		HideRates          bool `json:"hideRates"`
		SuppressRates      bool `json:"suppressRates"`
		ChildNamesProvided bool `json:"childNamesProvided"`
		TotalWLNights      int  `json:"totalWLNights"`
		NumberOfNights     int  `json:"numberOfNights"`
	} `json:"reservations"`
	GroupBooking              any    `json:"groupBooking"`
	SingleResOriginMode       string `json:"singleResOriginMode"`
	AllowChangeHotel          bool   `json:"allowChangeHotel"`
	AttendeeConsentAgreements []struct {
		ID                  int    `json:"id"`
		Name                any    `json:"name"`
		AttendeeID          int    `json:"attendeeId"`
		BusinessTextID      string `json:"businessTextId"`
		LocaleID            string `json:"localeId"`
		ConsentGiven        int    `json:"consentGiven"`
		MessageCreationDate int64  `json:"messageCreationDate"`
		DateOfConsent       any    `json:"dateOfConsent"`
		Required            bool   `json:"required"`
		ConsentType         string `json:"consentType"`
		ConsentLabel        string `json:"consentLabel"`
		ConsentMessage      string `json:"consentMessage"`
	} `json:"attendeeConsentAgreements"`
	Taxes []struct {
		ID                   int     `json:"id"`
		Name                 any     `json:"name"`
		Charge               float64 `json:"charge"`
		TaxName              string  `json:"taxName"`
		TaxAmount            float64 `json:"taxAmount"`
		HotelTaxAmountTypeID int     `json:"hotelTaxAmountTypeID"`
	} `json:"taxes"`
	ThirdPartyVerificationURL any `json:"thirdPartyVerificationUrl"`
	ReservationDates          struct {
		CheckIn  []int `json:"checkIn"`
		CheckOut []int `json:"checkOut"`
	} `json:"reservationDates"`
	GroupTypeID      int `json:"groupTypeId"`
	TotalBlocksCount struct {
		Num939057787 int `json:"939057787"`
	} `json:"totalBlocksCount"`
	Dates     [][]int `json:"dates"`
	AllGuests struct {
		Num19438448740 struct {
			ID              int64  `json:"id"`
			Name            any    `json:"name"`
			Prefix          any    `json:"prefix"`
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			MiddleName      any    `json:"middleName"`
			Suffix          any    `json:"suffix"`
			SpecialRequests string `json:"specialRequests"`
			PhoneNumber     string `json:"phoneNumber"`
			Email           string `json:"email"`
			ConfirmEmail    string `json:"confirmEmail"`
			PrevEmail       any    `json:"prevEmail"`
			Organization    any    `json:"organization"`
			Position        any    `json:"position"`
			Address         struct {
				Address1 string `json:"address1"`
				Address2 string `json:"address2"`
				City     string `json:"city"`
				State    string `json:"state"`
				Zip      string `json:"zip"`
				Country  struct {
					ID         int    `json:"id"`
					Name       any    `json:"name"`
					Alpha2Code string `json:"alpha2Code"`
				} `json:"country"`
			} `json:"address"`
			BillingInfo struct {
				ID      int `json:"id"`
				Name    any `json:"name"`
				Payment struct {
					CreditCard struct {
						ID          int    `json:"id"`
						Name        string `json:"name"`
						CardNumber  string `json:"cardNumber"`
						ExpiryMonth int    `json:"expiryMonth"`
						ExpiryYear  int    `json:"expiryYear"`
						CardTypeID  int    `json:"cardTypeId"`
						DepDate     any    `json:"depDate"`
						ValidMasked bool   `json:"validMasked"`
						Cvvcode     any    `json:"cvvcode"`
					} `json:"creditCard"`
					PaymentType  string `json:"paymentType"`
					OtherPayment any    `json:"otherPayment"`
				} `json:"payment"`
				Payer struct {
					HolderName  string `json:"holderName"`
					PhoneNumber string `json:"phoneNumber"`
					Address     struct {
						Address1 string `json:"address1"`
						Address2 any    `json:"address2"`
						City     string `json:"city"`
						State    string `json:"state"`
						Zip      string `json:"zip"`
						Country  struct {
							ID         int    `json:"id"`
							Name       any    `json:"name"`
							Alpha2Code string `json:"alpha2Code"`
						} `json:"country"`
					} `json:"address"`
				} `json:"payer"`
			} `json:"billingInfo"`
			TravelInfo struct {
				ArrDateTime any `json:"arrDateTime"`
				DepDateTime any `json:"depDateTime"`
				ArrAirline  struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"arrAirline"`
				DepAirline struct {
					ID   int `json:"id"`
					Name any `json:"name"`
				} `json:"depAirline"`
				ArrFlightNum   any    `json:"arrFlightNum"`
				DepFlightNum   any    `json:"depFlightNum"`
				ArrDate        any    `json:"arrDate"`
				ArrTime        any    `json:"arrTime"`
				DepDate        any    `json:"depDate"`
				DepTime        any    `json:"depTime"`
				OtherDetails   any    `json:"otherDetails"`
				Transportation string `json:"transportation"`
			} `json:"travelInfo"`
			Primary           bool   `json:"primary"`
			CopyContact       bool   `json:"copyContact"`
			ArrDateTime       []int  `json:"arrDateTime"`
			DepDateTime       []int  `json:"depDateTime"`
			ArrDate           []int  `json:"arrDate"`
			ArrTime           []int  `json:"arrTime"`
			DepDate           []int  `json:"depDate"`
			DepTime           []int  `json:"depTime"`
			AckSuffix         string `json:"ackSuffix"`
			FamilyName        any    `json:"familyName"`
			GivenName         any    `json:"givenName"`
			LocaleID          string `json:"localeId"`
			PartialNameString string `json:"partialNameString"`
			FullNameString    string `json:"fullNameString"`
		} `json:"19438448740"`
	} `json:"allGuests"`
	MaxGuestCount int   `json:"maxGuestCount"`
	TotalWLNights int   `json:"totalWLNights"`
	BlockIds      []int `json:"blockIds"`
	BlockDates    struct {
		Num939057787 [][]int `json:"939057787"`
	} `json:"blockDates"`
	BlocksCount struct {
		Num939057787 int `json:"939057787"`
	} `json:"blocksCount"`
	EditMode       bool    `json:"editMode"`
	Wlblocks       []any   `json:"wlblocks"`
	MaxChildCount  int     `json:"maxChildCount"`
	IntersectDates [][]int `json:"intersectDates"`
}
