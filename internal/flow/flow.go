package flow

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func StartFlow(task *UserTask) {
	//Create HTTP Thing + Cookie Jar
	fmt.Println("Starting Session...")
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error making session!")
	}
	client := &http.Client{
		Jar: jar,
	}
	task.Client.client = client
	var optsTitle []string
	var hotelTitle []string
	var roomTitle []string
	var index int
	var roomIndex int

	//Get Landing Page
	fmt.Println("Getting Landing Page...")
	res, opts := GetLanding(task)
	if res != 200 {
		fmt.Println("Error Getting Landing Page!")
		log.Panic()
	}

	for _, v := range opts {
		optsTitle = append(optsTitle, v.Title)
	}

	prompt := &survey.Select{
		Message: "Select Your Group:",
		Options: optsTitle,
	}

	survey.AskOne(prompt, &index)
	task.AttendeeInfo = opts[index]

	fmt.Println("Getting Home Page...")
	res = GetHome(task)
	if res != 200 {
		fmt.Println("Error Posting GroupID")
		log.Panic()
	}

	res, hotels := GetAllHotels(task)
	if res != 200 {
		fmt.Println("Error Getting All Hotels", res)
	}
	for _, v := range hotels {
		hotelTitle = append(hotelTitle, html.UnescapeString(v.Name))
	}

	prompt = &survey.Select{
		Message: "Select A Hotel: ",
		Options: hotelTitle,
	}
	survey.AskOne(prompt, &index)

	task.UserHotel.HotelID = hotels[index].ID

	if SendUpdate(task) != 200 {
		fmt.Println("Error Sending Update")
	}

	res, rooms := GetAllRooms(task)
	if res != 200 {
		fmt.Println("Error getting Rooms")
	}

	for _, v := range rooms[index].Blocks {
		roomTitle = append(roomTitle, html.UnescapeString(v.Name)+" - $"+strconv.Itoa(int(v.AverageBasicRate)))
	}

	prompt = &survey.Select{
		Message: "Select Which Room: ",
		Options: roomTitle,
	}

	survey.AskOne(prompt, &roomIndex)
	task.UserHotel.BlockID = rooms[index].Blocks[roomIndex].ID
	task.UserHotel.Charge = fmt.Sprintf("%.2f", rooms[index].Blocks[roomIndex].Charge)
	task.UserHotel.Subtotal = fmt.Sprintf("%.2f", rooms[index].Blocks[roomIndex].Charge)
	fmt.Printf("Selected: %s @ %s\n", rooms[index].Blocks[roomIndex].Name, rooms[index].Name)

	b := []Block{{
		HotelID:          task.UserHotel.HotelID,
		BlockID:          task.UserHotel.BlockID,
		CheckIn:          task.User.CheckIn,
		CheckOut:         task.User.CheckOut,
		NumberOfRooms:    task.User.NumberOfRooms,
		NumberOfGuests:   task.User.NumberOfGuests,
		NumberOfChildren: 0,
	}}

	bm := BlockMap{
		Blocks:      b,
		TotalGuests: task.User.NumberOfGuests,
		TotalRooms:  1,
	}

	up := UpdateStruct{
		HotelID:  task.UserHotel.HotelID,
		BlockMap: bm,
	}

	UpdateTotal(up, task)

	submitStr := EncodePersonalInformation(task)
	// fmt.Println(submitStr)

	SubmitInfo(submitStr, task)

	paymentStr := EncodePayment(task)
	// fmt.Println(paymentStr)
	SubmitPayment(task, paymentStr)

	SubmitReservation(task)

}

