package pa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var SPIRIT_CATEGORY_IDS = []string{
	"120", // Spirits
	"121", // Brandy & Cognac
	"122", // Unflavored Brandy
	"123", // Flavored Brandy
	"124", // Cognac
	"125", // Cordials & Liqueurs
}

// /inventories?ids=000004405&type=product&locationIds=5161,5105
// Can be used for Pickup Locations
// type=product shows product name
// if locationIds is blank, stock defaults to shipping

// /ccstore/v1/products?q=displayName%20co%20%22hennessy%22
// Able to do search queries q=displayName co "hennessy"

type Product struct {
	Name            string
	ProductID       string
	LastModified    time.Time
	Size            string
	Unit            any
	ComingSoon      string
	Volume          string
	Image           string
	AvailableOnline string
	URL             string
}

type Keywords struct {
	Positive []string
	Negative []string
}

type KeywordGroup struct {
	Name     string
	Keywords Keywords
}

type ProductList []Product

// If a SKU is used as a keyword it will directly look up that sku whereas a keyword pair will have to be searched
// through the entire site.
func SiteMonitor() {

	// Get All whiskey products the 1st time
	kws := Keywords{
		Positive: []string{"Michters"},
		Negative: []string{},
	}

	// for {
	// 	oldProd := GetAllSpirits(kws.Positive)
	// 	time.Sleep(15 * time.Second)
	// 	newProd := GetAllSpirits(kws.Positive)

	// }

	oldProds := GetAllSpirits(kws.Positive)

	for i := range oldProds {
		fmt.Println(oldProds[i].Name, oldProds[i].URL)
	}

}

// TODO - Move StartMonitor to pa.go
// Monitor Function used in Checkout Flow
func (e *PaTask) StartMonitor() {

	//If keyword is an item sku
	if e.Profile.LiquorKeywords[:2] == "00" || e.Profile.LiquorKeywords[:2] == "10" {
		p := GetSkuDetails(e.Profile.LiquorKeywords)

		e.ProductID = p.ProductID
		e.ProductName = p.Name
		e.ProductImage = p.Image
		e.ProductVol = p.Volume

		e.MonitorSKUs = append(e.MonitorSKUs, e.ProductID)
		e.UpdateStatus(CheckStock)
		e.CheckoutStartTime = time.Now()
		return
	}

	// if there are +/- keywords instead
	var posKws []string
	var negKws []string

	s := strings.Split(e.Profile.LiquorKeywords, ",")

	for _, v := range s {
		if string(v[0]) == "+" {
			posKws = append(posKws, v[1:])
		} else {
			negKws = append(negKws, v[1:])
		}

	}

	kws := Keywords{
		Positive: posKws,
		Negative: negKws,
	}

	kwsg := KeywordGroup{
		Name:     "Task",
		Keywords: kws,
	}

	prods := GetAllSpirits(posKws)
	pl := prods.SearchProducts(kwsg)

	//Checks if there is an active page for the sku
	if len(pl) > 0 {
		e.ProductID = pl[0].ProductID
		e.ProductName = pl[0].Name
		e.ProductImage = pl[0].Image
		e.ProductVol = pl[0].Volume

		e.UpdateStatus(AddToCart)
		e.CheckoutStartTime = time.Now()
	} else {
		fmt.Println("No Results Found... retrying")
		time.Sleep(3 * time.Second)
	}
}

func GetAllSpirits(posKws []string) ProductList {

	var prods []Product

	prod, ps := GetProdsFromQuery(0, posKws)

	fmt.Println(prod.TotalResults)

	pages := (prod.TotalResults / 250) + 1

	if pages > 1 {
		var wg sync.WaitGroup
		wg.Add(pages)
		for i := 0; i < pages; i++ {
			go func(offset int) {
				defer wg.Done()
				_, l := GetProdsFromQuery(offset, posKws)

				prods = append(prods, l...)
			}(i * 250)
		}
		wg.Wait()
	} else {
		prods = append(prods, ps...)
	}

	return prods
}

