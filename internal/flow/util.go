package flow

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EncodePersonalInformation(task *UserTask) string {

	var values string
	values = ("numberofadults" + "=" + strconv.Itoa(task.User.NumberOfGuests)) + "&"
	values += ("numberofchildren" + "=" + "0") + "&"
	values += ("reservations[0].ackNumber" + "=" + "") + "&"
	values += ("reservations[0].id" + "=" + "0") + "&"
	values += ("reservations[0].blockId" + "=" + strconv.Itoa(task.UserHotel.BlockID)) + "&"
	values += ("reservations[0].checkInDate" + "=" + task.Checkout.CheckoutDateIn) + "&"
	values += ("reservations[0].checkOutDate" + "=" + task.Checkout.CheckoutDateOut) + "&"
	values += ("reservations[0].eventId" + "=" + (task.Hotel.EventID)) + "&"
	values += ("reservations[0].groupTypeId" + "=" + task.AttendeeInfo.Value) + "&"
	values += ("reservations[0].hotelId" + "=" + strconv.Itoa(task.UserHotel.HotelID)) + "&"
	values += ("reservations[0].statusId" + "=" + "0") + "&"
	values += ("reservations[0].charge" + "=" + task.UserHotel.Charge) + "&"
	values += ("reservations[0].taxAmount" + "=" + "0") + "&"
	values += ("reservations[0].subtotal" + "=" + task.UserHotel.Charge) + "&"
	values += ("reservations[0].guests[0].id" + "=" + "0") + "&"
	values += ("reservations[0].guests[0].arrDate" + "=" + task.Checkout.CheckoutDateIn) + "&"
	values += ("reservations[0].guests[0].depDate" + "=" + task.Checkout.CheckoutDateOut) + "&"
	values += ("reservations[0].guests[0].prefix" + "=" + "") + "&"
	values += ("reservations[0].guests[0].firstName" + "=" + task.Personal.FirstName) + "&"
	values += ("reservations[0].guests[0].middleName" + "=" + "") + "&"
	values += ("reservations[0].guests[0].lastName" + "=" + task.Personal.LastName) + "&"
	values += ("reservations[0].guests[0].suffix" + "=" + "") + "&"
	values += ("reservations[0].guests[0].organization" + "=" + "Blehh") + "&"
	values += ("reservations[0].guests[0].position" + "=" + "") + "&"
	values += ("reservations[0].guests[0].email" + "=" + task.Personal.EmailAddress) + "&"
	values += ("reservations[0].guests[0].confirmEmail" + "=" + task.Personal.EmailAddress) + "&"
	values += ("reservations[0].guests[0].phoneNumber" + "=" + task.Personal.PhoneNumber) + "&"
	values += ("reservations[0].guests[0].familyName" + "=" + "") + "&"
	values += ("reservations[0].guests[0].givenName" + "=" + "") + "&"
	values += ("reservations[0].guests[0].address.country.alpha2Code" + "=" + "US") + "&"
	values += ("reservations[0].guests[0].address.address1" + "=" + task.Personal.Address1) + "&"
	values += ("reservations[0].guests[0].address.address2" + "=" + task.Personal.Address2) + "&"
	values += ("reservations[0].guests[0].address.city" + "=" + task.Personal.City) + "&"
	values += ("reservations[0].guests[0].address.state" + "=" + task.Personal.State) + "&"
	values += ("reservations[0].guests[0].address.zip" + "=" + task.Personal.Zip) + "&"
	values += ("reservations[0].rewardsProgram.id" + "=" + "184") + "&"
	values += ("reservations[0].rewardMemberAccount" + "=" + "") + "&"
	values += ("_reservations[0].accessibilityRequested" + "=" + "on") + "&"
	values += ("reservations[0].smokingPreference.id" + "=" + "2") + "&"
	values += ("reservations[0].guests[0].specialRequests" + "=" + "") + "&"
	values += ("reservations[0].specialRequests" + "=" + "") + "&"
	values += ("_reservations[0].optIn" + "=" + "true") + "&"

	values = strings.ReplaceAll(values, "[", "%5B")
	values = strings.ReplaceAll(values, "]", "%5D")
	values = strings.ReplaceAll(values, "-", "%2F")
	values = strings.ReplaceAll(values, "@", "%40")

	values += ("_csrf" + "=" + task.Client.CSRF)

	// str := ("numberofadults=1&numberofchildren=0&reservations%5B0%5D.ackNumber=&reservations%5B0%5D.id=0&reservations%5B0%5D.blockId=935655233&reservations%5B0%5D.checkInDate=9%2F13%2F24&reservations%5B0%5D.checkOutDate=9%2F16%2F24&reservations%5B0%5D.eventId=50383999&reservations%5B0%5D.groupTypeId=218613532&reservations%5B0%5D.hotelId=50099730&reservations%5B0%5D.statusId=0&reservations%5B0%5D.charge=489.00&reservations%5B0%5D.taxAmount=0&reservations%5B0%5D.subtotal=489.00&reservations%5B0%5D.guests%5B0%5D.id=0&reservations%5B0%5D.guests%5B0%5D.arrDate=9%2F13%2F24&reservations%5B0%5D.guests%5B0%5D.depDate=9%2F16%2F24&reservations%5B0%5D.guests%5B0%5D.prefix=&reservations%5B0%5D.guests%5B0%5D.firstName=Lenix&reservations%5B0%5D.guests%5B0%5D.middleName=&reservations%5B0%5D.guests%5B0%5D.lastName=Woof&reservations%5B0%5D.guests%5B0%5D.suffix=&reservations%5B0%5D.guests%5B0%5D.organization=Blehh&reservations%5B0%5D.guests%5B0%5D.position=&reservations%5B0%5D.guests%5B0%5D.email=Realemail%40gmail.com&reservations%5B0%5D.guests%5B0%5D.confirmEmail=Realemail%40gmail.com&reservations%5B0%5D.guests%5B0%5D.phoneNumber=3033334417&reservations%5B0%5D.guests%5B0%5D.familyName=&reservations%5B0%5D.guests%5B0%5D.givenName=&reservations%5B0%5D.guests%5B0%5D.address.country.alpha2Code=US&reservations%5B0%5D.guests%5B0%5D.address.address1=123+Main+St.&reservations%5B0%5D.guests%5B0%5D.address.address2=&reservations%5B0%5D.guests%5B0%5D.address.city=Denver&reservations%5B0%5D.guests%5B0%5D.address.state=CO&reservations%5B0%5D.guests%5B0%5D.address.zip=80222&reservations%5B0%5D.rewardsProgram.id=184&reservations%5B0%5D.rewardMemberAccount=&_reservations%5B0%5D.accessibilityRequested=on&reservations%5B0%5D.smokingPreference.id=2&reservations%5B0%5D.guests%5B0%5D.specialRequests=&reservations%5B0%5D.specialRequests=&_reservations%5B0%5D.optIn=on&_csrf=" + task.Client.CSRF)
	return (values)
}

