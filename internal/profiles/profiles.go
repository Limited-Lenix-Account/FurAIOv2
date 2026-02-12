package profiles

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadUserProfiles() []Profile {

	GetProfiles()

	var prf []Profile

	fmt.Println("Reading Profiles")
	file, err := os.Open("data/csvProfilesC.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&prf)
	if err != nil {
		log.Fatal(err)
	}

	return prf
}

func readCsvFile() [][]string {

	csvF, err := os.Open("internal/profiles/profiles.csv")
	if err != nil {
		fmt.Println("Error opening profile file")
	}
	csvR := csv.NewReader(csvF)
	records, err := csvR.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	return records

}

func writeCsvProfiles(profiles []Profile) {

	file, _ := os.Create("data/csvProfilesC.json")

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	encoder.Encode(profiles)

}

func GetProfiles() []Profile {

	r := readCsvFile()

	var profW []Profile

	for i := range r {
		// Ignore header row of csv
		if i == 0 {
			continue
		}

		b := BookingInformation{
			CheckInDate:  r[i][19],
			CheckOutDate: r[i][20],
		}

		p := PaymentInformation{
			CardNumber: r[i][14],
			ExpMon:     r[i][15],
			ExpYr:      r[i][16],
			CardType:   r[i][17],
			CVV:        r[i][18],
		}

		u := UserInformation{
			FirstName:          r[i][5],
			LastName:           r[i][6],
			AddressLine1:       r[i][7],
			AddressLine2:       r[i][8],
			City:               r[i][9],
			State:              r[i][10],
			Zip:                r[i][11],
			PhoneNumber:        r[i][12],
			Email:              r[i][13],
			PaymentInformation: p,
			BookingInformation: b,
		}

		pr := Profile{
			Site:            r[i][0],
			ProfileName:     r[i][1],
			HotelKeywords:   r[i][2],
			RoomKeywords:    r[i][3],
			LiquorKeywords:  r[i][4],
			UserInformation: u,
		}
		profW = append(profW, pr)
	}

	writeCsvProfiles(profW)

	return profW

}