func GetProdsFromQuery(page int, posKws []string) (*ProductResponse, []Product) {
	client := &http.Client{}

	// q=displayName%20co%20"Hennessy"

	// q = (displayName co "Hennessy" and displayName co "XO"...)

	// q = (displayName co "hennessy" and displayName co "XO")

	// q = (displayName co "Hennessy" and displayName co "XO")

	query := `(displayName%20co%20"` + strings.Join(posKws, `"%20and%20displayName%20co%20"`) + `")`
	// fmt.Println(query)

	urlString := "https://www.finewineandgoodspirits.com/ccstore/v1/products?includeChildren=true&offset=" + strconv.Itoa(page) + "&q=" + query
	// fmt.Printf("urlString: %v\n", urlString)
	if len(posKws) == 0 {
		urlString = "https://www.finewineandgoodspirits.com/ccstore/v1/products?categoryId=120&includeChildren=true&offset=" + strconv.Itoa(page)
	}

	// fmt.Println(urlString)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.finewineandgoodspirits.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	// req.Header.Set("Cookie", "_gcl_au=1.1.2111076155.1712769182; GSIDyk8q9gLSz40L=3a46ef7f-0269-4992-b81c-aa23a55257fd; STSID148379=e6f84471-0f49-4ba7-bf3f-0d50fcd8f76c; BVBRANDID=a82f17bc-93fa-4d18-94af-62fda45fc91e; _ga=GA1.1.2050355724.1712769182; __pdst=1b26f91811f54e848b091ce7573f7f59; __adroll_fpc=d8f13915972a0872f73533a05cdab84c-1712769183244; _pin_unauth=dWlkPU5HTTNZakV3TUdJdE16UmtaaTAwWmpaaUxUZ3paR1V0WW1Fek9UZGhaREExTTJVMA; _fbp=fb.1.1712769211703.185896143; ltkpopup-suppression-8e9eb133-da66-408a-880e-6f4844cae736=1; ccstoreroute=e377ede55d92f51d0b11dd3801f92e6b|cc2d3787e8999f4a634b9d592c9e2268; osfliveuiroute=7aa2295821121d5d9ba1a190c8adfaf6|a271690a033ecd4ba18a69565eca7c19; storePriceListGroupId=defaultPriceGroup; ltkSubscriber-Footer=eyJsdGtDaGFubmVsIjoiZW1haWwiLCJsdGtUcmlnZ2VyIjoibG9hZCIsImx0a0VtYWlsIjoiIn0%3D; AGEVERIFY=Over21; JSESSIONID=EasICFFJsGbk1aLUSMyECR0xpJHXHEdxf6siMc79GGntH098nNJ0!1108060648; BVBRANDSID=ead82b8b-3c0f-4544-9fd4-b5057e87104f; __ZEHIC7616=1713827487; _vuid=dd9375d6-68f8-46dd-8363-79cca64652ec; xva8383328c1PRD_siteUS=240b89f4%3A18f07efd1da%3A6139-4094298633; xda8383328c1PRD_siteUS=10C7QZZku5nlm-NsqcjIRupXsbiPr4LkQgg87waqVvXNTBQ9F92%3A1712769183705%3A1713812486905%3A1713827489944%3A16%3A16; __z_a=2181996688100876420110087; ltkpopup-session-depth=4-6; __ar_v4=YEV7UUOV4NHJNEIGTN5CUO%3A20240410%3A41%7CH422MU2HOBC4HLZGUQMCAG%3A20240410%3A41; __ZEHIC8348=1713827513; ak_bmsc=2C58347D51D1A44BA895A8119E2B9473~000000000000000000000000000000~YAAQCo0hF4/dhAaPAQAASD4bCBcvmt+G3gn+hyU5NUDq6L1wNId3JbalRvWF8leE98unzH7F1Uo+OXYEQo3N56hvEuHNoJMff4p6QLULFLS3phAQt+8NLKEqKmcTWsG8PSo+UaQMdOHYjKIa8AtpFIrpumCUcU9CRT0kKqVmElntOuBDhM1290eI2ER0kohool0+02wu/U/1BSERcPsJvW4SmU6lfsF3TUIWJsXKefJEdo3wEXpMl0yQQ+RTNxOE0S6v/zghULHpKrwkjagi97wiP2esEP311S4/VnY3UfdgMQzilt5OAuAGbCqK8fXooo1kaA6ZTWucPnrNDsCpGMcJpYMMplEp83gih+G+mCVzfcTTPbWmrXHVc6rueC83WxR8NNuXIkSQi3M=; bm_sz=F761E7F6814BBC3AEC764CB5747158FF~YAAQJo0hF9z1UNCOAQAAj+UdCBcMOu7/fOxwUuJvuvj0rNKs619SOO2O4WhiN9kqsk4+/HNJO1h9RMc+3DqZt0tjxpXZhTKe54aJKnhk7nWcPAIlduitIa8rfjZuoMPaXdjF1M4pvnfHniqalJGa43stCqgUQgnGYO4soIhs0FQEA3i3Osb3gTZ38ae54tA/SViOwzkH45XEf07ELAhDB4VPEBCtdgWb+K4bc6jeJLmUKLScbmzdHLXeYDc+GClVL4U8G5wKK1sATeycApNirKaj/yPiqFr+o8JbCCtMqMDPWyhCbha0tgwH3w00JVxApIsc4mAqwVDJ1QnUgaT0wKU9HP19DIqXQqRWEfMaM0EoE6NM0Cyrd4oieMIo/Gy8Z6xMBlo4VynjvVxV2X1f34W8dxIH+2Zefz6HN/p0WhI/P4wutIqMBI7LIUPqUepuimk0GLnWO3efcxq6EGKJzhnLswXdsOvZDtuL+fX5V217G7Zuv2qZxe0MvZtFNpcyt4/54ZY8AEDY/9SZhEL9LlnS07dMGL9rA7CAzQOrVw/e46U6pYsC2wK0CmfFicwvEFp6C1k6im+Bx59E+3k=~3749175~3159618; _ga_Y3ZSL11TM0=GS1.1.1713827488.14.1.1713828183.51.0.0; _abck=0CFD734DABA18D369949E71FA27D3CC8~-1~YAAQCo0hF7Q9hQaPAQAAz8QeCAuU+Yy0PYSAPBOqRAPC6HV1TsTH49ZNTpCO1mR4L3MBH12ikCOYhe+4hSoxDXcuq5+cMoJuqm1aeM1x+zom4r1Fh2teysslmrs//8Mjevmgdb84T0LFg8aC7n/NZqcKGBiFZIAkFlVHvmnhtb77krl5lA2B5oIhWCZ+msR2XRQYs1z+Tao6c9ekRfsvd1QUYLbAwwfr802T+YSaUVHUxr0fQmMvYhQMPS0WVvvtvRxBPbUl7Z9SwK5DF9CPyWz/B3+ZOk1xY09zKX22tVEd61aT5WJHaZbrttpuqFTLv5cFCBCyG8n1Vj8RMe8HNl6RwBzj/iEGv+OpsMHbggQYMnnD19TBomUVTTlJ1ekPDqfkqQqDhdfaa6FmBNCeDeF6AICIXmctTFA=~-1~-1~-1; bm_sv=3DBCF9CA7A4F450A599B6C347CB42E54~YAAQCo0hF7U9hQaPAQAAz8QeCBcQoMVonmwPyB+PjMH/juijOWPeSro0n3CTK03++Sc5T0xxUGNWnW+pqQQTv4R3Lqs6hj/JSl3GsVBac2yKbWpG453LeJZEOzr8TUVk7nNhLMX51EFEKVwfc6Zm32MfRlEX2WCXa2467SwKfjt7aDC5BhmwHb0+nbqwJCm+PsaLnvbyOjvWaJNKTOQj6q/pX1kv3/GB2kFlltNFKeAkHmIpDkEm46IFcB8GmREjL6ESh1uyxQHfSG4zhrUbz4Y=~1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prod ProductResponse
	var prods []Product
	err = json.Unmarshal(body, &prod)
	if err != nil {
		fmt.Println("Error Unmarshalling products")
	}

	for _, v := range prod.Items {
		p := Product{
			Name:            v.DisplayName,
			ProductID:       v.ID,
			LastModified:    v.LastModifiedDate,
			Size:            v.B2CSize,
			Unit:            v.UnitOfMeasure,
			ComingSoon:      v.B2CComingSoon,
			Volume:          v.B2CSize,
			Image:           "https://www.finewineandgoodspirits.com" + v.PrimarySourceImageURL,
			AvailableOnline: v.B2COnlineExclusive,
			URL:             "https://www.finewineandgoodspirits.com" + v.Route,
		}

		prods = append(prods, p)
	}
	// fmt.Println(prods[0])
	return &prod, prods
}

