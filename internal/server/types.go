package server

type SkuData []struct {
	Route          string `json:"route"`
	CustomStock    string `json:"customStock"`
	ProtectedStock string `json:"protectedStock"`
	Available      bool   `json:"available"`
	Sku            string `json:"sku"`
	ProdName       string `json:"prodName"`
	ProdVolume     string `json:"prodVol"`
	MaxQuantity    string `json:"maxQuantity"`
}
