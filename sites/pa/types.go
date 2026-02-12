package pa

import "time"

const (
	Start              = "Start"
	CreateBrowser      = "Creating Browser"
	SetCookies         = "Setting Cookies"
	CreateCart         = "Creating Cart"
	AddToCart          = "Add To Cart"
	CheckStock         = "Check Stock"
	MakeSession        = "Make Session"
	OutOfStock         = "Out Of Stock!"
	ProductNotLoaded   = "Product Not Loaded!"
	Retrying           = "Retrying..."
	SubmitShipping     = "Submit Shipping"
	GetShipRates       = "Get Ship Rates"
	SubmitShippingRate = "Submit Shipping Rate"
	EncryptCard        = "Encrypt Card"
	SubmitOrder        = "Submit Order"
	Stop               = "Stop"
	Login              = "Login"
	Monitor            = "Monitor"
	Listen             = "Listen"
	Shipping           = "Shipping"
	Pickup             = "Pickup"
	AddPreloadItem     = "Add Preload Item"
	CheckOrder         = "Check Order"
	RemoveItem         = "Remove Item"
	CheckItems         = "Check Items"
	PaymentDeclined    = "Payment Declined"
	Success            = "Success"
)

type PaError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Errors    []struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
		MoreInfo  string `json:"moreInfo"`
		Status    string `json:"status"`
	} `json:"errors"`
	Status string `json:"status"`
}