// Returns a []Product of items that match the keyword group
func (p *ProductList) SearchProducts(kws KeywordGroup) []Product {

	var prods []Product

	for _, v := range *p {
		if containsKeywords(v.Name, kws.Keywords.Positive, kws.Keywords.Negative) {
			// fmt.Println(v.Name, v.ProductID)
			prods = append(prods, v)
		}
	}

	return prods

}

// Prints products by most recently modified
func (p *ProductList) SearchProductsByDate() {

	sort.Slice(*p, func(i, j int) bool {
		// This will sort in ascending order. For descending order, use '(*p)[j].LastModified.Before((*p)[i].LastModified)'
		return (*p)[j].LastModified.Before((*p)[i].LastModified)
	})

	for i := 0; i < 5; i++ {
		fmt.Println((*p)[i].Name)
	}
}

// Prints products which are labed as Coming Soon
func (p *ProductList) SearchProductsComingSoon() {
	fmt.Println("Coming Soon Products")
	for _, v := range *p {
		if v.ComingSoon == "Y" {
			fmt.Println(v.Name)
		}
	}

	fmt.Println()
}

// Uses a sku to directly send a request to get full sku information
func GetSkuDetails(sku string) Product {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.finewineandgoodspirits.com/ccstore/v1/skus/"+sku, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="99", "Chromium";v="124"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", body)

	var prod SkuDetails
	err = json.Unmarshal(body, &prod)
	if err != nil {
		fmt.Println("Error unmarshalling SKU response", err)
	}

	v := prod.ParentProducts[0]

	p := Product{
		Name:         v.DisplayName,
		ProductID:    v.ID,
		LastModified: v.LastModifiedDate,
		Size:         v.B2CSize,
		Unit:         v.UnitOfMeasure,
		ComingSoon:   v.B2CComingSoon,
		Volume:       v.B2CSize,
		Image:        "https://www.finewineandgoodspirits.com" + v.PrimarySourceImageURL,
	}
	return p
}

// Helper function that returns a bool if a product name matches the keyword group
func containsKeywords(s string, positiveKeywords []string, negativeKeywords []string) bool {
	// Check if string contains all positive keywords
	for _, keyword := range positiveKeywords {
		if !strings.Contains(s, keyword) {
			return false
		}
	}

	// Check if string contains any negative keywords
	for _, keyword := range negativeKeywords {
		if strings.Contains(s, keyword) {
			return false
		}
	}
	return true
}

func CompareProducts(oldProd, newProd *ProductList) {

	// for _, v := range *newProd {

	// }

}
