package pa

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type PostCartJson struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Quantity   int     `json:"quantity"`
	ProductID  string  `json:"productId"`
	CatRefID   string  `json:"catRefId"`
	LocationID *string `json:"locationId,omitempty"`
	// ShippingID *string `json:"shippingGroupId,omitempty"`
}

type CheckItem struct {
	ExternalPriceQuantity int     `json:"externalPriceQuantity"`
	ExternalPrice         float64 `json:"externalPrice"`
	ProductID             string  `json:"productId"`
	Quantity              int     `json:"quantity"`
	CatRefID              string  `json:"catRefId"`
	LocationID            *string `json:"locationId"`
	ShippingGroupID       string  `json:"shippingGroupId"`
}
type CheckCartStruct struct {
	Items []CheckItem `json:"items"`
}

type StockStruct struct {
	Method   string   `json:"method"`
	Location *string  `json:"location"`
	Items    []string `json:"items"`
}

func (e *PaTask) EncodeCartCheck() string {

	p := []CheckItem{{
		ExternalPriceQuantity: -1,
		ExternalPrice:         e.ProductPrice,
		ProductID:             e.ProductID,
		Quantity:              e.QuantityInt,
		CatRefID:              e.ProductID,
		ShippingGroupID:       e.ShippingGroup,
	}}

	c := CheckCartStruct{
		Items: p,
	}

	cartStr, err := json.Marshal(c)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}

	return string(cartStr)

}

func (e *PaTask) EncodeCartJson() string {

	q, err := strconv.Atoi(e.QuantityStr)
	if err != nil {
		q = 1
	}

	e.QuantityInt = q

	p := []CartItem{{

		Quantity:  e.QuantityInt,
		ProductID: e.ProductID,
		CatRefID:  e.ProductID,
		// LocationID: e.LocationID,
		// ShippingID: &e.ShippingGroup,
	},
	}

	c := PostCartJson{
		Items: p,
	}

	cartStr, err := json.Marshal(c)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}

	fmt.Println(string(cartStr))

	return string(cartStr)

}

func (e *PaTask) EncodeStockJson() string {

	var c StockStruct

	if e.Method == Shipping {
		c = StockStruct{
			Method: "b2cShip",
			//Location: ,
			Items: e.MonitorSKUs,
		}
	} else if e.Method == Pickup {
		c = StockStruct{
			Method:   "pickup",
			Location: &e.LocationID,
			Items:    e.MonitorSKUs,
		}
	}
	cartStr, err := json.Marshal(c)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}
	return string(cartStr)

}

func (e *PaTask) EncodeShippingAddress() string {
	var c ShippingStruct

	a := ShippingAddress{
		FirstName:   e.Profile.UserInformation.FirstName,
		LastName:    e.Profile.UserInformation.LastName,
		Address1:    e.Profile.UserInformation.AddressLine1,
		Address2:    e.Profile.UserInformation.AddressLine2,
		City:        e.Profile.UserInformation.City,
		State:       e.Profile.UserInformation.State,
		PostalCode:  e.Profile.UserInformation.Zip,
		Country:     "US",
		FaxNumber:   "standardSg",
		Email:       e.Profile.UserInformation.Email,
		PhoneNumber: e.Profile.UserInformation.PhoneNumber,
	}

	c = ShippingStruct{
		ShippingAddress: a,
	}

	cartStr, err := json.Marshal(c)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}

	return string(cartStr)
}

// {"payments":[{"billingAddress":{"lastName":"GORW","country":"US","address3":"","address2":"","city":"PHILADELPHIA","prefix":"","address1":"123 MAIN ST","jobTitle":"","companyName":"","postalCode":"19127-2108","county":"","suffix":"","firstName":"Bhiwgriuherg","phoneNumber":"3938402900","faxNumber":"","middleName":"","state":"PA","email":"qgwroiuherg@gmail.com","company":null},"cardNumber":"9443550680705555","cardType":"visa","expiryMonth":"09","expiryYear":"29","nameOnCard":"Lenix Woof","customProperties":{"token":"9443550680705555","type":"visa","exp":"0929"},"amount":53.99,"type":"generic"}]}
func (e *PaTask) EncodeBillingInformation() string {

	var pl []Payments

	b := BillingAddress{
		LastName:    e.Profile.UserInformation.LastName,
		Country:     "US",
		Address3:    "",
		Address2:    e.Profile.UserInformation.AddressLine2,
		City:        e.Profile.UserInformation.City,
		Prefix:      "",
		Address1:    e.Profile.UserInformation.AddressLine1,
		JobTitle:    "",
		CompanyName: "",
		PostalCode:  e.Profile.UserInformation.Zip,
		County:      "",
		Suffix:      "",
		FirstName:   e.Profile.UserInformation.FirstName,
		PhoneNumber: e.Profile.UserInformation.PhoneNumber,
		FaxNumber:   "",
		MiddleName:  "",
		State:       "PA",
		Email:       e.Profile.UserInformation.Email,
	}

	exp := e.Profile.UserInformation.PaymentInformation.ExpMon + e.Profile.UserInformation.PaymentInformation.ExpYr[2:]

	c := CustomCardProp{
		Token: e.EncryptedCard,
		Type:  e.Profile.UserInformation.PaymentInformation.CardType,
		Exp:   exp,
	}

	p := Payments{
		BillingAddress:   b,
		CardNumber:       e.EncryptedCard,
		CardType:         e.Profile.UserInformation.PaymentInformation.CardType,
		ExpiryMonth:      e.Profile.UserInformation.PaymentInformation.ExpMon,
		ExpiryYear:       e.Profile.UserInformation.PaymentInformation.ExpYr[2:],
		NameOnCard:       e.Profile.UserInformation.FirstName + " " + e.Profile.UserInformation.LastName,
		CustomProperties: c,
		Amount:           e.OrderTotal,
		Type:             "generic",
	}

	pl = append(pl, p)

	bs := BillingStruct{
		Payments: pl,
	}

	cartStr, err := json.Marshal(bs)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}

	return string(cartStr)

}

func (e *PaTask) EncodeCardStruct() string {
	//{"account":"4444555555555555","source":"iToke","encryptionhandler":null,"unique":false,"expiry":null,"cvv":"998"}

	s := EncodeStruct{
		Account: e.Profile.UserInformation.PaymentInformation.CardNumber,
		Source:  "iToke",
		Unique:  false,
		Cvv:     e.Profile.UserInformation.PaymentInformation.CVV,
	}

	cartStr, err := json.Marshal(s)
	if err != nil {
		log.Fatal("cannot encode cart string", err)
	}
	return string(cartStr)

}

func (t *PaTask) UpdateStatus(stage string) {
	t.Task.Stage = stage
	log.Printf("Task ID: %s | Mode: %s | Profile: %s | Product: %s | Size %s | Status %s\r\n", t.Task.ID, t.Mode, t.Profile.ProfileName, t.ProductName, t.ProductVol, t.Task.Stage)
}

func (t *PaTask) PrintStatus(message string) {
	log.Printf("Task ID: %s | Mode: %s | Profile: %s | Product: %s | Size %s | Status %s\r\n", t.Task.ID, t.Mode, t.Profile.ProfileName, t.ProductName, t.ProductVol, message)
}