type AddToCartResp struct {
	TotalResults string `json:"totalResults"`
	Offset       string `json:"offset"`
	HasMore      string `json:"hasMore"`
	Limit        string `json:"limit"`
	Links        []struct {
		Method string `json:"method"`
		Rel    string `json:"rel"`
		Href   string `json:"href"`
	} `json:"links"`
	Embedded struct {
		Order struct {
			ShippingGroups []struct {
				TaxPriceInfo struct {
					CityTax                    float64 `json:"cityTax"`
					SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
					Amount                     float64 `json:"amount"`
					ValueAddedTax              float64 `json:"valueAddedTax"`
					CountyTax                  float64 `json:"countyTax"`
					IsTaxIncluded              bool    `json:"isTaxIncluded"`
					MiscTax                    float64 `json:"miscTax"`
					DistrictTax                float64 `json:"districtTax"`
					StateTax                   float64 `json:"stateTax"`
					CountryTax                 float64 `json:"countryTax"`
				} `json:"taxPriceInfo"`
				PriceInfo struct {
					Amount                  float64 `json:"amount"`
					Total                   float64 `json:"total"`
					LkpValExcludingFreeShip any     `json:"lkpValExcludingFreeShip"`
					Shipping                float64 `json:"shipping"`
					ShippingSurchargeValue  float64 `json:"shippingSurchargeValue"`
					Tax                     float64 `json:"tax"`
					SubTotal                float64 `json:"subTotal"`
					CurrencyCode            string  `json:"currencyCode"`
					TotalWithoutTax         float64 `json:"totalWithoutTax"`
				} `json:"priceInfo"`
				DiscountInfo struct {
					OrderDiscount    float64 `json:"orderDiscount"`
					DiscountDescList []any   `json:"discountDescList"`
					ShippingDiscount float64 `json:"shippingDiscount"`
				} `json:"discountInfo"`
				ShippingMethod struct {
					SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
					ShippingTax                float64 `json:"shippingTax"`
					Cost                       float64 `json:"cost"`
					TaxIncluded                bool    `json:"taxIncluded"`
					ExternalID                 any     `json:"externalId"`
					TaxCode                    string  `json:"taxCode"`
					Value                      string  `json:"value"`
					ShippingMethodDescription  string  `json:"shippingMethodDescription"`
				} `json:"shippingMethod"`
				ShippingGroupID string `json:"shippingGroupId"`
				ShippingAddress struct {
					LastName    any `json:"lastName"`
					Country     any `json:"country"`
					Address3    any `json:"address3"`
					Address2    any `json:"address2"`
					City        any `json:"city"`
					Prefix      any `json:"prefix"`
					Address1    any `json:"address1"`
					PostalCode  any `json:"postalCode"`
					CompanyName any `json:"companyName"`
					JobTitle    any `json:"jobTitle"`
					County      any `json:"county"`
					Suffix      any `json:"suffix"`
					FirstName   any `json:"firstName"`
					PhoneNumber any `json:"phoneNumber"`
					FaxNumber   any `json:"faxNumber"`
					Alias       any `json:"alias"`
					MiddleName  any `json:"middleName"`
					State       any `json:"state"`
					Email       any `json:"email"`
				} `json:"shippingAddress,omitempty"`
				Type     string `json:"type"`
				Items    []any  `json:"items"`
				LastName any    `json:"lastName,omitempty"`
				Store    struct {
					Country      string `json:"country"`
					Hours        any    `json:"hours"`
					Address3     any    `json:"address3"`
					Address2     any    `json:"address2"`
					City         string `json:"city"`
					Address1     string `json:"address1"`
					PostalCode   string `json:"postalCode"`
					County       any    `json:"county"`
					StateAddress string `json:"stateAddress"`
					PhoneNumber  any    `json:"phoneNumber"`
					LocationID   string `json:"locationId"`
					Name         string `json:"name"`
					FaxNumber    any    `json:"faxNumber"`
					Email        any    `json:"email"`
				} `json:"store,omitempty"`
				FirstName   any `json:"firstName,omitempty"`
				PhoneNumber any `json:"phoneNumber,omitempty"`
				MiddleName  any `json:"middleName,omitempty"`
				Email       any `json:"email,omitempty"`
			} `json:"shippingGroups"`
			CreationSiteID         string `json:"creationSiteId"`
			OrderID                string `json:"orderId"`
			AllowAlternateCurrency bool   `json:"allowAlternateCurrency"`
			DynamicProperties      []struct {
				Value any    `json:"value"`
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"dynamicProperties"`
			Payments       []any `json:"payments"`
			PriceListGroup struct {
				IsTaxIncluded bool   `json:"isTaxIncluded"`
				EndDate       any    `json:"endDate"`
				DisplayName   string `json:"displayName"`
				ListPriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"listPriceList"`
				Active                     bool   `json:"active"`
				IsPointsBased              bool   `json:"isPointsBased"`
				Locale                     string `json:"locale"`
				ShippingSurchargePriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"shippingSurchargePriceList"`
				Deleted            bool   `json:"deleted"`
				TaxCalculationType string `json:"taxCalculationType"`
				RepositoryID       string `json:"repositoryId"`
				SalePriceList      struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"salePriceList"`
				Currency struct {
					CurrencyType     any    `json:"currencyType"`
					Symbol           string `json:"symbol"`
					Deleted          bool   `json:"deleted"`
					DisplayName      string `json:"displayName"`
					RepositoryID     string `json:"repositoryId"`
					FractionalDigits int    `json:"fractionalDigits"`
					CurrencyCode     string `json:"currencyCode"`
					NumericCode      string `json:"numericCode"`
				} `json:"currency"`
				ID                 string `json:"id"`
				IncludeAllProducts bool   `json:"includeAllProducts"`
				StartDate          any    `json:"startDate"`
			} `json:"priceListGroup"`
			OrderAction string `json:"orderAction"`
			PriceInfo   struct {
				Amount                 float64 `json:"amount"`
				Total                  float64 `json:"total"`
				Shipping               float64 `json:"shipping"`
				ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
				Tax                    float64 `json:"tax"`
				SubTotal               float64 `json:"subTotal"`
				CurrencyCode           string  `json:"currencyCode"`
				TotalWithoutTax        float64 `json:"totalWithoutTax"`
			} `json:"priceInfo"`
			DiscountInfo struct {
				UnclaimedCouponMultiPromotions struct {
				} `json:"unclaimedCouponMultiPromotions"`
				OrderCouponsMap struct {
				} `json:"orderCouponsMap"`
				OrderDiscount             float64 `json:"orderDiscount"`
				ShippingDiscount          float64 `json:"shippingDiscount"`
				OrderImplicitDiscountList []any   `json:"orderImplicitDiscountList"`
				UnclaimedCouponsMap       struct {
				} `json:"unclaimedCouponsMap"`
				ClaimedCouponMultiPromotions struct {
				} `json:"claimedCouponMultiPromotions"`
			} `json:"discountInfo"`
			ShoppingCart struct {
				Items []struct {
					PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
					RawTotalPrice        float64 `json:"rawTotalPrice"`
					DisplayName          string  `json:"displayName"`
					DynamicProperties    []struct {
						Value any    `json:"value"`
						ID    string `json:"id"`
						Label string `json:"label"`
					} `json:"dynamicProperties"`
					ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
					DiscountAmount         float64 `json:"discountAmount"`
					ExternalData           []any   `json:"externalData"`
					Description            any     `json:"description"`
					IsItemValid            bool    `json:"isItemValid"`
					ItemDiscountInfos      []any   `json:"itemDiscountInfos"`
					CommerceItemID         string  `json:"commerceItemId"`
					Price                  float64 `json:"price"`
					Variant                []struct {
						OptionName  string `json:"optionName"`
						OptionValue string `json:"optionValue"`
					} `json:"variant"`
					PrimaryImageAltText string  `json:"primaryImageAltText"`
					OnSale              bool    `json:"onSale"`
					ID                  string  `json:"id"`
					State               string  `json:"state"`
					StateKey            string  `json:"stateKey"`
					UnitPrice           float64 `json:"unitPrice"`
					PrimaryImageTitle   string  `json:"primaryImageTitle"`
					ChildSKUs           []struct {
						PrimaryThumbImageURL any `json:"primaryThumbImageURL"`
					} `json:"childSKUs"`
					Amount                float64 `json:"amount"`
					Quantity              int     `json:"quantity"`
					ProductID             string  `json:"productId"`
					PointOfNoRevision     bool    `json:"pointOfNoRevision"`
					SalePrice             float64 `json:"salePrice"`
					OrderDiscountInfos    []any   `json:"orderDiscountInfos"`
					DetailedItemPriceInfo []struct {
						Discounted                 bool    `json:"discounted"`
						SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
						Amount                     float64 `json:"amount"`
						Quantity                   int     `json:"quantity"`
						ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
						Tax                        float64 `json:"tax"`
						OrderDiscountShare         float64 `json:"orderDiscountShare"`
						DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
						CurrencyCode               string  `json:"currencyCode"`
					} `json:"detailedItemPriceInfo"`
					GiftWithPurchaseCommerceItemMarkers []any  `json:"giftWithPurchaseCommerceItemMarkers"`
					OriginalCommerceItemID              any    `json:"originalCommerceItemId"`
					TaxCode                             any    `json:"taxCode"`
					CatRefID                            string `json:"catRefId"`
					SkuProperties                       []struct {
						Name         string `json:"name"`
						ID           string `json:"id"`
						Value        any    `json:"value"`
						PropertyType string `json:"propertyType"`
					} `json:"skuProperties"`
					Route        string `json:"route"`
					DiscountInfo []any  `json:"discountInfo"`
					SiteID       string `json:"siteId"`
					ShopperInput struct {
					} `json:"shopperInput"`
					Asset     bool    `json:"asset"`
					ListPrice float64 `json:"listPrice"`
				} `json:"items"`
				NumberOfItems int `json:"numberOfItems"`
			} `json:"shoppingCart"`
			GiftWithPurchaseInfo         []any  `json:"giftWithPurchaseInfo"`
			SiteID                       string `json:"siteId"`
			Markers                      []any  `json:"markers"`
			GiftWithPurchaseOrderMarkers []any  `json:"giftWithPurchaseOrderMarkers"`
		} `json:"order"`
	} `json:"embedded"`
	Items []struct {
		PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
		RawTotalPrice        float64 `json:"rawTotalPrice"`
		DisplayName          string  `json:"displayName"`
		DynamicProperties    []struct {
			Value any    `json:"value"`
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"dynamicProperties"`
		ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
		DiscountAmount         float64 `json:"discountAmount"`
		ExternalData           []any   `json:"externalData"`
		ItemDiscountInfos      []any   `json:"itemDiscountInfos"`
		CommerceItemID         string  `json:"commerceItemId"`
		Price                  float64 `json:"price"`
		Variant                []struct {
			OptionName  string `json:"optionName"`
			OptionValue string `json:"optionValue"`
		} `json:"variant"`
		PrimaryImageAltText string  `json:"primaryImageAltText"`
		OnSale              bool    `json:"onSale"`
		ID                  string  `json:"id"`
		State               string  `json:"state"`
		StateKey            string  `json:"stateKey"`
		UnitPrice           float64 `json:"unitPrice"`
		PrimaryImageTitle   string  `json:"primaryImageTitle"`
		ChildSKUs           []struct {
			PrimaryThumbImageURL any `json:"primaryThumbImageURL"`
		} `json:"childSKUs"`
		Amount                float64 `json:"amount"`
		Quantity              int     `json:"quantity"`
		ProductID             string  `json:"productId"`
		PointOfNoRevision     bool    `json:"pointOfNoRevision"`
		SalePrice             float64 `json:"salePrice"`
		OrderDiscountInfos    []any   `json:"orderDiscountInfos"`
		DetailedItemPriceInfo []struct {
			Discounted                 bool    `json:"discounted"`
			SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
			Amount                     float64 `json:"amount"`
			Quantity                   int     `json:"quantity"`
			ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
			Tax                        float64 `json:"tax"`
			OrderDiscountShare         float64 `json:"orderDiscountShare"`
			DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
			CurrencyCode               string  `json:"currencyCode"`
		} `json:"detailedItemPriceInfo"`
		GiftWithPurchaseCommerceItemMarkers []any  `json:"giftWithPurchaseCommerceItemMarkers"`
		CatRefID                            string `json:"catRefId"`
		SkuProperties                       []struct {
			Name         string `json:"name"`
			ID           string `json:"id"`
			Value        any    `json:"value"`
			PropertyType string `json:"propertyType"`
		} `json:"skuProperties"`
		Route        string `json:"route"`
		DiscountInfo []any  `json:"discountInfo"`
		SiteID       string `json:"siteId"`
		ShopperInput struct {
		} `json:"shopperInput"`
		Asset     bool    `json:"asset"`
		ListPrice float64 `json:"listPrice"`
	} `json:"items"`
}

type CatalogResp struct {
	TotalResults int `json:"totalResults"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	Links        []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Category struct {
		LongDescription any    `json:"longDescription"`
		Route           string `json:"route"`
		CategoryImages  []any  `json:"categoryImages"`
		DisplayName     string `json:"displayName"`
		RepositoryID    string `json:"repositoryId"`
		Active          bool   `json:"active"`
		Description     string `json:"description"`
		ID              string `json:"id"`
	} `json:"category"`
	Items []Item `json:"items"`
}

type Item struct {
	B2CExpertRatings any `json:"b2c_expertRatings"`
	B2CAge           any `json:"b2c_age"`
	OrderLimit       any `json:"orderLimit"`
	ListPrices       struct {
		DefaultPriceGroup float64 `json:"defaultPriceGroup"`
	} `json:"listPrices"`
	XVolumeUOM                   any    `json:"x_volumeUOM"`
	Type                         string `json:"type"`
	B2CTastingNotesBody          any    `json:"b2c_tastingNotesBody"`
	GcDescription                any    `json:"gc_description"`
	B2CLotteryPackageDescription any    `json:"b2c_lotteryPackageDescription"`
	Shippable                    bool   `json:"shippable"`
	B2CSizeSort                  string `json:"b2c_size_sort"`
	PrimaryImageAltText          string `json:"primaryImageAltText"`
	ID                           string `json:"id"`
	Brand                        string `json:"brand"`
	ParentCategories             []struct {
		RepositoryID          string `json:"repositoryId"`
		FixedParentCategories []struct {
			RepositoryID          string `json:"repositoryId"`
			FixedParentCategories []struct {
				RepositoryID          string `json:"repositoryId"`
				FixedParentCategories []any  `json:"fixedParentCategories"`
			} `json:"fixedParentCategories"`
		} `json:"fixedParentCategories"`
	} `json:"parentCategories"`
	Height                   any      `json:"height"`
	DefaultProductListingSku any      `json:"defaultProductListingSku"`
	Assetable                bool     `json:"assetable"`
	UnitOfMeasure            any      `json:"unitOfMeasure"`
	TargetAddOnProducts      []any    `json:"targetAddOnProducts"`
	B2CGlutenFree            string   `json:"b2c_glutenFree"`
	B2CDisableBopis          string   `json:"b2c_disableBopis"`
	B2CChairmansSelection    string   `json:"b2c_chairmansSelection"`
	SeoURLSlugDerived        string   `json:"seoUrlSlugDerived"`
	B2CRegion                any      `json:"b2c_region"`
	Active                   bool     `json:"active"`
	B2CUpc                   string   `json:"b2c_upc"`
	XSalesTaxIndicator       string   `json:"x_salesTaxIndicator"`
	ThumbImageURLs           []string `json:"thumbImageURLs"`
	B2CType                  string   `json:"b2c_type"`
	B2CPaResidencyOnly       string   `json:"b2c_paResidencyOnly"`
	B2COfferID               any      `json:"b2c_offerId"`
	B2CTastingNotes          string   `json:"b2c_tastingNotes"`
	B2CChairmansSpirits      string   `json:"b2c_chairmansSpirits"`
	Route                    string   `json:"route"`
	XVolumeOZ                string   `json:"x_volumeOZ"`
	RelatedArticles          []any    `json:"relatedArticles"`
	B2CProductIDNumber       string   `json:"b2c_productIdNumber"`
	MediumImageURLs          []string `json:"mediumImageURLs"`
	PrimarySourceImageURL    string   `json:"primarySourceImageURL"`
	SourceImageURLs          []string `json:"sourceImageURLs"`
	PrimaryThumbImageURL     string   `json:"primaryThumbImageURL"`
	DirectCatalogs           []any    `json:"directCatalogs"`
	Nonreturnable            bool     `json:"nonreturnable"`
	DisplayName              string   `json:"displayName"`
	B2CTaste                 any      `json:"b2c_taste"`
	B2CMostPopular           any      `json:"b2c_mostPopular"`
	PrimaryFullImageURL      string   `json:"primaryFullImageURL"`
	XFreightCost             any      `json:"x_freightCost"`
	ProductVariantOptions    []struct {
		VariantBasedDisplay     bool   `json:"variantBasedDisplay"`
		OptionID                string `json:"optionId"`
		ListingVariant          bool   `json:"listingVariant"`
		MapKeyPropertyAttribute string `json:"mapKeyPropertyAttribute"`
		OptionName              string `json:"optionName"`
		OptionValueMap          struct {
			Eaches int `json:"eaches"`
		} `json:"optionValueMap"`
	} `json:"productVariantOptions"`
	B2CExpertReviews          any    `json:"b2c_expertReviews"`
	PrimaryLargeImageURL      string `json:"primaryLargeImageURL"`
	B2CHighlyAllocatedProduct string `json:"b2c_highlyAllocatedProduct"`
	B2CInventoryAvailability  any    `json:"b2c_inventoryAvailability"`
	B2CVarietal               any    `json:"b2c_varietal"`
	SaleVolumePrices          struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"saleVolumePrices"`
	ChildSKUs []struct {
		DynamicPropertyMapLong struct {
			SkuB2CProductB2BLabel int `json:"sku-B2CProduct_b2b_label"`
		} `json:"dynamicPropertyMapLong"`
		BundleLinks     []any `json:"bundleLinks"`
		LargeImage      any   `json:"largeImage"`
		SmallImage      any   `json:"smallImage"`
		ListVolumePrice any   `json:"listVolumePrice"`
		OnlineOnly      bool  `json:"onlineOnly"`
		ListPrices      struct {
			DefaultPriceGroup float64 `json:"defaultPriceGroup"`
		} `json:"listPrices"`
		XUOM                  string `json:"x_uOM"`
		ConfigurationMetadata []any  `json:"configurationMetadata"`
		LargeImageURLs        []any  `json:"largeImageURLs"`
		ProductLine           any    `json:"productLine"`
		ListVolumePrices      struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"listVolumePrices"`
		DerivedSalePriceFrom        string `json:"derivedSalePriceFrom"`
		Model                       any    `json:"model"`
		Barcode                     any    `json:"barcode"`
		XSupplier                   any    `json:"x_supplier"`
		SalePriceEndDate            any    `json:"salePriceEndDate"`
		Images                      []any  `json:"images"`
		UnitOfMeasure               any    `json:"unitOfMeasure"`
		PrimaryMediumImageURL       any    `json:"primaryMediumImageURL"`
		DynamicPropertyMapBigString struct {
		} `json:"dynamicPropertyMapBigString"`
		Active                bool   `json:"active"`
		ThumbImageURLs        []any  `json:"thumbImageURLs"`
		XCaseSize             string `json:"x_caseSize"`
		MediumImageURLs       []any  `json:"mediumImageURLs"`
		PrimarySourceImageURL any    `json:"primarySourceImageURL"`
		SourceImageURLs       []any  `json:"sourceImageURLs"`
		PrimarySmallImageURL  any    `json:"primarySmallImageURL"`
		ProductFamily         any    `json:"productFamily"`
		PrimaryThumbImageURL  any    `json:"primaryThumbImageURL"`
		Nonreturnable         bool   `json:"nonreturnable"`
		DisplayName           any    `json:"displayName"`
		SalePrices            struct {
			DefaultPriceGroup float64 `json:"defaultPriceGroup"`
		} `json:"salePrices"`
		B2BBottleLabel               string `json:"b2b_bottleLabel"`
		PrimaryFullImageURL          any    `json:"primaryFullImageURL"`
		XSearchableSKU               string `json:"x_searchableSKU"`
		B2BLabel                     string `json:"b2b_label"`
		ProductListingSku            any    `json:"productListingSku"`
		PrimaryLargeImageURL         any    `json:"primaryLargeImageURL"`
		DerivedOnlineOnly            bool   `json:"derivedOnlineOnly"`
		SmallImageURLs               []any  `json:"smallImageURLs"`
		DerivedShippingSurchargeFrom string `json:"derivedShippingSurchargeFrom"`
		ShippingSurcharges           struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"shippingSurcharges"`
		ThumbnailImage   any `json:"thumbnailImage"`
		SaleVolumePrices struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"saleVolumePrices"`
		SaleVolumePrice    any       `json:"saleVolumePrice"`
		SalePriceStartDate any       `json:"salePriceStartDate"`
		Quantity           any       `json:"quantity"`
		LastModifiedDate   time.Time `json:"lastModifiedDate"`
		SalePrice          float64   `json:"salePrice"`
		FullImageURLs      []any     `json:"fullImageURLs"`
		VariantValuesOrder struct {
		} `json:"variantValuesOrder"`
		SoldAsPackage        bool    `json:"soldAsPackage"`
		ListingSKUID         any     `json:"listingSKUId"`
		RepositoryID         string  `json:"repositoryId"`
		DerivedListPriceFrom string  `json:"derivedListPriceFrom"`
		ShippingSurcharge    any     `json:"shippingSurcharge"`
		Configurable         bool    `json:"configurable"`
		ListPrice            float64 `json:"listPrice"`
	} `json:"childSKUs"`
	B2CCustomerRatingsFilterSplit any      `json:"b2c_customerRatingsFilterSplit"`
	SalePrice                     float64  `json:"salePrice"`
	XVolume                       string   `json:"x_volume"`
	B2CSortWeighting              any      `json:"b2c_sortWeighting"`
	B2CScotchType                 any      `json:"b2c_scotchType"`
	B2CTastingNotesOakInfluence   any      `json:"b2c_tastingNotesOakInfluence"`
	B2CTastingNotesSweetness      any      `json:"b2c_tastingNotesSweetness"`
	XDeliveryFee                  any      `json:"x_deliveryFee"`
	NotForIndividualSale          bool     `json:"notForIndividualSale"`
	Width                         any      `json:"width"`
	B2CExpertRatingsFilter        any      `json:"b2c_expertRatingsFilter"`
	DerivedListPriceFrom          string   `json:"derivedListPriceFrom"`
	DefaultParentCategory         any      `json:"defaultParentCategory"`
	B2CRegionFilterSplit          any      `json:"b2c_regionFilterSplit"`
	ListPrice                     float64  `json:"listPrice"`
	B2CQuotedAtPrice              any      `json:"b2c_quotedAtPrice"`
	XAlcoholicOrNonalcoholic      string   `json:"x_alcoholicOrNonalcoholic"`
	ListVolumePrice               any      `json:"listVolumePrice"`
	B2CFeaturedFilterSplit        any      `json:"b2c_featuredFilterSplit"`
	ExcludeFromSitemap            bool     `json:"excludeFromSitemap"`
	B2CFreightIncludedSalePrice   float64  `json:"b2c_freightIncludedSalePrice"`
	RelatedProducts               any      `json:"relatedProducts"`
	OnlineOnly                    bool     `json:"onlineOnly"`
	LargeImageURLs                []string `json:"largeImageURLs"`
	B2CFreightIncludedListPrice   float64  `json:"b2c_freightIncludedListPrice"`
	ListVolumePrices              struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"listVolumePrices"`
	AddOnProducts         []any     `json:"addOnProducts"`
	DerivedSalePriceFrom  string    `json:"derivedSalePriceFrom"`
	XType                 string    `json:"x_type"`
	B2CNew                any       `json:"b2c_new"`
	B2CTopCustomerReviews any       `json:"b2c_topCustomerReviews"`
	PrimaryMediumImageURL string    `json:"primaryMediumImageURL"`
	B2CCountry            string    `json:"b2c_country"`
	Weight                any       `json:"weight"`
	CreationDate          time.Time `json:"creationDate"`
	ParentCategoryIDPath  string    `json:"parentCategoryIdPath"`
	XTypeDisplay          string    `json:"x_typeDisplay"`
	ParentCategory        struct {
		RepositoryID          string `json:"repositoryId"`
		FixedParentCategories []struct {
			RepositoryID          string `json:"repositoryId"`
			FixedParentCategories []struct {
				RepositoryID          string `json:"repositoryId"`
				FixedParentCategories []any  `json:"fixedParentCategories"`
			} `json:"fixedParentCategories"`
		} `json:"fixedParentCategories"`
	} `json:"parentCategory"`
	PrimarySmallImageURL            string `json:"primarySmallImageURL"`
	B2CSalePriceType                string `json:"b2c_salePriceType"`
	B2CLimitPerOrder                any    `json:"b2c_limitPerOrder"`
	AvgCustRating                   any    `json:"avgCustRating"`
	B2CFeatured                     any    `json:"b2c_featured"`
	LongDescription                 any    `json:"longDescription"`
	B2CVintage                      any    `json:"b2c_vintage"`
	B2COnlineAvailable              string `json:"b2c_onlineAvailable"`
	B2CExpertRatingsFilterSplitSort any    `json:"b2c_expertRatingsFilterSplitSort"`
	Description                     any    `json:"description"`
	SalePrices                      struct {
		DefaultPriceGroup float64 `json:"defaultPriceGroup"`
	} `json:"salePrices"`
	B2COnlineExclusive                string   `json:"b2c_onlineExclusive"`
	B2CSpecialOrderAddressShip        string   `json:"b2c_specialOrderAddressShip"`
	B2CFreightIncludedActivePrice     float64  `json:"b2c_freightIncludedActivePrice"`
	B2CLotteryProduct                 string   `json:"b2c_lotteryProduct"`
	B2CLotteryAvailabilityDescription any      `json:"b2c_lotteryAvailabilityDescription"`
	SmallImageURLs                    []string `json:"smallImageURLs"`
	B2BHasCase                        any      `json:"b2b_hasCase"`
	DerivedShippingSurchargeFrom      string   `json:"derivedShippingSurchargeFrom"`
	ShippingSurcharges                struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"shippingSurcharges"`
	B2COrganic                        string    `json:"b2c_organic"`
	SaleVolumePrice                   any       `json:"saleVolumePrice"`
	PrimaryImageTitle                 string    `json:"primaryImageTitle"`
	B2CExpertRatingsFilterSplit       any       `json:"b2c_expertRatingsFilterSplit"`
	B2CSpecialOrderProduct            string    `json:"b2c_specialOrderProduct"`
	B2CClearance                      string    `json:"b2c_clearance"`
	RelatedMediaContent               []any     `json:"relatedMediaContent"`
	LastModifiedDate                  time.Time `json:"lastModifiedDate"`
	FullImageURLs                     []string  `json:"fullImageURLs"`
	B2CSize                           string    `json:"b2c_size"`
	Length                            any       `json:"length"`
	B2CProof                          string    `json:"b2c_proof"`
	DerivedDirectCatalogs             []any     `json:"derivedDirectCatalogs"`
	B2CFuturesProduct                 string    `json:"b2c_futuresProduct"`
	B2CLotteryRegistrationDescription any       `json:"b2c_lotteryRegistrationDescription"`
	B2CComingSoon                     string    `json:"b2c_comingSoon"`
	VariantValuesOrder                struct {
	} `json:"variantValuesOrder"`
	RepositoryID          string `json:"repositoryId"`
	ShippingSurcharge     any    `json:"shippingSurcharge"`
	ProductImagesMetadata []struct {
	} `json:"productImagesMetadata"`
	B2CMadeInPa  string `json:"b2c_madeInPa"`
	Configurable bool   `json:"configurable"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
	Links       []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
}

type CreateCartResp struct {
	TotalResults string `json:"totalResults"`
	Offset       string `json:"offset"`
	HasMore      string `json:"hasMore"`
	Limit        string `json:"limit"`
	Links        []struct {
		Method string `json:"method"`
		Rel    string `json:"rel"`
		Href   string `json:"href"`
	} `json:"links"`
	Embedded struct {
		Order struct {
			ShippingGroups []struct {
				ShippingGroupID string `json:"shippingGroupId"`
				ShippingAddress struct {
					LastName    any `json:"lastName"`
					Country     any `json:"country"`
					Address3    any `json:"address3"`
					Address2    any `json:"address2"`
					City        any `json:"city"`
					Prefix      any `json:"prefix"`
					Address1    any `json:"address1"`
					PostalCode  any `json:"postalCode"`
					CompanyName any `json:"companyName"`
					JobTitle    any `json:"jobTitle"`
					County      any `json:"county"`
					Suffix      any `json:"suffix"`
					FirstName   any `json:"firstName"`
					PhoneNumber any `json:"phoneNumber"`
					FaxNumber   any `json:"faxNumber"`
					Alias       any `json:"alias"`
					MiddleName  any `json:"middleName"`
					State       any `json:"state"`
					Email       any `json:"email"`
				} `json:"shippingAddress"`
				DiscountInfo struct {
					OrderDiscount    float64 `json:"orderDiscount"`
					DiscountDescList []any   `json:"discountDescList"`
					ShippingDiscount float64 `json:"shippingDiscount"`
				} `json:"discountInfo"`
				Type           string `json:"type"`
				Items          []any  `json:"items"`
				ShippingMethod struct {
					SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
					ShippingTax                float64 `json:"shippingTax"`
					Cost                       float64 `json:"cost"`
					TaxIncluded                bool    `json:"taxIncluded"`
					ExternalID                 any     `json:"externalId"`
					TaxCode                    string  `json:"taxCode"`
					Value                      string  `json:"value"`
					ShippingMethodDescription  string  `json:"shippingMethodDescription"`
				} `json:"shippingMethod"`
			} `json:"shippingGroups"`
			OrderID                string `json:"orderId"`
			AllowAlternateCurrency bool   `json:"allowAlternateCurrency"`
			DynamicProperties      []struct {
				Value any    `json:"value"`
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"dynamicProperties"`
			Payments       []any `json:"payments"`
			PriceListGroup struct {
				IsTaxIncluded bool   `json:"isTaxIncluded"`
				EndDate       any    `json:"endDate"`
				DisplayName   string `json:"displayName"`
				ListPriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"listPriceList"`
				Active                     bool   `json:"active"`
				IsPointsBased              bool   `json:"isPointsBased"`
				Locale                     string `json:"locale"`
				ShippingSurchargePriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"shippingSurchargePriceList"`
				Deleted            bool   `json:"deleted"`
				TaxCalculationType string `json:"taxCalculationType"`
				RepositoryID       string `json:"repositoryId"`
				SalePriceList      struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"salePriceList"`
				Currency struct {
					CurrencyType     any    `json:"currencyType"`
					Symbol           string `json:"symbol"`
					Deleted          bool   `json:"deleted"`
					DisplayName      string `json:"displayName"`
					RepositoryID     string `json:"repositoryId"`
					FractionalDigits int    `json:"fractionalDigits"`
					CurrencyCode     string `json:"currencyCode"`
					NumericCode      string `json:"numericCode"`
				} `json:"currency"`
				ID                 string `json:"id"`
				IncludeAllProducts bool   `json:"includeAllProducts"`
				StartDate          any    `json:"startDate"`
			} `json:"priceListGroup"`
			OrderAction string `json:"orderAction"`
			PriceInfo   struct {
				Amount                 float64 `json:"amount"`
				Total                  float64 `json:"total"`
				Shipping               float64 `json:"shipping"`
				ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
				Tax                    float64 `json:"tax"`
				SubTotal               float64 `json:"subTotal"`
				CurrencyCode           string  `json:"currencyCode"`
				TotalWithoutTax        float64 `json:"totalWithoutTax"`
			} `json:"priceInfo"`
			DiscountInfo struct {
				UnclaimedCouponMultiPromotions struct {
				} `json:"unclaimedCouponMultiPromotions"`
				OrderCouponsMap struct {
				} `json:"orderCouponsMap"`
				OrderDiscount             float64 `json:"orderDiscount"`
				ShippingDiscount          float64 `json:"shippingDiscount"`
				OrderImplicitDiscountList []any   `json:"orderImplicitDiscountList"`
				UnclaimedCouponsMap       struct {
				} `json:"unclaimedCouponsMap"`
				ClaimedCouponMultiPromotions struct {
				} `json:"claimedCouponMultiPromotions"`
			} `json:"discountInfo"`
			ShoppingCart struct {
				Items         []any `json:"items"`
				NumberOfItems int   `json:"numberOfItems"`
			} `json:"shoppingCart"`
			GiftWithPurchaseInfo         []any  `json:"giftWithPurchaseInfo"`
			SiteID                       string `json:"siteId"`
			Markers                      []any  `json:"markers"`
			GiftWithPurchaseOrderMarkers []any  `json:"giftWithPurchaseOrderMarkers"`
		} `json:"order"`
	} `json:"embedded"`
	Items []struct {
		DiscountInfo struct {
			OrderDiscount    float64 `json:"orderDiscount"`
			DiscountDescList []any   `json:"discountDescList"`
			ShippingDiscount float64 `json:"shippingDiscount"`
		} `json:"discountInfo"`
		TrackingInfo   []any `json:"trackingInfo"`
		ShippingMethod struct {
			SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
			ShippingTax                float64 `json:"shippingTax"`
			Cost                       float64 `json:"cost"`
			TaxIncluded                bool    `json:"taxIncluded"`
			ExternalID                 any     `json:"externalId"`
			TaxCode                    string  `json:"taxCode"`
			Value                      string  `json:"value"`
			ShippingMethodDescription  string  `json:"shippingMethodDescription"`
		} `json:"shippingMethod"`
		ShippingGroupID string `json:"shippingGroupId"`
		ShippingAddress struct {
			LastName    any `json:"lastName"`
			Country     any `json:"country"`
			Address3    any `json:"address3"`
			Address2    any `json:"address2"`
			City        any `json:"city"`
			Prefix      any `json:"prefix"`
			Address1    any `json:"address1"`
			PostalCode  any `json:"postalCode"`
			CompanyName any `json:"companyName"`
			JobTitle    any `json:"jobTitle"`
			County      any `json:"county"`
			Suffix      any `json:"suffix"`
			FirstName   any `json:"firstName"`
			PhoneNumber any `json:"phoneNumber"`
			FaxNumber   any `json:"faxNumber"`
			Alias       any `json:"alias"`
			MiddleName  any `json:"middleName"`
			State       any `json:"state"`
			Email       any `json:"email"`
		} `json:"shippingAddress,omitempty"`
		Type   string `json:"type"`
		Items  []any  `json:"items"`
		Status string `json:"status"`
	} `json:"items"`
}

type AccountInformation struct {
	LastPurchaseDate              any    `json:"lastPurchaseDate"`
	LastName                      string `json:"lastName"`
	GDPRProfileP13NConsentDate    any    `json:"GDPRProfileP13nConsentDate"`
	GDPRProfileP13NConsentGranted bool   `json:"GDPRProfileP13nConsentGranted"`
	Gender                        string `json:"gender"`
	Catalog                       struct {
		RepositoryID string `json:"repositoryId"`
	} `json:"catalog"`
	DynamicProperties []struct {
		UIEditorType string `json:"uiEditorType"`
		Default      string `json:"default"`
		Length       any    `json:"length"`
		ID           string `json:"id"`
		Label        string `json:"label"`
		Type         string `json:"type"`
		Value        any    `json:"value"`
		Required     bool   `json:"required"`
	} `json:"dynamicProperties"`
	Roles              []any `json:"roles"`
	SecondaryAddresses struct {
		Address struct {
			Country           string `json:"country"`
			LastName          string `json:"lastName"`
			Types             []any  `json:"types"`
			Address3          any    `json:"address3"`
			City              string `json:"city"`
			Address2          any    `json:"address2"`
			Prefix            any    `json:"prefix"`
			Address1          string `json:"address1"`
			PostalCode        string `json:"postalCode"`
			JobTitle          any    `json:"jobTitle"`
			CompanyName       any    `json:"companyName"`
			County            any    `json:"county"`
			Suffix            any    `json:"suffix"`
			FirstName         string `json:"firstName"`
			ExternalAddressID any    `json:"externalAddressId"`
			PhoneNumber       any    `json:"phoneNumber"`
			RepositoryID      string `json:"repositoryId"`
			FaxNumber         any    `json:"faxNumber"`
			MiddleName        any    `json:"middleName"`
			State             string `json:"state"`
		} `json:"Address"`
	} `json:"secondaryAddresses"`
	NumberOfOrders     int       `json:"numberOfOrders"`
	Login              string    `json:"login"`
	Locale             string    `json:"locale"`
	ParentOrganization any       `json:"parentOrganization"`
	ReceiveEmailDate   any       `json:"receiveEmailDate"`
	FirstPurchaseDate  any       `json:"firstPurchaseDate"`
	LifetimeSpend      float64   `json:"lifetimeSpend"`
	LoyaltyPrograms    []any     `json:"loyaltyPrograms"`
	LastPurchaseAmount float64   `json:"lastPurchaseAmount"`
	RegistrationDate   time.Time `json:"registrationDate"`
	Links              []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	ID                     string  `json:"id"`
	LifetimeAOV            float64 `json:"lifetimeAOV"`
	Email                  string  `json:"email"`
	DaytimeTelephoneNumber any     `json:"daytimeTelephoneNumber"`
	NumberOfVisits         int     `json:"numberOfVisits"`
	CustomerContactID      any     `json:"customerContactId"`
	TaxExempt              bool    `json:"taxExempt"`
	DerivedOrderPriceLimit any     `json:"derivedOrderPriceLimit"`
	ContactBillingAddress  any     `json:"contactBillingAddress"`
	ReceiveEmail           string  `json:"receiveEmail"`
	PriceListGroup         struct {
		IsTaxIncluded bool   `json:"isTaxIncluded"`
		EndDate       any    `json:"endDate"`
		DisplayName   string `json:"displayName"`
		ListPriceList struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"listPriceList"`
		Active                     bool   `json:"active"`
		IsPointsBased              bool   `json:"isPointsBased"`
		Locale                     string `json:"locale"`
		ShippingSurchargePriceList struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"shippingSurchargePriceList"`
		Deleted            bool   `json:"deleted"`
		TaxCalculationType string `json:"taxCalculationType"`
		RepositoryID       string `json:"repositoryId"`
		SalePriceList      struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"salePriceList"`
		Currency struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"currency"`
		ID                 string `json:"id"`
		IncludeAllProducts bool   `json:"includeAllProducts"`
		StartDate          any    `json:"startDate"`
	} `json:"priceListGroup"`
	DateOfBirth            any `json:"dateOfBirth"`
	ContactShippingAddress struct {
		Country           string `json:"country"`
		LastName          string `json:"lastName"`
		Types             []any  `json:"types"`
		Address3          any    `json:"address3"`
		City              string `json:"city"`
		Address2          any    `json:"address2"`
		Prefix            any    `json:"prefix"`
		Address1          string `json:"address1"`
		PostalCode        string `json:"postalCode"`
		JobTitle          any    `json:"jobTitle"`
		CompanyName       any    `json:"companyName"`
		County            any    `json:"county"`
		Suffix            any    `json:"suffix"`
		FirstName         string `json:"firstName"`
		ExternalAddressID any    `json:"externalAddressId"`
		PhoneNumber       any    `json:"phoneNumber"`
		RepositoryID      string `json:"repositoryId"`
		FaxNumber         any    `json:"faxNumber"`
		MiddleName        any    `json:"middleName"`
		State             string `json:"state"`
	} `json:"contactShippingAddress"`
	LastVisitDate           time.Time `json:"lastVisitDate"`
	PreviousVisitDate       time.Time `json:"previousVisitDate"`
	FirstName               string    `json:"firstName"`
	DerivedApprovalRequired bool      `json:"derivedApprovalRequired"`
	LifetimeCurrencyCode    any       `json:"lifetimeCurrencyCode"`
	RepositoryID            string    `json:"repositoryId"`
	ShippingAddress         struct {
		Country           string `json:"country"`
		LastName          string `json:"lastName"`
		Types             []any  `json:"types"`
		Address3          any    `json:"address3"`
		City              string `json:"city"`
		Address2          any    `json:"address2"`
		Prefix            any    `json:"prefix"`
		Address1          string `json:"address1"`
		PostalCode        string `json:"postalCode"`
		JobTitle          any    `json:"jobTitle"`
		CompanyName       any    `json:"companyName"`
		County            any    `json:"county"`
		Suffix            any    `json:"suffix"`
		FirstName         string `json:"firstName"`
		ExternalAddressID any    `json:"externalAddressId"`
		PhoneNumber       any    `json:"phoneNumber"`
		RepositoryID      string `json:"repositoryId"`
		FaxNumber         any    `json:"faxNumber"`
		Alias             string `json:"alias"`
		MiddleName        any    `json:"middleName"`
		State             string `json:"state"`
	} `json:"shippingAddress"`
	FirstVisitDate         time.Time `json:"firstVisitDate"`
	CurrentOrganization    any       `json:"currentOrganization"`
	SecondaryOrganizations []any     `json:"secondaryOrganizations"`
	ShippingAddresses      []struct {
		Country           string `json:"country"`
		LastName          string `json:"lastName"`
		Types             []any  `json:"types"`
		Address3          any    `json:"address3"`
		City              string `json:"city"`
		Address2          any    `json:"address2"`
		Prefix            any    `json:"prefix"`
		Address1          string `json:"address1"`
		PostalCode        string `json:"postalCode"`
		JobTitle          any    `json:"jobTitle"`
		CompanyName       any    `json:"companyName"`
		RegionName        string `json:"regionName"`
		County            any    `json:"county"`
		IsDefaultAddress  bool   `json:"isDefaultAddress"`
		Suffix            any    `json:"suffix"`
		FirstName         string `json:"firstName"`
		ExternalAddressID any    `json:"externalAddressId"`
		PhoneNumber       any    `json:"phoneNumber"`
		RepositoryID      string `json:"repositoryId"`
		FaxNumber         any    `json:"faxNumber"`
		Alias             string `json:"alias"`
		MiddleName        any    `json:"middleName"`
		State             string `json:"state"`
		CountryName       string `json:"countryName"`
	} `json:"shippingAddresses"`
}

type OrderBody struct {
	TaxPriceInfo struct {
		CityTax                    float64 `json:"cityTax"`
		SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
		Amount                     float64 `json:"amount"`
		ValueAddedTax              float64 `json:"valueAddedTax"`
		CountyTax                  float64 `json:"countyTax"`
		IsTaxIncluded              bool    `json:"isTaxIncluded"`
		MiscTax                    float64 `json:"miscTax"`
		DistrictTax                float64 `json:"districtTax"`
		StateTax                   float64 `json:"stateTax"`
		CountryTax                 float64 `json:"countryTax"`
	} `json:"taxPriceInfo"`
	PriceInfo struct {
		Amount                  float64 `json:"amount"`
		Total                   float64 `json:"total"`
		LkpValExcludingFreeShip any     `json:"lkpValExcludingFreeShip"`
		Shipping                float64 `json:"shipping"`
		ShippingSurchargeValue  float64 `json:"shippingSurchargeValue"`
		Tax                     float64 `json:"tax"`
		SubTotal                float64 `json:"subTotal"`
		CurrencyCode            string  `json:"currencyCode"`
		TotalWithoutTax         float64 `json:"totalWithoutTax"`
	} `json:"priceInfo"`
	DiscountInfo struct {
		OrderDiscount    float64 `json:"orderDiscount"`
		DiscountDescList []any   `json:"discountDescList"`
		ShippingDiscount float64 `json:"shippingDiscount"`
	} `json:"discountInfo"`
	TrackingInfo   []any `json:"trackingInfo"`
	ShippingMethod struct {
		SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
		ShippingTax                float64 `json:"shippingTax"`
		Cost                       float64 `json:"cost"`
		TaxIncluded                bool    `json:"taxIncluded"`
		ExternalID                 string  `json:"externalId"`
		TaxCode                    string  `json:"taxCode"`
		Value                      string  `json:"value"`
		ShippingMethodDescription  string  `json:"shippingMethodDescription"`
	} `json:"shippingMethod"`
	ShippingGroupID string `json:"shippingGroupId"`
	ShippingAddress struct {
		LastName    string `json:"lastName"`
		Country     string `json:"country"`
		Address3    any    `json:"address3"`
		Address2    any    `json:"address2"`
		City        string `json:"city"`
		Prefix      any    `json:"prefix"`
		Address1    string `json:"address1"`
		PostalCode  string `json:"postalCode"`
		CompanyName any    `json:"companyName"`
		JobTitle    any    `json:"jobTitle"`
		County      any    `json:"county"`
		Suffix      any    `json:"suffix"`
		FirstName   string `json:"firstName"`
		PhoneNumber string `json:"phoneNumber"`
		FaxNumber   string `json:"faxNumber"`
		Alias       any    `json:"alias"`
		MiddleName  any    `json:"middleName"`
		State       string `json:"state"`
		Email       string `json:"email"`
	} `json:"shippingAddress"`
	Links []struct {
		Method string `json:"method"`
		Rel    string `json:"rel"`
		Href   string `json:"href"`
	} `json:"links"`
	Type  string `json:"type"`
	Items []struct {
		PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
		RawTotalPrice        float64 `json:"rawTotalPrice"`
		ReturnedQuantity     int     `json:"returnedQuantity"`
		DynamicProperties    []struct {
			Value any    `json:"value"`
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"dynamicProperties"`
		DisplayName            string  `json:"displayName"`
		ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
		AvailabilityDate       any     `json:"availabilityDate"`
		ExternalData           []any   `json:"externalData"`
		DiscountAmount         float64 `json:"discountAmount"`
		PreOrderQuantity       int     `json:"preOrderQuantity"`
		CommerceItemID         string  `json:"commerceItemId"`
		Price                  float64 `json:"price"`
		Variant                []struct {
			OptionName  string `json:"optionName"`
			OptionValue string `json:"optionValue"`
		} `json:"variant"`
		OnSale                bool    `json:"onSale"`
		PrimaryImageAltText   string  `json:"primaryImageAltText"`
		StateDetailsAsUser    string  `json:"stateDetailsAsUser"`
		CommerceID            string  `json:"commerceId"`
		UnitPrice             float64 `json:"unitPrice"`
		PrimaryImageTitle     string  `json:"primaryImageTitle"`
		Amount                float64 `json:"amount"`
		Quantity              int     `json:"quantity"`
		PointOfNoRevision     bool    `json:"pointOfNoRevision"`
		RelationshipType      string  `json:"relationshipType"`
		ProductID             string  `json:"productId"`
		SalePrice             float64 `json:"salePrice"`
		DetailedItemPriceInfo []struct {
			Discounted                 bool    `json:"discounted"`
			SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
			Amount                     float64 `json:"amount"`
			Quantity                   int     `json:"quantity"`
			ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
			Tax                        float64 `json:"tax"`
			OrderDiscountShare         float64 `json:"orderDiscountShare"`
			DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
			CurrencyCode               string  `json:"currencyCode"`
		} `json:"detailedItemPriceInfo"`
		Active        bool   `json:"active"`
		CatRefID      string `json:"catRefId"`
		SkuProperties []struct {
			Name         string `json:"name"`
			ID           string `json:"id"`
			Value        any    `json:"value"`
			PropertyType string `json:"propertyType"`
		} `json:"skuProperties"`
		DiscountInfo []any  `json:"discountInfo"`
		Route        string `json:"route"`
		SiteID       string `json:"siteId"`
		ShopperInput struct {
		} `json:"shopperInput"`
		Asset             bool    `json:"asset"`
		BackOrderQuantity int     `json:"backOrderQuantity"`
		ListPrice         float64 `json:"listPrice"`
		Status            string  `json:"status"`
	} `json:"items"`
	Embedded struct {
		Order struct {
			ShippingGroups []struct {
				TaxPriceInfo struct {
					CityTax                    float64 `json:"cityTax"`
					SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
					Amount                     float64 `json:"amount"`
					ValueAddedTax              float64 `json:"valueAddedTax"`
					CountyTax                  float64 `json:"countyTax"`
					IsTaxIncluded              bool    `json:"isTaxIncluded"`
					MiscTax                    float64 `json:"miscTax"`
					DistrictTax                float64 `json:"districtTax"`
					StateTax                   float64 `json:"stateTax"`
					CountryTax                 float64 `json:"countryTax"`
				} `json:"taxPriceInfo"`
				PriceInfo struct {
					Amount                  float64 `json:"amount"`
					Total                   float64 `json:"total"`
					LkpValExcludingFreeShip any     `json:"lkpValExcludingFreeShip"`
					Shipping                float64 `json:"shipping"`
					ShippingSurchargeValue  float64 `json:"shippingSurchargeValue"`
					Tax                     float64 `json:"tax"`
					SubTotal                float64 `json:"subTotal"`
					CurrencyCode            string  `json:"currencyCode"`
					TotalWithoutTax         float64 `json:"totalWithoutTax"`
				} `json:"priceInfo"`
				DiscountInfo struct {
					OrderDiscount    float64 `json:"orderDiscount"`
					DiscountDescList []any   `json:"discountDescList"`
					ShippingDiscount float64 `json:"shippingDiscount"`
				} `json:"discountInfo"`
				ShippingMethod struct {
					SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
					ShippingTax                float64 `json:"shippingTax"`
					Cost                       float64 `json:"cost"`
					TaxIncluded                bool    `json:"taxIncluded"`
					ExternalID                 string  `json:"externalId"`
					TaxCode                    string  `json:"taxCode"`
					Value                      string  `json:"value"`
					ShippingMethodDescription  string  `json:"shippingMethodDescription"`
				} `json:"shippingMethod"`
				ShippingGroupID string `json:"shippingGroupId"`
				ShippingAddress struct {
					LastName    string `json:"lastName"`
					Country     string `json:"country"`
					Address3    any    `json:"address3"`
					Address2    any    `json:"address2"`
					City        string `json:"city"`
					Prefix      any    `json:"prefix"`
					Address1    string `json:"address1"`
					PostalCode  string `json:"postalCode"`
					CompanyName any    `json:"companyName"`
					JobTitle    any    `json:"jobTitle"`
					County      any    `json:"county"`
					Suffix      any    `json:"suffix"`
					FirstName   string `json:"firstName"`
					PhoneNumber string `json:"phoneNumber"`
					FaxNumber   string `json:"faxNumber"`
					Alias       any    `json:"alias"`
					MiddleName  any    `json:"middleName"`
					State       string `json:"state"`
					Email       string `json:"email"`
				} `json:"shippingAddress"`
				Type  string `json:"type"`
				Items []struct {
					RawTotalPrice     float64 `json:"rawTotalPrice"`
					ReturnedQuantity  int     `json:"returnedQuantity"`
					DynamicProperties []struct {
						Value any    `json:"value"`
						ID    string `json:"id"`
						Label string `json:"label"`
					} `json:"dynamicProperties"`
					ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
					AvailabilityDate       any     `json:"availabilityDate"`
					ExternalData           []any   `json:"externalData"`
					DiscountAmount         float64 `json:"discountAmount"`
					PreOrderQuantity       int     `json:"preOrderQuantity"`
					CommerceItemID         string  `json:"commerceItemId"`
					Price                  float64 `json:"price"`
					OnSale                 bool    `json:"onSale"`
					StateDetailsAsUser     string  `json:"stateDetailsAsUser"`
					CommerceID             string  `json:"commerceId"`
					UnitPrice              float64 `json:"unitPrice"`
					Amount                 float64 `json:"amount"`
					Quantity               int     `json:"quantity"`
					PointOfNoRevision      bool    `json:"pointOfNoRevision"`
					RelationshipType       string  `json:"relationshipType"`
					ProductID              string  `json:"productId"`
					SalePrice              float64 `json:"salePrice"`
					DetailedItemPriceInfo  []struct {
						Discounted                 bool    `json:"discounted"`
						SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
						Amount                     float64 `json:"amount"`
						Quantity                   int     `json:"quantity"`
						ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
						Tax                        float64 `json:"tax"`
						OrderDiscountShare         float64 `json:"orderDiscountShare"`
						DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
						CurrencyCode               string  `json:"currencyCode"`
					} `json:"detailedItemPriceInfo"`
					CatRefID     string `json:"catRefId"`
					DiscountInfo []any  `json:"discountInfo"`
					SiteID       string `json:"siteId"`
					ShopperInput struct {
					} `json:"shopperInput"`
					Asset             bool    `json:"asset"`
					BackOrderQuantity int     `json:"backOrderQuantity"`
					ListPrice         float64 `json:"listPrice"`
					Status            string  `json:"status"`
				} `json:"items"`
			} `json:"shippingGroups"`
			CreationSiteID         string `json:"creationSiteId"`
			OrderID                string `json:"orderId"`
			AllowAlternateCurrency bool   `json:"allowAlternateCurrency"`
			DynamicProperties      []struct {
				Value any    `json:"value"`
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"dynamicProperties"`
			Payments       []any `json:"payments"`
			ShippingMethod struct {
				ShippingTax float64 `json:"shippingTax"`
				Value       string  `json:"value"`
				Cost        float64 `json:"cost"`
			} `json:"shippingMethod"`
			PriceListGroup struct {
				IsTaxIncluded bool   `json:"isTaxIncluded"`
				EndDate       any    `json:"endDate"`
				DisplayName   string `json:"displayName"`
				ListPriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"listPriceList"`
				Active                     bool   `json:"active"`
				IsPointsBased              bool   `json:"isPointsBased"`
				Locale                     string `json:"locale"`
				ShippingSurchargePriceList struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"shippingSurchargePriceList"`
				Deleted            bool   `json:"deleted"`
				TaxCalculationType string `json:"taxCalculationType"`
				RepositoryID       string `json:"repositoryId"`
				SalePriceList      struct {
					RepositoryID string `json:"repositoryId"`
				} `json:"salePriceList"`
				Currency struct {
					CurrencyType     any    `json:"currencyType"`
					Symbol           string `json:"symbol"`
					Deleted          bool   `json:"deleted"`
					DisplayName      string `json:"displayName"`
					RepositoryID     string `json:"repositoryId"`
					FractionalDigits int    `json:"fractionalDigits"`
					CurrencyCode     string `json:"currencyCode"`
					NumericCode      string `json:"numericCode"`
				} `json:"currency"`
				ID                 string `json:"id"`
				IncludeAllProducts bool   `json:"includeAllProducts"`
				StartDate          any    `json:"startDate"`
			} `json:"priceListGroup"`
			OrderAction string `json:"orderAction"`
			PriceInfo   struct {
				Amount                 float64 `json:"amount"`
				Total                  float64 `json:"total"`
				Shipping               float64 `json:"shipping"`
				ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
				Tax                    float64 `json:"tax"`
				SubTotal               float64 `json:"subTotal"`
				CurrencyCode           string  `json:"currencyCode"`
				TotalWithoutTax        float64 `json:"totalWithoutTax"`
			} `json:"priceInfo"`
			DiscountInfo struct {
				UnclaimedCouponMultiPromotions struct {
				} `json:"unclaimedCouponMultiPromotions"`
				OrderCouponsMap struct {
				} `json:"orderCouponsMap"`
				OrderDiscount             float64 `json:"orderDiscount"`
				ShippingDiscount          float64 `json:"shippingDiscount"`
				OrderImplicitDiscountList []any   `json:"orderImplicitDiscountList"`
				UnclaimedCouponsMap       struct {
				} `json:"unclaimedCouponsMap"`
				ClaimedCouponMultiPromotions struct {
				} `json:"claimedCouponMultiPromotions"`
			} `json:"discountInfo"`
			ShoppingCart struct {
				Items []struct {
					PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
					RawTotalPrice        float64 `json:"rawTotalPrice"`
					DisplayName          string  `json:"displayName"`
					DynamicProperties    []struct {
						Value any    `json:"value"`
						ID    string `json:"id"`
						Label string `json:"label"`
					} `json:"dynamicProperties"`
					ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
					DiscountAmount         float64 `json:"discountAmount"`
					ExternalData           []any   `json:"externalData"`
					Description            any     `json:"description"`
					IsItemValid            bool    `json:"isItemValid"`
					ItemDiscountInfos      []any   `json:"itemDiscountInfos"`
					CommerceItemID         string  `json:"commerceItemId"`
					Price                  float64 `json:"price"`
					Variant                []struct {
						OptionName  string `json:"optionName"`
						OptionValue string `json:"optionValue"`
					} `json:"variant"`
					PrimaryImageAltText string  `json:"primaryImageAltText"`
					OnSale              bool    `json:"onSale"`
					ID                  string  `json:"id"`
					State               string  `json:"state"`
					StateKey            string  `json:"stateKey"`
					UnitPrice           float64 `json:"unitPrice"`
					PrimaryImageTitle   string  `json:"primaryImageTitle"`
					ChildSKUs           []struct {
						PrimaryThumbImageURL any `json:"primaryThumbImageURL"`
					} `json:"childSKUs"`
					Amount                float64 `json:"amount"`
					Quantity              int     `json:"quantity"`
					ProductID             string  `json:"productId"`
					PointOfNoRevision     bool    `json:"pointOfNoRevision"`
					SalePrice             float64 `json:"salePrice"`
					OrderDiscountInfos    []any   `json:"orderDiscountInfos"`
					DetailedItemPriceInfo []struct {
						Discounted                 bool    `json:"discounted"`
						SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
						Amount                     float64 `json:"amount"`
						Quantity                   int     `json:"quantity"`
						ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
						Tax                        float64 `json:"tax"`
						OrderDiscountShare         float64 `json:"orderDiscountShare"`
						DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
						CurrencyCode               string  `json:"currencyCode"`
					} `json:"detailedItemPriceInfo"`
					GiftWithPurchaseCommerceItemMarkers []any  `json:"giftWithPurchaseCommerceItemMarkers"`
					OriginalCommerceItemID              any    `json:"originalCommerceItemId"`
					TaxCode                             any    `json:"taxCode"`
					CatRefID                            string `json:"catRefId"`
					SkuProperties                       []struct {
						Name         string `json:"name"`
						ID           string `json:"id"`
						Value        any    `json:"value"`
						PropertyType string `json:"propertyType"`
					} `json:"skuProperties"`
					Route        string `json:"route"`
					DiscountInfo []any  `json:"discountInfo"`
					SiteID       string `json:"siteId"`
					ShopperInput struct {
					} `json:"shopperInput"`
					Asset     bool    `json:"asset"`
					ListPrice float64 `json:"listPrice"`
				} `json:"items"`
				NumberOfItems int `json:"numberOfItems"`
			} `json:"shoppingCart"`
			GiftWithPurchaseInfo         []any  `json:"giftWithPurchaseInfo"`
			SiteID                       string `json:"siteId"`
			Markers                      []any  `json:"markers"`
			GiftWithPurchaseOrderMarkers []any  `json:"giftWithPurchaseOrderMarkers"`
		} `json:"order"`
	} `json:"embedded"`
	Status string `json:"status"`
}

type OrderResponse struct {
	CreationTime      int64  `json:"creationTime"`
	SourceSystem      string `json:"sourceSystem"`
	DynamicProperties []struct {
		Value any    `json:"value"`
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"dynamicProperties"`
	Payments []struct {
		PaymentGroupID          string  `json:"paymentGroupId"`
		Amount                  float64 `json:"amount"`
		CustomPaymentProperties struct {
			CardType string `json:"cardType"`
			Exp      string `json:"exp"`
			Token    string `json:"token"`
		} `json:"customPaymentProperties"`
		GatewayName       string `json:"gatewayName"`
		UIIntervention    any    `json:"uiIntervention"`
		PaymentMethod     string `json:"paymentMethod"`
		IsAmountRemaining bool   `json:"isAmountRemaining"`
		PaymentState      string `json:"paymentState"`
		Message           string `json:"message"`
		Type              string `json:"type"`
		CurrencyCode      string `json:"currencyCode"`
	} `json:"payments"`
	UUID      string `json:"uuid"`
	PriceInfo struct {
		Amount                 float64 `json:"amount"`
		Total                  float64 `json:"total"`
		Shipping               float64 `json:"shipping"`
		ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
		OrderTotalBySite       struct {
			SiteUS float64 `json:"siteUS"`
		} `json:"orderTotalBySite"`
		Tax             float64 `json:"tax"`
		SubTotal        float64 `json:"subTotal"`
		CurrencyCode    string  `json:"currencyCode"`
		TotalWithoutTax float64 `json:"totalWithoutTax"`
	} `json:"priceInfo"`
	ShoppingCart struct {
		Items []struct {
			PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
			RawTotalPrice        float64 `json:"rawTotalPrice"`
			DisplayName          string  `json:"displayName"`
			DynamicProperties    []struct {
				Value any    `json:"value"`
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"dynamicProperties"`
			ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
			DiscountAmount         float64 `json:"discountAmount"`
			ExternalData           []any   `json:"externalData"`
			ItemDiscountInfos      []any   `json:"itemDiscountInfos"`
			ExternalPrice          float64 `json:"externalPrice"`
			CommerceItemID         string  `json:"commerceItemId"`
			Price                  float64 `json:"price"`
			Variant                []struct {
				OptionName  string `json:"optionName"`
				OptionValue string `json:"optionValue"`
			} `json:"variant"`
			PrimaryImageAltText string  `json:"primaryImageAltText"`
			ID                  string  `json:"id"`
			State               string  `json:"state"`
			StateKey            string  `json:"stateKey"`
			UnitPrice           float64 `json:"unitPrice"`
			PrimaryImageTitle   string  `json:"primaryImageTitle"`
			ChildSKUs           []struct {
				PrimaryThumbImageURL any `json:"primaryThumbImageURL"`
			} `json:"childSKUs"`
			Amount                float64 `json:"amount"`
			Quantity              int     `json:"quantity"`
			ProductID             string  `json:"productId"`
			PointOfNoRevision     bool    `json:"pointOfNoRevision"`
			SalePrice             float64 `json:"salePrice"`
			OrderDiscountInfos    []any   `json:"orderDiscountInfos"`
			DetailedItemPriceInfo []struct {
				Discounted                 bool    `json:"discounted"`
				SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
				Amount                     float64 `json:"amount"`
				Quantity                   int     `json:"quantity"`
				ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
				Tax                        float64 `json:"tax"`
				OrderDiscountShare         float64 `json:"orderDiscountShare"`
				DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
				CurrencyCode               string  `json:"currencyCode"`
			} `json:"detailedItemPriceInfo"`
			GiftWithPurchaseCommerceItemMarkers []any  `json:"giftWithPurchaseCommerceItemMarkers"`
			OriginalCommerceItemID              any    `json:"originalCommerceItemId"`
			ExternalPriceQuantity               int    `json:"externalPriceQuantity"`
			CatRefID                            string `json:"catRefId"`
			SkuProperties                       []struct {
				Name         string `json:"name"`
				ID           string `json:"id"`
				Value        any    `json:"value"`
				PropertyType string `json:"propertyType"`
			} `json:"skuProperties"`
			Route        string `json:"route"`
			DiscountInfo []any  `json:"discountInfo"`
			SiteID       string `json:"siteId"`
			ShopperInput struct {
			} `json:"shopperInput"`
			Asset     bool    `json:"asset"`
			ListPrice float64 `json:"listPrice"`
		} `json:"items"`
	} `json:"shoppingCart"`
	Links []struct {
		Method string `json:"method"`
		Rel    string `json:"rel"`
		Href   string `json:"href"`
	} `json:"links"`
	ID           string `json:"id"`
	State        string `json:"state"`
	TaxPriceInfo struct {
		CityTax                    float64 `json:"cityTax"`
		SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
		Amount                     float64 `json:"amount"`
		ValueAddedTax              float64 `json:"valueAddedTax"`
		CountyTax                  float64 `json:"countyTax"`
		IsTaxIncluded              bool    `json:"isTaxIncluded"`
		MiscTax                    float64 `json:"miscTax"`
		DistrictTax                float64 `json:"districtTax"`
		StateTax                   float64 `json:"stateTax"`
		CountryTax                 float64 `json:"countryTax"`
	} `json:"taxPriceInfo"`
	ShippingGroups []struct {
		TaxPriceInfo struct {
			CityTax                    float64 `json:"cityTax"`
			SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
			Amount                     float64 `json:"amount"`
			ValueAddedTax              float64 `json:"valueAddedTax"`
			CountyTax                  float64 `json:"countyTax"`
			IsTaxIncluded              bool    `json:"isTaxIncluded"`
			MiscTax                    float64 `json:"miscTax"`
			DistrictTax                float64 `json:"districtTax"`
			StateTax                   float64 `json:"stateTax"`
			CountryTax                 float64 `json:"countryTax"`
		} `json:"taxPriceInfo"`
		ShippingMethod struct {
			SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
			ShippingTax                float64 `json:"shippingTax"`
			Cost                       float64 `json:"cost"`
			TaxIncluded                bool    `json:"taxIncluded"`
			ExternalID                 string  `json:"externalId"`
			TaxCode                    string  `json:"taxCode"`
			Value                      string  `json:"value"`
			ShippingMethodDescription  string  `json:"shippingMethodDescription"`
		} `json:"shippingMethod"`
		ShippingGroupID string `json:"shippingGroupId"`
		Type            string `json:"type"`
		SubmittedDate   any    `json:"submittedDate"`
		PriceInfo       struct {
			Amount                  float64 `json:"amount"`
			Total                   float64 `json:"total"`
			LkpValExcludingFreeShip any     `json:"lkpValExcludingFreeShip"`
			Shipping                float64 `json:"shipping"`
			ShippingSurchargeValue  float64 `json:"shippingSurchargeValue"`
			Tax                     float64 `json:"tax"`
			SubTotal                float64 `json:"subTotal"`
			CurrencyCode            string  `json:"currencyCode"`
			TotalWithoutTax         float64 `json:"totalWithoutTax"`
		} `json:"priceInfo"`
		DiscountInfo struct {
			OrderDiscount    float64 `json:"orderDiscount"`
			DiscountDescList []any   `json:"discountDescList"`
			ShippingDiscount float64 `json:"shippingDiscount"`
		} `json:"discountInfo"`
		ShipOnDate      any   `json:"shipOnDate"`
		TrackingInfo    []any `json:"trackingInfo"`
		ActualShipDate  any   `json:"actualShipDate"`
		ShippingAddress struct {
			LastName    string `json:"lastName"`
			Country     string `json:"country"`
			Address3    any    `json:"address3"`
			Address2    any    `json:"address2"`
			City        string `json:"city"`
			Prefix      any    `json:"prefix"`
			Address1    string `json:"address1"`
			PostalCode  string `json:"postalCode"`
			CompanyName any    `json:"companyName"`
			JobTitle    any    `json:"jobTitle"`
			County      any    `json:"county"`
			Suffix      any    `json:"suffix"`
			FirstName   string `json:"firstName"`
			PhoneNumber string `json:"phoneNumber"`
			FaxNumber   string `json:"faxNumber"`
			Alias       any    `json:"alias"`
			MiddleName  any    `json:"middleName"`
			State       string `json:"state"`
			Email       string `json:"email"`
		} `json:"shippingAddress"`
		Items []struct {
			PrimaryThumbImageURL string  `json:"primaryThumbImageURL"`
			RawTotalPrice        float64 `json:"rawTotalPrice"`
			ReturnedQuantity     int     `json:"returnedQuantity"`
			DynamicProperties    []struct {
				Value any    `json:"value"`
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"dynamicProperties"`
			DisplayName            string  `json:"displayName"`
			ShippingSurchargeValue float64 `json:"shippingSurchargeValue"`
			AvailabilityDate       any     `json:"availabilityDate"`
			ExternalData           []any   `json:"externalData"`
			DiscountAmount         float64 `json:"discountAmount"`
			PreOrderQuantity       int     `json:"preOrderQuantity"`
			ExternalPrice          float64 `json:"externalPrice"`
			CommerceItemID         string  `json:"commerceItemId"`
			Price                  float64 `json:"price"`
			Variant                []struct {
				OptionName  string `json:"optionName"`
				OptionValue string `json:"optionValue"`
			} `json:"variant"`
			OnSale                bool    `json:"onSale"`
			PrimaryImageAltText   string  `json:"primaryImageAltText"`
			StateDetailsAsUser    string  `json:"stateDetailsAsUser"`
			CommerceID            string  `json:"commerceId"`
			UnitPrice             float64 `json:"unitPrice"`
			PrimaryImageTitle     string  `json:"primaryImageTitle"`
			Amount                float64 `json:"amount"`
			Quantity              int     `json:"quantity"`
			PointOfNoRevision     bool    `json:"pointOfNoRevision"`
			RelationshipType      string  `json:"relationshipType"`
			ProductID             string  `json:"productId"`
			SalePrice             float64 `json:"salePrice"`
			DetailedItemPriceInfo []struct {
				Discounted                 bool    `json:"discounted"`
				SecondaryCurrencyTaxAmount float64 `json:"secondaryCurrencyTaxAmount"`
				Amount                     float64 `json:"amount"`
				Quantity                   int     `json:"quantity"`
				ConfigurationDiscountShare float64 `json:"configurationDiscountShare"`
				Tax                        float64 `json:"tax"`
				OrderDiscountShare         float64 `json:"orderDiscountShare"`
				DetailedUnitPrice          float64 `json:"detailedUnitPrice"`
				CurrencyCode               string  `json:"currencyCode"`
			} `json:"detailedItemPriceInfo"`
			Active                bool   `json:"active"`
			ExternalPriceQuantity int    `json:"externalPriceQuantity"`
			CatRefID              string `json:"catRefId"`
			SkuProperties         []struct {
				Name         string `json:"name"`
				ID           string `json:"id"`
				Value        any    `json:"value"`
				PropertyType string `json:"propertyType"`
			} `json:"skuProperties"`
			DiscountInfo []any  `json:"discountInfo"`
			Route        string `json:"route"`
			SiteID       string `json:"siteId"`
			ShopperInput struct {
			} `json:"shopperInput"`
			Asset             bool    `json:"asset"`
			BackOrderQuantity int     `json:"backOrderQuantity"`
			ListPrice         float64 `json:"listPrice"`
			Status            string  `json:"status"`
		} `json:"items"`
		TrackingNumber any    `json:"trackingNumber"`
		Status         string `json:"status"`
	} `json:"shippingGroups"`
	LastModifiedDate       time.Time `json:"lastModifiedDate"`
	CreationSiteID         string    `json:"creationSiteId"`
	AllowAlternateCurrency bool      `json:"allowAlternateCurrency"`
	ApprovalSystemMessages []any     `json:"approvalSystemMessages"`
	PriceListGroup         struct {
		IsTaxIncluded bool   `json:"isTaxIncluded"`
		EndDate       any    `json:"endDate"`
		DisplayName   string `json:"displayName"`
		ListPriceList struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"listPriceList"`
		Active                     bool   `json:"active"`
		IsPointsBased              bool   `json:"isPointsBased"`
		Locale                     string `json:"locale"`
		ShippingSurchargePriceList struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"shippingSurchargePriceList"`
		Deleted            bool   `json:"deleted"`
		TaxCalculationType string `json:"taxCalculationType"`
		RepositoryID       string `json:"repositoryId"`
		SalePriceList      struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"salePriceList"`
		Currency struct {
			CurrencyType     any    `json:"currencyType"`
			Symbol           string `json:"symbol"`
			Deleted          bool   `json:"deleted"`
			DisplayName      string `json:"displayName"`
			RepositoryID     string `json:"repositoryId"`
			FractionalDigits int    `json:"fractionalDigits"`
			CurrencyCode     string `json:"currencyCode"`
			NumericCode      string `json:"numericCode"`
		} `json:"currency"`
		ID                 string `json:"id"`
		IncludeAllProducts bool   `json:"includeAllProducts"`
		StartDate          any    `json:"startDate"`
	} `json:"priceListGroup"`
	CreationDate   time.Time `json:"creationDate"`
	OrderProfileID string    `json:"orderProfileId"`
	OrderAction    string    `json:"orderAction"`
	DiscountInfo   struct {
		OrderDiscount         float64 `json:"orderDiscount"`
		OrderDiscountDescList []any   `json:"orderDiscountDescList"`
		ShippingDiscount      float64 `json:"shippingDiscount"`
	} `json:"discountInfo"`
	TrackingInfo   []any  `json:"trackingInfo"`
	OrderLocale    string `json:"orderLocale"`
	SiteID         string `json:"siteId"`
	BillingAddress struct {
		LastName    string `json:"lastName"`
		Country     string `json:"country"`
		Address3    string `json:"address3"`
		Address2    string `json:"address2"`
		City        string `json:"city"`
		Prefix      string `json:"prefix"`
		Address1    string `json:"address1"`
		PostalCode  string `json:"postalCode"`
		CompanyName string `json:"companyName"`
		JobTitle    string `json:"jobTitle"`
		County      string `json:"county"`
		Suffix      string `json:"suffix"`
		FirstName   string `json:"firstName"`
		PhoneNumber string `json:"phoneNumber"`
		FaxNumber   string `json:"faxNumber"`
		Alias       any    `json:"alias"`
		MiddleName  string `json:"middleName"`
		State       string `json:"state"`
		Email       string `json:"email"`
	} `json:"billingAddress"`
	Markers                      []any `json:"markers"`
	GiftWithPurchaseOrderMarkers []any `json:"giftWithPurchaseOrderMarkers"`
}

type EncryptedCard struct {
	Message   string `json:"message"`
	Errorcode int    `json:"errorcode"`
	Token     string `json:"token"`
}

type ProductResponse struct {
	TotalResults int `json:"totalResults"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	Links        []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Sort []struct {
		Property string `json:"property"`
		Order    string `json:"order"`
	} `json:"sort"`
	Category struct {
		LongDescription any    `json:"longDescription"`
		Route           string `json:"route"`
		CategoryImages  []any  `json:"categoryImages"`
		DisplayName     string `json:"displayName"`
		RepositoryID    string `json:"repositoryId"`
		Active          bool   `json:"active"`
		Description     string `json:"description"`
		ID              string `json:"id"`
	} `json:"category"`
	Items []struct {
		B2CExpertRatings any `json:"b2c_expertRatings"`
		B2CAge           any `json:"b2c_age"`
		OrderLimit       any `json:"orderLimit"`
		ListPrices       struct {
			DefaultPriceGroup float64 `json:"defaultPriceGroup"`
		} `json:"listPrices"`
		XVolumeUOM                   any    `json:"x_volumeUOM"`
		Type                         string `json:"type"`
		B2CTastingNotesBody          any    `json:"b2c_tastingNotesBody"`
		GcDescription                any    `json:"gc_description"`
		B2CLotteryPackageDescription any    `json:"b2c_lotteryPackageDescription"`
		Shippable                    bool   `json:"shippable"`
		B2CSizeSort                  string `json:"b2c_size_sort"`
		PrimaryImageAltText          string `json:"primaryImageAltText"`
		ID                           string `json:"id"`
		Brand                        string `json:"brand"`
		ParentCategories             []struct {
			RepositoryID          string `json:"repositoryId"`
			FixedParentCategories []struct {
				RepositoryID          string `json:"repositoryId"`
				FixedParentCategories []struct {
					RepositoryID          string `json:"repositoryId"`
					FixedParentCategories []struct {
						RepositoryID string `json:"repositoryId"`
					} `json:"fixedParentCategories"`
				} `json:"fixedParentCategories"`
			} `json:"fixedParentCategories"`
		} `json:"parentCategories"`
		Height                   any      `json:"height"`
		DefaultProductListingSku any      `json:"defaultProductListingSku"`
		Assetable                bool     `json:"assetable"`
		UnitOfMeasure            any      `json:"unitOfMeasure"`
		TargetAddOnProducts      []any    `json:"targetAddOnProducts"`
		B2CGlutenFree            string   `json:"b2c_glutenFree"`
		B2CDisableBopis          string   `json:"b2c_disableBopis"`
		B2CChairmansSelection    string   `json:"b2c_chairmansSelection"`
		SeoURLSlugDerived        string   `json:"seoUrlSlugDerived"`
		B2CRegion                any      `json:"b2c_region"`
		Active                   bool     `json:"active"`
		B2CUpc                   string   `json:"b2c_upc"`
		XSalesTaxIndicator       string   `json:"x_salesTaxIndicator"`
		ThumbImageURLs           []string `json:"thumbImageURLs"`
		B2CType                  string   `json:"b2c_type"`
		B2CPaResidencyOnly       string   `json:"b2c_paResidencyOnly"`
		B2COfferID               string   `json:"b2c_offerId"`
		B2CTastingNotes          string   `json:"b2c_tastingNotes"`
		B2CChairmansSpirits      string   `json:"b2c_chairmansSpirits"`
		Route                    string   `json:"route"`
		XVolumeOZ                string   `json:"x_volumeOZ"`
		RelatedArticles          []any    `json:"relatedArticles"`
		B2CProductIDNumber       string   `json:"b2c_productIdNumber"`
		MediumImageURLs          []string `json:"mediumImageURLs"`
		PrimarySourceImageURL    string   `json:"primarySourceImageURL"`
		SourceImageURLs          []string `json:"sourceImageURLs"`
		PrimaryThumbImageURL     string   `json:"primaryThumbImageURL"`
		DirectCatalogs           []any    `json:"directCatalogs"`
		Nonreturnable            bool     `json:"nonreturnable"`
		DisplayName              string   `json:"displayName"`
		B2CTaste                 any      `json:"b2c_taste"`
		B2CMostPopular           any      `json:"b2c_mostPopular"`
		PrimaryFullImageURL      string   `json:"primaryFullImageURL"`
		XFreightCost             any      `json:"x_freightCost"`
		ProductVariantOptions    []struct {
			VariantBasedDisplay     bool   `json:"variantBasedDisplay"`
			OptionID                string `json:"optionId"`
			ListingVariant          bool   `json:"listingVariant"`
			MapKeyPropertyAttribute string `json:"mapKeyPropertyAttribute"`
			OptionName              string `json:"optionName"`
			OptionValueMap          struct {
				Eaches int `json:"eaches"`
			} `json:"optionValueMap"`
		} `json:"productVariantOptions"`
		B2CExpertReviews          any    `json:"b2c_expertReviews"`
		PrimaryLargeImageURL      string `json:"primaryLargeImageURL"`
		B2CHighlyAllocatedProduct string `json:"b2c_highlyAllocatedProduct"`
		B2CInventoryAvailability  any    `json:"b2c_inventoryAvailability"`
		B2CVarietal               any    `json:"b2c_varietal"`
		SaleVolumePrices          struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"saleVolumePrices"`
		ChildSKUs []struct {
			DynamicPropertyMapLong struct {
				SkuB2CProductB2BLabel int `json:"sku-B2CProduct_b2b_label"`
			} `json:"dynamicPropertyMapLong"`
			BundleLinks     []any `json:"bundleLinks"`
			LargeImage      any   `json:"largeImage"`
			SmallImage      any   `json:"smallImage"`
			ListVolumePrice any   `json:"listVolumePrice"`
			OnlineOnly      bool  `json:"onlineOnly"`
			ListPrices      struct {
				DefaultPriceGroup float64 `json:"defaultPriceGroup"`
			} `json:"listPrices"`
			XUOM                  string `json:"x_uOM"`
			ConfigurationMetadata []any  `json:"configurationMetadata"`
			LargeImageURLs        []any  `json:"largeImageURLs"`
			ProductLine           any    `json:"productLine"`
			ListVolumePrices      struct {
				DefaultPriceGroup any `json:"defaultPriceGroup"`
			} `json:"listVolumePrices"`
			DerivedSalePriceFrom        string `json:"derivedSalePriceFrom"`
			Model                       any    `json:"model"`
			Barcode                     any    `json:"barcode"`
			XSupplier                   any    `json:"x_supplier"`
			SalePriceEndDate            any    `json:"salePriceEndDate"`
			Images                      []any  `json:"images"`
			UnitOfMeasure               any    `json:"unitOfMeasure"`
			PrimaryMediumImageURL       any    `json:"primaryMediumImageURL"`
			DynamicPropertyMapBigString struct {
			} `json:"dynamicPropertyMapBigString"`
			Active                bool   `json:"active"`
			ThumbImageURLs        []any  `json:"thumbImageURLs"`
			XCaseSize             string `json:"x_caseSize"`
			MediumImageURLs       []any  `json:"mediumImageURLs"`
			PrimarySourceImageURL any    `json:"primarySourceImageURL"`
			SourceImageURLs       []any  `json:"sourceImageURLs"`
			PrimarySmallImageURL  any    `json:"primarySmallImageURL"`
			ProductFamily         any    `json:"productFamily"`
			PrimaryThumbImageURL  any    `json:"primaryThumbImageURL"`
			Nonreturnable         bool   `json:"nonreturnable"`
			DisplayName           any    `json:"displayName"`
			SalePrices            struct {
				DefaultPriceGroup any `json:"defaultPriceGroup"`
			} `json:"salePrices"`
			B2BBottleLabel               string `json:"b2b_bottleLabel"`
			PrimaryFullImageURL          any    `json:"primaryFullImageURL"`
			XSearchableSKU               string `json:"x_searchableSKU"`
			B2BLabel                     string `json:"b2b_label"`
			ProductListingSku            any    `json:"productListingSku"`
			PrimaryLargeImageURL         any    `json:"primaryLargeImageURL"`
			DerivedOnlineOnly            bool   `json:"derivedOnlineOnly"`
			SmallImageURLs               []any  `json:"smallImageURLs"`
			DerivedShippingSurchargeFrom string `json:"derivedShippingSurchargeFrom"`
			ShippingSurcharges           struct {
				DefaultPriceGroup any `json:"defaultPriceGroup"`
			} `json:"shippingSurcharges"`
			ThumbnailImage   any `json:"thumbnailImage"`
			SaleVolumePrices struct {
				DefaultPriceGroup any `json:"defaultPriceGroup"`
			} `json:"saleVolumePrices"`
			SaleVolumePrice    any       `json:"saleVolumePrice"`
			SalePriceStartDate any       `json:"salePriceStartDate"`
			Quantity           any       `json:"quantity"`
			LastModifiedDate   time.Time `json:"lastModifiedDate"`
			SalePrice          any       `json:"salePrice"`
			FullImageURLs      []any     `json:"fullImageURLs"`
			VariantValuesOrder struct {
			} `json:"variantValuesOrder"`
			SoldAsPackage        bool    `json:"soldAsPackage"`
			ListingSKUID         any     `json:"listingSKUId"`
			RepositoryID         string  `json:"repositoryId"`
			DerivedListPriceFrom string  `json:"derivedListPriceFrom"`
			ShippingSurcharge    any     `json:"shippingSurcharge"`
			Configurable         bool    `json:"configurable"`
			ListPrice            float64 `json:"listPrice"`
		} `json:"childSKUs"`
		B2CCustomerRatingsFilterSplit any      `json:"b2c_customerRatingsFilterSplit"`
		SalePrice                     any      `json:"salePrice"`
		XVolume                       string   `json:"x_volume"`
		B2CSortWeighting              any      `json:"b2c_sortWeighting"`
		B2CScotchType                 any      `json:"b2c_scotchType"`
		B2CTastingNotesOakInfluence   any      `json:"b2c_tastingNotesOakInfluence"`
		B2CTastingNotesSweetness      any      `json:"b2c_tastingNotesSweetness"`
		XDeliveryFee                  any      `json:"x_deliveryFee"`
		NotForIndividualSale          bool     `json:"notForIndividualSale"`
		Width                         any      `json:"width"`
		B2CExpertRatingsFilter        any      `json:"b2c_expertRatingsFilter"`
		DerivedListPriceFrom          string   `json:"derivedListPriceFrom"`
		DefaultParentCategory         any      `json:"defaultParentCategory"`
		B2CRegionFilterSplit          any      `json:"b2c_regionFilterSplit"`
		ListPrice                     float64  `json:"listPrice"`
		B2CQuotedAtPrice              any      `json:"b2c_quotedAtPrice"`
		XAlcoholicOrNonalcoholic      string   `json:"x_alcoholicOrNonalcoholic"`
		ListVolumePrice               any      `json:"listVolumePrice"`
		B2CFeaturedFilterSplit        any      `json:"b2c_featuredFilterSplit"`
		ExcludeFromSitemap            bool     `json:"excludeFromSitemap"`
		B2CFreightIncludedSalePrice   float64  `json:"b2c_freightIncludedSalePrice"`
		RelatedProducts               any      `json:"relatedProducts"`
		OnlineOnly                    bool     `json:"onlineOnly"`
		LargeImageURLs                []string `json:"largeImageURLs"`
		B2CFreightIncludedListPrice   float64  `json:"b2c_freightIncludedListPrice"`
		ListVolumePrices              struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"listVolumePrices"`
		AddOnProducts         []any     `json:"addOnProducts"`
		DerivedSalePriceFrom  string    `json:"derivedSalePriceFrom"`
		XType                 string    `json:"x_type"`
		B2CNew                any       `json:"b2c_new"`
		B2CTopCustomerReviews any       `json:"b2c_topCustomerReviews"`
		PrimaryMediumImageURL string    `json:"primaryMediumImageURL"`
		B2CCountry            string    `json:"b2c_country"`
		Weight                any       `json:"weight"`
		CreationDate          time.Time `json:"creationDate"`
		ParentCategoryIDPath  string    `json:"parentCategoryIdPath"`
		XTypeDisplay          string    `json:"x_typeDisplay"`
		ParentCategory        struct {
			RepositoryID          string `json:"repositoryId"`
			FixedParentCategories []struct {
				RepositoryID          string `json:"repositoryId"`
				FixedParentCategories []struct {
					RepositoryID          string `json:"repositoryId"`
					FixedParentCategories []struct {
						RepositoryID string `json:"repositoryId"`
					} `json:"fixedParentCategories"`
				} `json:"fixedParentCategories"`
			} `json:"fixedParentCategories"`
		} `json:"parentCategory"`
		PrimarySmallImageURL            string `json:"primarySmallImageURL"`
		B2CSalePriceType                any    `json:"b2c_salePriceType"`
		B2CLimitPerOrder                any    `json:"b2c_limitPerOrder"`
		AvgCustRating                   any    `json:"avgCustRating"`
		B2CFeatured                     any    `json:"b2c_featured"`
		LongDescription                 any    `json:"longDescription"`
		B2CVintage                      any    `json:"b2c_vintage"`
		B2COnlineAvailable              string `json:"b2c_onlineAvailable"`
		B2CExpertRatingsFilterSplitSort any    `json:"b2c_expertRatingsFilterSplitSort"`
		Description                     any    `json:"description"`
		SalePrices                      struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"salePrices"`
		B2COnlineExclusive                string   `json:"b2c_onlineExclusive"`
		B2CSpecialOrderAddressShip        string   `json:"b2c_specialOrderAddressShip"`
		B2CFreightIncludedActivePrice     float64  `json:"b2c_freightIncludedActivePrice"`
		B2CLotteryProduct                 string   `json:"b2c_lotteryProduct"`
		B2CLotteryAvailabilityDescription any      `json:"b2c_lotteryAvailabilityDescription"`
		SmallImageURLs                    []string `json:"smallImageURLs"`
		B2BHasCase                        any      `json:"b2b_hasCase"`
		DerivedShippingSurchargeFrom      string   `json:"derivedShippingSurchargeFrom"`
		ShippingSurcharges                struct {
			DefaultPriceGroup any `json:"defaultPriceGroup"`
		} `json:"shippingSurcharges"`
		B2COrganic                        string    `json:"b2c_organic"`
		SaleVolumePrice                   any       `json:"saleVolumePrice"`
		PrimaryImageTitle                 string    `json:"primaryImageTitle"`
		B2CExpertRatingsFilterSplit       any       `json:"b2c_expertRatingsFilterSplit"`
		B2CSpecialOrderProduct            string    `json:"b2c_specialOrderProduct"`
		B2CClearance                      string    `json:"b2c_clearance"`
		RelatedMediaContent               []any     `json:"relatedMediaContent"`
		LastModifiedDate                  time.Time `json:"lastModifiedDate"`
		FullImageURLs                     []string  `json:"fullImageURLs"`
		B2CSize                           string    `json:"b2c_size"`
		Length                            any       `json:"length"`
		B2CProof                          string    `json:"b2c_proof"`
		DerivedDirectCatalogs             []any     `json:"derivedDirectCatalogs"`
		B2CFuturesProduct                 string    `json:"b2c_futuresProduct"`
		B2CLotteryRegistrationDescription any       `json:"b2c_lotteryRegistrationDescription"`
		B2CComingSoon                     string    `json:"b2c_comingSoon"`
		VariantValuesOrder                struct {
		} `json:"variantValuesOrder"`
		RepositoryID          string `json:"repositoryId"`
		ShippingSurcharge     any    `json:"shippingSurcharge"`
		ProductImagesMetadata []struct {
		} `json:"productImagesMetadata"`
		B2CMadeInPa  string `json:"b2c_madeInPa"`
		Configurable bool   `json:"configurable"`
	} `json:"items"`
}

type ShippingStruct struct {
	ShippingAddress `json:"shippingAddress"`
}

type ShippingAddress struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	CompanyName *string `json:"companyName"`
	Address1    string  `json:"address1"`
	Address2    string  `json:"address2"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	FaxNumber   string  `json:"faxNumber"`
	JobTitle    *string `json:"jobTitle"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
}

type BillingStruct struct {
	Payments []Payments `json:"payments"`
}

type Payments struct {
	BillingAddress   BillingAddress `json:"billingAddress"`
	CardNumber       string         `json:"cardNumber"`
	CardType         string         `json:"cardType"`
	ExpiryMonth      string         `json:"expiryMonth"`
	ExpiryYear       string         `json:"expiryYear"`
	NameOnCard       string         `json:"nameOnCard"`
	CustomProperties CustomCardProp `json:"customProperties"`
	Amount           float64        `json:"amount"`
	Type             string         `json:"type"`
}

type CustomCardProp struct {
	Token string `json:"token"`
	Type  string `json:"type"`
	Exp   string `json:"exp"`
}

type BillingAddress struct {
	LastName    string  `json:"lastName"`
	Country     string  `json:"country"`
	Address3    string  `json:"address3"`
	Address2    string  `json:"address2"`
	City        string  `json:"city"`
	Prefix      string  `json:"prefix"`
	Address1    string  `json:"address1"`
	JobTitle    string  `json:"jobTitle"`
	CompanyName string  `json:"companyName"`
	PostalCode  string  `json:"postalCode"`
	County      string  `json:"county"`
	Suffix      string  `json:"suffix"`
	FirstName   string  `json:"firstName"`
	PhoneNumber string  `json:"phoneNumber"`
	FaxNumber   string  `json:"faxNumber"`
	MiddleName  string  `json:"middleName"`
	State       string  `json:"state"`
	Email       string  `json:"email"`
	Company     *string `json:"company"`
}

type EncodeStruct struct {
	Account           string  `json:"account"`
	Source            string  `json:"source"`
	Encryptionhandler *string `json:"encryptionhandler"`
	Unique            bool    `json:"unique"`
	Expiry            *string `json:"expiry"`
	Cvv               string  `json:"cvv"`
}

type SkuDetails struct {
	DynamicPropertyMapLong struct {
		SkuB2CProductB2BLabel int `json:"sku-B2CProduct_b2b_label"`
	} `json:"dynamicPropertyMapLong"`
	BundleLinks     []any `json:"bundleLinks"`
	LargeImage      any   `json:"largeImage"`
	SmallImage      any   `json:"smallImage"`
	ListVolumePrice any   `json:"listVolumePrice"`
	ListPrices      struct {
		DefaultPriceGroup float64 `json:"defaultPriceGroup"`
	} `json:"listPrices"`
	XUOM                  string `json:"x_uOM"`
	ConfigurationMetadata []any  `json:"configurationMetadata"`
	LargeImageURLs        []any  `json:"largeImageURLs"`
	ProductLine           any    `json:"productLine"`
	ListVolumePrices      struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"listVolumePrices"`
	DerivedSalePriceFrom string `json:"derivedSalePriceFrom"`
	Model                any    `json:"model"`
	Links                []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Barcode                     any   `json:"barcode"`
	XSupplier                   any   `json:"x_supplier"`
	SalePriceEndDate            any   `json:"salePriceEndDate"`
	Images                      []any `json:"images"`
	UnitOfMeasure               any   `json:"unitOfMeasure"`
	PrimaryMediumImageURL       any   `json:"primaryMediumImageURL"`
	DynamicPropertyMapBigString struct {
	} `json:"dynamicPropertyMapBigString"`
	Active         bool   `json:"active"`
	ThumbImageURLs []any  `json:"thumbImageURLs"`
	XCaseSize      string `json:"x_caseSize"`
	ParentProducts []struct {
		B2CExpertRatings             any    `json:"b2c_expertRatings"`
		B2CAge                       any    `json:"b2c_age"`
		OrderLimit                   any    `json:"orderLimit"`
		ListPrices                   any    `json:"listPrices"`
		XVolumeUOM                   any    `json:"x_volumeUOM"`
		Type                         string `json:"type"`
		B2CTastingNotesBody          any    `json:"b2c_tastingNotesBody"`
		GcDescription                any    `json:"gc_description"`
		B2CLotteryPackageDescription any    `json:"b2c_lotteryPackageDescription"`
		Shippable                    bool   `json:"shippable"`
		B2CSizeSort                  string `json:"b2c_size_sort"`
		PrimaryImageAltText          string `json:"primaryImageAltText"`
		ID                           string `json:"id"`
		Brand                        string `json:"brand"`
		ParentCategories             []struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"parentCategories"`
		Height                             any      `json:"height"`
		DefaultProductListingSku           any      `json:"defaultProductListingSku"`
		Assetable                          bool     `json:"assetable"`
		SecondaryCurrencyShippingSurcharge any      `json:"secondaryCurrencyShippingSurcharge"`
		UnitOfMeasure                      any      `json:"unitOfMeasure"`
		TargetAddOnProducts                []any    `json:"targetAddOnProducts"`
		B2CGlutenFree                      string   `json:"b2c_glutenFree"`
		B2CDisableBopis                    string   `json:"b2c_disableBopis"`
		B2CChairmansSelection              string   `json:"b2c_chairmansSelection"`
		SeoURLSlugDerived                  string   `json:"seoUrlSlugDerived"`
		B2CRegion                          any      `json:"b2c_region"`
		Active                             bool     `json:"active"`
		B2CUpc                             string   `json:"b2c_upc"`
		XSalesTaxIndicator                 string   `json:"x_salesTaxIndicator"`
		ThumbImageURLs                     []string `json:"thumbImageURLs"`
		B2CType                            string   `json:"b2c_type"`
		B2CPaResidencyOnly                 string   `json:"b2c_paResidencyOnly"`
		B2COfferID                         any      `json:"b2c_offerId"`
		B2CTastingNotes                    string   `json:"b2c_tastingNotes"`
		B2CChairmansSpirits                string   `json:"b2c_chairmansSpirits"`
		Route                              string   `json:"route"`
		XVolumeOZ                          string   `json:"x_volumeOZ"`
		RelatedArticles                    []any    `json:"relatedArticles"`
		B2CProductIDNumber                 string   `json:"b2c_productIdNumber"`
		MediumImageURLs                    []string `json:"mediumImageURLs"`
		PrimarySourceImageURL              string   `json:"primarySourceImageURL"`
		SourceImageURLs                    []string `json:"sourceImageURLs"`
		PrimaryThumbImageURL               string   `json:"primaryThumbImageURL"`
		DirectCatalogs                     []any    `json:"directCatalogs"`
		Nonreturnable                      bool     `json:"nonreturnable"`
		DisplayName                        string   `json:"displayName"`
		B2CTaste                           any      `json:"b2c_taste"`
		B2CMostPopular                     any      `json:"b2c_mostPopular"`
		PrimaryFullImageURL                string   `json:"primaryFullImageURL"`
		XFreightCost                       any      `json:"x_freightCost"`
		ProductVariantOptions              any      `json:"productVariantOptions"`
		ParentCategoriesAncestorCategories []struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"parentCategoriesAncestorCategories"`
		B2CExpertReviews                any      `json:"b2c_expertReviews"`
		PrimaryLargeImageURL            string   `json:"primaryLargeImageURL"`
		B2CHighlyAllocatedProduct       string   `json:"b2c_highlyAllocatedProduct"`
		B2CInventoryAvailability        any      `json:"b2c_inventoryAvailability"`
		B2CVarietal                     any      `json:"b2c_varietal"`
		IsOnSale                        any      `json:"isOnSale"`
		SaleVolumePrices                any      `json:"saleVolumePrices"`
		B2CCustomerRatingsFilterSplit   string   `json:"b2c_customerRatingsFilterSplit"`
		SalePrice                       any      `json:"salePrice"`
		XVolume                         string   `json:"x_volume"`
		B2CSortWeighting                any      `json:"b2c_sortWeighting"`
		B2CScotchType                   any      `json:"b2c_scotchType"`
		B2CTastingNotesOakInfluence     any      `json:"b2c_tastingNotesOakInfluence"`
		B2CTastingNotesSweetness        any      `json:"b2c_tastingNotesSweetness"`
		XDeliveryFee                    any      `json:"x_deliveryFee"`
		NotForIndividualSale            bool     `json:"notForIndividualSale"`
		Width                           any      `json:"width"`
		B2CExpertRatingsFilter          any      `json:"b2c_expertRatingsFilter"`
		DerivedListPriceFrom            any      `json:"derivedListPriceFrom"`
		DefaultParentCategory           any      `json:"defaultParentCategory"`
		B2CRegionFilterSplit            any      `json:"b2c_regionFilterSplit"`
		PriceRange                      any      `json:"priceRange"`
		ListPrice                       any      `json:"listPrice"`
		B2CQuotedAtPrice                any      `json:"b2c_quotedAtPrice"`
		XAlcoholicOrNonalcoholic        string   `json:"x_alcoholicOrNonalcoholic"`
		AncestorCategoriesForFullDeploy []any    `json:"ancestorCategoriesForFullDeploy"`
		ListVolumePrice                 any      `json:"listVolumePrice"`
		B2CFeaturedFilterSplit          any      `json:"b2c_featuredFilterSplit"`
		ExcludeFromSitemap              bool     `json:"excludeFromSitemap"`
		B2CFreightIncludedSalePrice     float64  `json:"b2c_freightIncludedSalePrice"`
		RelatedProducts                 any      `json:"relatedProducts"`
		OnlineOnly                      bool     `json:"onlineOnly"`
		LargeImageURLs                  []string `json:"largeImageURLs"`
		B2CFreightIncludedListPrice     float64  `json:"b2c_freightIncludedListPrice"`
		ListVolumePrices                any      `json:"listVolumePrices"`
		AddOnProducts                   []any    `json:"addOnProducts"`
		DerivedSalePriceFrom            any      `json:"derivedSalePriceFrom"`
		XType                           string   `json:"x_type"`
		B2CNew                          any      `json:"b2c_new"`
		B2CTopCustomerReviews           string   `json:"b2c_topCustomerReviews"`
		PrimaryMediumImageURL           string   `json:"primaryMediumImageURL"`
		B2CCountry                      string   `json:"b2c_country"`
		Weight                          any      `json:"weight"`
		ParentCategoryIDPath            string   `json:"parentCategoryIdPath"`
		XTypeDisplay                    string   `json:"x_typeDisplay"`
		ParentCategory                  struct {
			RepositoryID string `json:"repositoryId"`
		} `json:"parentCategory"`
		PrimarySmallImageURL              string    `json:"primarySmallImageURL"`
		B2CSalePriceType                  any       `json:"b2c_salePriceType"`
		B2CLimitPerOrder                  any       `json:"b2c_limitPerOrder"`
		AvgCustRating                     any       `json:"avgCustRating"`
		B2CFeatured                       any       `json:"b2c_featured"`
		LongDescription                   any       `json:"longDescription"`
		B2CVintage                        any       `json:"b2c_vintage"`
		B2COnlineAvailable                string    `json:"b2c_onlineAvailable"`
		B2CExpertRatingsFilterSplitSort   any       `json:"b2c_expertRatingsFilterSplitSort"`
		Description                       any       `json:"description"`
		SalePrices                        any       `json:"salePrices"`
		B2COnlineExclusive                string    `json:"b2c_onlineExclusive"`
		B2CSpecialOrderAddressShip        string    `json:"b2c_specialOrderAddressShip"`
		B2CFreightIncludedActivePrice     float64   `json:"b2c_freightIncludedActivePrice"`
		B2CLotteryProduct                 string    `json:"b2c_lotteryProduct"`
		B2CLotteryAvailabilityDescription any       `json:"b2c_lotteryAvailabilityDescription"`
		SmallImageURLs                    []string  `json:"smallImageURLs"`
		B2BHasCase                        any       `json:"b2b_hasCase"`
		DerivedShippingSurchargeFrom      any       `json:"derivedShippingSurchargeFrom"`
		ShippingSurcharges                any       `json:"shippingSurcharges"`
		B2COrganic                        string    `json:"b2c_organic"`
		SaleVolumePrice                   any       `json:"saleVolumePrice"`
		PrimaryImageTitle                 string    `json:"primaryImageTitle"`
		B2CExpertRatingsFilterSplit       any       `json:"b2c_expertRatingsFilterSplit"`
		B2CSpecialOrderProduct            string    `json:"b2c_specialOrderProduct"`
		B2CClearance                      string    `json:"b2c_clearance"`
		RelatedMediaContent               []any     `json:"relatedMediaContent"`
		LastModifiedDate                  time.Time `json:"lastModifiedDate"`
		FullImageURLs                     []string  `json:"fullImageURLs"`
		B2CSize                           string    `json:"b2c_size"`
		Length                            any       `json:"length"`
		B2CProof                          string    `json:"b2c_proof"`
		DerivedDirectCatalogs             []any     `json:"derivedDirectCatalogs"`
		B2CFuturesProduct                 string    `json:"b2c_futuresProduct"`
		B2CLotteryRegistrationDescription any       `json:"b2c_lotteryRegistrationDescription"`
		B2CComingSoon                     string    `json:"b2c_comingSoon"`
		VariantValuesOrder                any       `json:"variantValuesOrder"`
		WasPriceRange                     any       `json:"wasPriceRange"`
		RepositoryID                      string    `json:"repositoryId"`
		ShippingSurcharge                 any       `json:"shippingSurcharge"`
		FractionalQuantitiesAllowed       bool      `json:"fractionalQuantitiesAllowed"`
		ProductImagesMetadata             []struct {
		} `json:"productImagesMetadata"`
		B2CMadeInPa  string `json:"b2c_madeInPa"`
		Configurable bool   `json:"configurable"`
	} `json:"parentProducts"`
	MediumImageURLs       []any `json:"mediumImageURLs"`
	PrimarySourceImageURL any   `json:"primarySourceImageURL"`
	SourceImageURLs       []any `json:"sourceImageURLs"`
	PrimarySmallImageURL  any   `json:"primarySmallImageURL"`
	ProductFamily         any   `json:"productFamily"`
	PrimaryThumbImageURL  any   `json:"primaryThumbImageURL"`
	Nonreturnable         bool  `json:"nonreturnable"`
	DisplayName           any   `json:"displayName"`
	SalePrices            struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"salePrices"`
	B2BBottleLabel        string `json:"b2b_bottleLabel"`
	PrimaryFullImageURL   any    `json:"primaryFullImageURL"`
	XSearchableSKU        string `json:"x_searchableSKU"`
	B2BLabel              string `json:"b2b_label"`
	ProductVariantOptions [][]struct {
		OptionValue string `json:"optionValue"`
		OptionName  string `json:"optionName"`
	} `json:"productVariantOptions"`
	ProductListingSku            any    `json:"productListingSku"`
	PrimaryLargeImageURL         any    `json:"primaryLargeImageURL"`
	DerivedOnlineOnly            bool   `json:"derivedOnlineOnly"`
	SmallImageURLs               []any  `json:"smallImageURLs"`
	DerivedShippingSurchargeFrom string `json:"derivedShippingSurchargeFrom"`
	ShippingSurcharges           struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"shippingSurcharges"`
	ThumbnailImage   any `json:"thumbnailImage"`
	SaleVolumePrices struct {
		DefaultPriceGroup any `json:"defaultPriceGroup"`
	} `json:"saleVolumePrices"`
	SaleVolumePrice    any       `json:"saleVolumePrice"`
	SalePriceStartDate any       `json:"salePriceStartDate"`
	Quantity           any       `json:"quantity"`
	LastModifiedDate   time.Time `json:"lastModifiedDate"`
	SalePrice          any       `json:"salePrice"`
	FullImageURLs      []any     `json:"fullImageURLs"`
	VariantValuesOrder struct {
	} `json:"variantValuesOrder"`
	SoldAsPackage               bool    `json:"soldAsPackage"`
	ListingSKUID                any     `json:"listingSKUId"`
	RepositoryID                string  `json:"repositoryId"`
	DerivedListPriceFrom        string  `json:"derivedListPriceFrom"`
	ShippingSurcharge           any     `json:"shippingSurcharge"`
	FractionalQuantitiesAllowed bool    `json:"fractionalQuantitiesAllowed"`
	Configurable                bool    `json:"configurable"`
	ListPrice                   float64 `json:"listPrice"`
}
