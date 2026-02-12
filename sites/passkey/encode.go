package passkey

import (
	"log"
	"strconv"
	"strings"
)

func (e *PasskeyTask) EncodePersonalInformation() {

	var values string

	// formattedInTest := "3/6/24"
	// formattedOutTest := "3/8/24"

	values = ("numberofadults" + "=" + strconv.Itoa(1)) + "&" // Number of Guests
	values += ("numberofchildren" + "=" + "0") + "&"
	values += ("reservations[0].ackNumber" + "=" + "") + "&"
	values += ("reservations[0].id" + "=" + "0") + "&"
	values += ("reservations[0].blockId" + "=" + strconv.Itoa(e.BlockID)) + "&"
	values += ("reservations[0].checkInDate" + "=" + e.CheckoutDateIn) + "&"
	values += ("reservations[0].checkOutDate" + "=" + e.CheckoutDateOut) + "&"
	values += ("reservations[0].eventId" + "=" + (e.EventID)) + "&"
	values += ("reservations[0].groupTypeId" + "=" + e.GroupID) + "&"
	values += ("reservations[0].hotelId" + "=" + strconv.Itoa(e.HotelID)) + "&"
	values += ("reservations[0].statusId" + "=" + "0") + "&"
	values += ("reservations[0].charge" + "=" + e.Charge) + "&"
	values += ("reservations[0].taxAmount" + "=" + e.TaxAmount) + "&"
	values += ("reservations[0].subtotal" + "=" + e.Subtotal) + "&"
	values += ("reservations[0].guests[0].id" + "=" + "0") + "&"
	values += ("reservations[0].guests[0].arrDate" + "=" + e.CheckoutDateIn) + "&"
	values += ("reservations[0].guests[0].depDate" + "=" + e.CheckoutDateOut) + "&"
	values += ("reservations[0].guests[0].prefix" + "=" + "") + "&"
	values += ("reservations[0].guests[0].firstName" + "=" + e.Task.Profile.UserInformation.FirstName) + "&"
	values += ("reservations[0].guests[0].middleName" + "=" + "") + "&"
	values += ("reservations[0].guests[0].lastName" + "=" + e.Task.Profile.UserInformation.LastName) + "&"
	values += ("reservations[0].guests[0].suffix" + "=" + "") + "&"
	values += ("reservations[0].guests[0].organization" + "=" + "") + "&"
	values += ("reservations[0].guests[0].position" + "=" + "") + "&"
	values += ("reservations[0].guests[0].email" + "=" + e.Task.Profile.UserInformation.Email) + "&"
	values += ("reservations[0].guests[0].confirmEmail" + "=" + e.Task.Profile.UserInformation.Email) + "&"
	values += ("reservations[0].guests[0].phoneNumber" + "=" + e.Task.Profile.UserInformation.PhoneNumber) + "&"
	// values += ("reservations[0].guests[0].familyName" + "=" + "") + "&"
	// values += ("reservations[0].guests[0].givenName" + "=" + "") + "&"
	values += ("reservations[0].guests[0].address.country.alpha2Code" + "=" + "US") + "&"
	values += ("reservations[0].guests[0].address.address1" + "=" + e.Task.Profile.UserInformation.AddressLine1) + "&"
	values += ("reservations[0].guests[0].address.address2" + "=" + e.Task.Profile.UserInformation.AddressLine2) + "&"
	values += ("reservations[0].guests[0].address.city" + "=" + e.Task.Profile.UserInformation.City) + "&"
	values += ("reservations[0].guests[0].address.state" + "=" + e.Task.Profile.UserInformation.State) + "&"
	values += ("reservations[0].guests[0].address.zip" + "=" + e.Task.Profile.UserInformation.Zip) + "&"
	values += ("reservations[0].rewardsProgram.id" + "=" + "120") + "&"
	values += ("reservations[0].rewardMemberAccount" + "=" + "") + "&"
	values += ("_reservations[0].accessibilityRequested" + "=" + "on") + "&"
	values += ("reservations[0].smokingPreference.id" + "=" + "2") + "&"
	values += ("reservations[0].guests[0].specialRequests" + "=" + "") + "&"
	values += ("reservations[0].specialRequests" + "=" + "") + "&"
	values += ("reservations[0].optIn" + "=" + "true") + "&"
	values += ("_reservations[0].optIn" + "=" + "on") + "&"

	values = strings.ReplaceAll(values, "[", "%5B")
	values = strings.ReplaceAll(values, "]", "%5D")
	values = strings.ReplaceAll(values, "-", "%2F")
	values = strings.ReplaceAll(values, "@", "%40")

	values += ("_csrf" + "=" + e.CSRF)

	// str := ("numberofadults=1&numberofchildren=0&reservations%5B0%5D.ackNumber=&reservations%5B0%5D.id=0&reservations%5B0%5D.blockId=935655233&reservations%5B0%5D.checkInDate=9%2F13%2F24&reservations%5B0%5D.checkOutDate=9%2F16%2F24&reservations%5B0%5D.eventId=50383999&reservations%5B0%5D.groupTypeId=218613532&reservations%5B0%5D.hotelId=50099730&reservations%5B0%5D.statusId=0&reservations%5B0%5D.charge=489.00&reservations%5B0%5D.taxAmount=0&reservations%5B0%5D.subtotal=489.00&reservations%5B0%5D.guests%5B0%5D.id=0&reservations%5B0%5D.guests%5B0%5D.arrDate=9%2F13%2F24&reservations%5B0%5D.guests%5B0%5D.depDate=9%2F16%2F24&reservations%5B0%5D.guests%5B0%5D.prefix=&reservations%5B0%5D.guests%5B0%5D.firstName=Lenix&reservations%5B0%5D.guests%5B0%5D.middleName=&reservations%5B0%5D.guests%5B0%5D.lastName=Woof&reservations%5B0%5D.guests%5B0%5D.suffix=&reservations%5B0%5D.guests%5B0%5D.organization=Blehh&reservations%5B0%5D.guests%5B0%5D.position=&reservations%5B0%5D.guests%5B0%5D.email=Realemail%40gmail.com&reservations%5B0%5D.guests%5B0%5D.confirmEmail=Realemail%40gmail.com&reservations%5B0%5D.guests%5B0%5D.phoneNumber=3033334417&reservations%5B0%5D.guests%5B0%5D.familyName=&reservations%5B0%5D.guests%5B0%5D.givenName=&reservations%5B0%5D.guests%5B0%5D.address.country.alpha2Code=US&reservations%5B0%5D.guests%5B0%5D.address.address1=123+Main+St.&reservations%5B0%5D.guests%5B0%5D.address.address2=&reservations%5B0%5D.guests%5B0%5D.address.city=Denver&reservations%5B0%5D.guests%5B0%5D.address.state=CO&reservations%5B0%5D.guests%5B0%5D.address.zip=80222&reservations%5B0%5D.rewardsProgram.id=184&reservations%5B0%5D.rewardMemberAccount=&_reservations%5B0%5D.accessibilityRequested=on&reservations%5B0%5D.smokingPreference.id=2&reservations%5B0%5D.guests%5B0%5D.specialRequests=&reservations%5B0%5D.specialRequests=&_reservations%5B0%5D.optIn=on&_csrf=" + task.Client.CSRF)
	e.EncodedInfo = values
}

