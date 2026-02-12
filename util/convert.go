package util

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ConvertDate(date []int) string {

	year := strconv.Itoa(int(date[0]))
	var month string
	day := strconv.Itoa(int(date[2]))

	switch date[1] {
	case 1:
		month = "Jan"
	case 2:
		month = "Feb"
	case 3:
		month = "Mar"
	case 4:
		month = "April"
	case 5:
		month = "May"
	case 6:
		month = "June"
	case 7:
		month = "July"
	case 8:
		month = "August"
	case 9:
		month = "September"
	case 10:
		month = "October"
	case 11:
		month = "November"
	case 12:
		month = "December"
	}

	str := fmt.Sprintf("%s %s, %s", month, day, year)
	return str

}

func GetCheckoutDates(availDates [][]int, selection []int) (string, string, string, string) {

	end := selection[len(selection)-1]
	beginDate := availDates[selection[0]]
	endDate := availDates[end]

	checkInDate := fmt.Sprintf("%d-%02d-%02d", int(beginDate[0]), int(beginDate[1]), int(beginDate[2]))
	checkOutDate := fmt.Sprintf("%d-%02d-%02d", int(endDate[0]), int(endDate[1]), int(endDate[2]+1))

	CheckoutDateIn := fmt.Sprintf("%d-%d-%d", int(beginDate[1]), int(beginDate[2]), int(beginDate[0])%100)
	CheckoutDateOut := fmt.Sprintf("%d-%d-%d", int(endDate[1]), int(endDate[2]+1), int(endDate[0])%100)

	return checkInDate, checkOutDate, CheckoutDateIn, CheckoutDateOut

}

func ConvertProfileDates(checkIn string, checkOut string) (string, string) {

	beginDate := strings.Split(checkIn, "-")
	endDate := strings.Split(checkOut, "-")
	var beginInt []int
	var endInt []int

	for i := range beginDate {
		s, _ := strconv.Atoi(beginDate[i])
		v, _ := strconv.Atoi(endDate[i])
		beginInt = append(beginInt, s)
		endInt = append(endInt, v)
	}
	checkoutInDate := fmt.Sprintf("%d-%d-%d", int(beginInt[1]), int(beginInt[2]), int(beginInt[0])%100)
	checkoutOutDate := fmt.Sprintf("%d-%d-%d", int(endInt[1]), int(endInt[2]), int(endInt[0])%100)
	return checkoutInDate, checkoutOutDate

}

func SaveHTML(html io.ReadCloser) {

	path := "data/submithtml/"

	file, err := os.Create(path + strconv.Itoa(int(time.Now().Unix())) + ".html")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, html)
	if err != nil {
		fmt.Println("Error copying response to file:", err)
		return
	}

}