//numberofadults=1&numberofchildren=0&reservations%5B0%5D.ackNumber=&reservations%5B0%5D.id=0&reservations%5B0%5D.blockId=935655233&reservations%5B0%5D.checkInDate=9%2F13%2F24&reservations%5B0%5D.checkOutDate=9%2F16%2F24&reservations%5B0%5D.eventId=50383999&reservations%5B0%5D.groupTypeId=218613532&reservations%5B0%5D.hotelId=50099730&reservations%5B0%5D.statusId=0&reservations%5B0%5D.charge=489.00&reservations%5B0%5D.taxAmount=0&reservations%5B0%5D.subtotal=489.00&reservations%5B0%5D.guests%5B0%5D.id=0&reservations%5B0%5D.guests%5B0%5D.arrDate=9%2F13%2F24&reservations%5B0%5D.guests%5B0%5D.depDate=9%2F16%2F24&reservations%5B0%5D.guests%5B0%5D.prefix=&reservations%5B0%5D.guests%5B0%5D.firstName=First&reservations%5B0%5D.guests%5B0%5D.middleName=&reservations%5B0%5D.guests%5B0%5D.lastName=Last&reservations%5B0%5D.guests%5B0%5D.suffix=&reservations%5B0%5D.guests%5B0%5D.organization=Org&reservations%5B0%5D.guests%5B0%5D.position=&reservations%5B0%5D.guests%5B0%5D.email=email%40gmail.com&reservations%5B0%5D.guests%5B0%5D.confirmEmail=email%40gmail.com&reservations%5B0%5D.guests%5B0%5D.phoneNumber=3033334417&reservations%5B0%5D.guests%5B0%5D.familyName=&reservations%5B0%5D.guests%5B0%5D.givenName=&reservations%5B0%5D.guests%5B0%5D.address.country.alpha2Code=US&reservations%5B0%5D.guests%5B0%5D.address.address1=123+Main+St&reservations%5B0%5D.guests%5B0%5D.address.address2=&reservations%5B0%5D.guests%5B0%5D.address.city=Denver&reservations%5B0%5D.guests%5B0%5D.address.state=CO&reservations%5B0%5D.guests%5B0%5D.address.zip=80222&reservations%5B0%5D.rewardsProgram.id=184&reservations%5B0%5D.rewardMemberAccount=&_reservations%5B0%5D.accessibilityRequested=on&reservations%5B0%5D.smokingPreference.id=2&reservations%5B0%5D.guests%5B0%5D.specialRequests=&reservations%5B0%5D.specialRequests=&reservations%5B0%5D.optIn=true&_reservations%5B0%5D.optIn=on&_csrf=fbed38c6-f759-468c-a2b4-33e8e8a86a85
//numberofadults=1&numberofchildren=0&reservations%5B0%5D.ackNumber=&reservations%5B0%5D.id=0&reservations%5B0%5D.blockId=935655233&reservations%5B0%5D.checkInDate=9%2F13%2F24&reservations%5B0%5D.checkOutDate=9%2F16%2F24&reservations%5B0%5D.eventId=50383999&reservations%5B0%5D.groupTypeId=218613532&reservations%5B0%5D.hotelId=50099730&reservations%5B0%5D.statusId=0&reservations%5B0%5D.charge=828.00&reservations%5B0%5D.taxAmount=0&reservations%5B0%5D.subtotal=828.00&reservations%5B0%5D.guests%5B0%5D.id=0&reservations%5B0%5D.guests%5B0%5D.arrDate=9%2F13%2F24&reservations%5B0%5D.guests%5B0%5D.depDate=9%2F16%2F24&reservations%5B0%5D.guests%5B0%5D.prefix=&reservations%5B0%5D.guests%5B0%5D.firstName=Lenix&reservations%5B0%5D.guests%5B0%5D.middleName=&reservations%5B0%5D.guests%5B0%5D.lastName=Woof&reservations%5B0%5D.guests%5B0%5D.suffix=&reservations%5B0%5D.guests%5B0%5D.organization=Org&reservations%5B0%5D.guests%5B0%5D.position=&reservations%5B0%5D.guests%5B0%5D.email=email%40gmail.com&reservations%5B0%5D.guests%5B0%5D.confirmEmail=email%40gmail.com&reservations%5B0%5D.guests%5B0%5D.phoneNumber=3033334417&reservations%5B0%5D.guests%5B0%5D.familyName=&reservations%5B0%5D.guests%5B0%5D.givenName=&reservations%5B0%5D.guests%5B0%5D.address.country.alpha2Code=US&reservations%5B0%5D.guests%5B0%5D.address.address1=123+Main+St&reservations%5B0%5D.guests%5B0%5D.address.address2=&reservations%5B0%5D.guests%5B0%5D.address.city=Denver&reservations%5B0%5D.guests%5B0%5D.address.state=CO&reservations%5B0%5D.guests%5B0%5D.address.zip=80222&reservations%5B0%5D.rewardsProgram.id=184&reservations%5B0%5D.rewardMemberAccount=&_reservations%5B0%5D.accessibilityRequested=on&reservations%5B0%5D.smokingPreference.id=2&reservations%5B0%5D.guests%5B0%5D.specialRequests=&reservations%5B0%5D.specialRequests=&_reservations%5B0%5D.optIn=true&_reservations%5B0%5D.optIn=on_csrf=a9f41311%2F60ab%2F43a1%2F90f6%2Fd3e98dabbf34