func (e *PasskeyTask) EncodePayment() {
	var values string

	// Card Types:
	// 100 - Visa
	// 101 - MasterCard
	// 102 - AmEx
	// 103 - Discover
	// 104 - Diners Club
	// 105 JCB

	values += ("billingInfo%5B0%5D.payment.paymentType" + "=" + "CCPayment") + "&"
	values += ("billingInfo%5B0%5D.payer.holderName" + "=" + e.Task.Profile.UserInformation.FirstName + "+" + e.Task.Profile.UserInformation.LastName) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.cardTypeId" + "=" + "100") + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.depDate" + "=" + "") + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.cardNumber" + "=" + e.Task.Profile.UserInformation.PaymentInformation.CardNumber) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.expiryMonth" + "=" + e.Task.Profile.UserInformation.PaymentInformation.ExpMon) + "&"
	values += ("billingInfo%5B0%5D.payment.creditCard.expiryYear" + "=" + e.Task.Profile.UserInformation.PaymentInformation.ExpYr) + "&"
	values += ("billingInfo%5B0%5D.payer.address.country.alpha2Code" + "=" + "US") + "&"
	values += ("billingInfo%5B0%5D.payer.address.address1" + "=" + e.Task.Profile.UserInformation.AddressLine1) + "&"
	values += ("billingInfo%5B0%5D.payer.address.address2" + "=" + e.Task.Profile.UserInformation.AddressLine2) + "&"
	values += ("billingInfo%5B0%5D.payer.phoneNumber" + "=" + e.Task.Profile.UserInformation.PhoneNumber) + "&"
	values += ("billingInfo%5B0%5D.payer.address.city" + "=" + e.Task.Profile.UserInformation.City) + "&"
	values += ("billingInfo%5B0%5D.payer.address.state" + "=" + e.Task.Profile.UserInformation.State) + "&"
	values += ("billingInfo%5B0%5D.payer.address.zip" + "=" + e.Task.Profile.UserInformation.Zip) + "&"

	values = strings.ReplaceAll(values, "[", "%5B")
	values = strings.ReplaceAll(values, "]", "%5D")
	// values = strings.ReplaceAll(values, "-", "%2F")
	values = strings.ReplaceAll(values, "@", "%40")

	// you must change this value to 101 or 102 (not sure why yet)
	values += ("splitFolio%5B0%5D" + "=" + e.PaymentFolio) + "&"
	values += ("_csrf" + "=" + e.CSRF)

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
	e.EncodedPayment = values
}

//Not Encoding but idk anywhere else to put it

func (t *PasskeyTask) UpdateStatus(stage string) {
	t.Task.Stage = stage
	log.Printf("Task ID: %s | Mode: %s | Profile: %s | Event: %s | Hotel: %s | Room: %s | Status: %s\r\n", t.Task.ID, t.Mode, t.Profile.ProfileName, t.EventName, t.HotelName, t.BlockName, t.Task.Stage)

}
