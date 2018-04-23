/*
Event Markup Language Compliance Test

	Tests whether an API created from EML by a builder such as les-node complies with the EML specification.

	* Are all the command processor end points working?
	* Are the read models working?
	* Are all the business rules implemented?
	* Are the required validation errors returned when executing invalid commands?

*/
package main

import (
	"os"

	"github.com/verdverm/frisby"
)

const apiURI = "http://localhost:3001/api/v1"
const testUserID = "4481c18058aa494fb033b65665c38de2"

func main() {
	registeringUser()

	validatingRegisteredUserSucceeds()

	validatingUserWhoIsntRegisteredFails()

	deletingUser()

	validatingUserWhoHasBeenDeletedFails()

	frisby.Global.PrintReport()
	if frisby.Global.NumErrored > 0 {
		os.Exit(-1)
	}
}

func registeringUser() {
	frisby.Create("Executing valid registeringUser Command succeeds.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI + "/User/Register").
		SetJson(user{UserID: testUserID, NotificationEmail: "fake@mail.com"}).
		Send().
		ExpectStatus(202)
}

func deletingUser() {
	frisby.Create("Delete User Command succeeds.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI + "/User/DeleteUser").
		SetJson(user{UserID: testUserID}).
		Send().
		ExpectStatus(202)
}

func validatingRegisteredUserSucceeds() {
	frisby.Create("Validating registered user succeeds.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI + "/User/MarkUserAsValidated").
		SetJson(user{UserID: testUserID}).
		Send().
		ExpectStatus(202)
}

func validatingUserWhoIsntRegisteredFails() {
	frisby.Create("Validating user who isn't registered fails.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI+"/User/MarkUserAsValidated").
		SetJson(user{UserID: "123456789"}).
		Send().
		ExpectStatus(400).
		ExpectJson("message.0.field", "").
		ExpectJson("message.0.msg", "Cannot MarkUserAsValidated. UserRegistered must have occurred.").
		ExpectJsonLength("message", 1)
}

func validatingUserWhoHasBeenDeletedFails() {
	frisby.Create("Validating user who isn't registered fails.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI+"/User/MarkUserAsValidated").
		SetJson(user{UserID: testUserID}).
		Send().
		ExpectStatus(400).
		ExpectJson("message.0.field", "").
		ExpectJson("message.0.msg", "Cannot MarkUserAsValidated. UserDeleted must not have occurred.").
		ExpectJsonLength("message", 1)
}

type user struct {
	UserID            string `json:"userId"`
	NotificationEmail string `json:"notificationEmail"`
}