func EncodePayment(task *UserTask) string {
	var values string

	// Card Types:
	// 100 - Visa
	// 101 - MasterCard
	// 102 - AmEx
	// 103 - Discover
	// 104 - Diners Club
	// 105 JCB

	values += ("billingInfo%5B0%5D.payment.paymentType" + "=" + "CCPayment") + "&"
	values += ("billingInfo%5B0%5D.payer.holderName" + "=" + task.Personal.FirstName + "+" + task.Personal.LastName) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.cardTypeId" + "=" + "100") + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.depDate" + "=" + "") + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.cardNumber" + "=" + task.Personal.Payment.CreditNumber) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.expiryMonth" + "=" + task.Personal.Payment.ExpMonth) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.expiryYear" + "=" + task.Personal.Payment.ExpYear) + "&"
	values += ("billingInfo%5B0%5D.payer.address.country.alpha2Code" + "=" + "US") + "&"
	values += ("billingInfo%5B0%5D.payer.address.address1" + "=" + task.Personal.Address1) + "&"
	values += ("billingInfo%5B0%5D.payer.address.address2" + "=" + task.Personal.Address2) + "&"
	values += ("billingInfo%5B0%5D.payer.phoneNumber" + "=" + task.Personal.PhoneNumber) + "&"
	values += ("billingInfo%5B0%5D.payer.address.city" + "=" + task.Personal.City) + "&"
	values += ("billingInfo%5B0%5D.payer.address.state" + "=" + task.Personal.State) + "&"
	values += ("billingInfo%5B0%5D.payer.address.zip" + "=" + task.Personal.Zip) + "&"

	values = strings.ReplaceAll(values, "[", "%5B")
	values = strings.ReplaceAll(values, "]", "%5D")
	// values = strings.ReplaceAll(values, "-", "%2F")
	values = strings.ReplaceAll(values, "@", "%40")

	values += ("splitFolio%5B0%5D" + "=" + "101") + "&"
	values += ("_csrf" + "=" + task.Client.CSRF)

	// billingInfo%5B0%5D.payment.paymentType=CCPayment&
	// billingInfo%5B0%5D.payer.holderName=First+Last&
	// billingInfo%5B0%5D.payment.creditCard.cardTypeId=100&
	// billingInfo%5B0%5D.payment.creditCard.depDate=&
	// billingInfo%5B0%5D.payment.creditCard.cardNumber=4444-5555-5555-5555&
	// billingInfo%5B0%5D.payment.creditCard.expiryMonth=02&
	// billingInfo%5B0%5D.payment.creditCard.expiryYear=2029&
	// billingInfo%5B0%5D.payer.address.country.alpha2Code=US&
	// billingInfo%5B0%5D.payer.address.address1=123+Main+St&
	// billingInfo%5B0%5D.payer.address.address2=&
	// billingInfo%5B0%5D.payer.phoneNumber=3033334417&
	// billingInfo%5B0%5D.payer.address.city=Denver&
	// billingInfo%5B0%5D.payer.address.state=CO&
	// billingInfo%5B0%5D.payer.address.zip=80222&
	// splitFolio%5B0%5D=101&
	// _csrf=fbed38c6-f759-468c-a2b4-33e8e8a86a85
	return values
}

func SaveSuccessfulCaptcha(raw, sol string) error {

	filePath := "data/successCaptcha/" + sol + ".png"
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err
	}

	// Write the decoded data to a file
	err = os.WriteFile(filePath, decoded, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Image saved to: %s\n", filePath)
	return nil

}

func SaveUnsuccessfulCaptcha(raw, sol string) error {

	filePath := "data/failedCaptcha/" + sol + ".png"
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err
	}

	// Write the decoded data to a file
	err = os.WriteFile(filePath, decoded, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Image saved to: %s\n", filePath)
	return nil

}
