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
	"time"

	"github.com/verdverm/frisby"
)

const apiURI = "http://localhost:3001/api/v1"
const testUserID = "4481c18058aa494fb033b65665c38de2"

func main() {
	addingTodoListItemForUnknownUserFails()

	// When
	registeringUser()
	// Then
	userLookupReadmodelContainsUser()

	// When
	addingTodoListItem()
	// Then
	todoListContainsNewItem()

	omittingMandatoryCommandParameterFails()

	frisby.Global.PrintReport()
	if frisby.Global.NumErrored > 0 {
		os.Exit(-1)
	}
}

func addingTodoListItemForUnknownUserFails() {
	frisby.Create("Executing Command with id missing in read model fails.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI+"/TodoItem/AddItem").
		SetJson(todoItem{
			UserID:      testUserID,
			Description: "Carpe Diem",
			TodoitemID:  "10000"}).
		Send().
		ExpectStatus(400).
		ExpectJson("message.0.field", "userId").
		ExpectJson("message.0.msg", "userId does not exist.").
		ExpectJsonLength("message", 1)
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

func addingTodoListItem() {
	frisby.Create("Executing valid Command succeeds.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI + "/TodoItem/AddItem").
		SetJson(todoItem{UserID: testUserID, Description: "Carpe Diem", TodoitemID: "10000"}).
		Send().
		ExpectStatus(202)
}

func userLookupReadmodelContainsUser() {
	time.Sleep(100 * time.Millisecond)
	frisby.Create("GET UserLookup Read Model succeeds.").
		Get(apiURI+"/r/UserLookup").
		Send().
		ExpectStatus(200).
		ExpectJson("0.userId", testUserID).
		ExpectJson("0.notificationEmail", "fake@mail.com")
}

func todoListContainsNewItem() {
	time.Sleep(100 * time.Millisecond)
	frisby.Create("GET Read Model succeeds.").
		Get(apiURI+"/r/TODOList").
		Send().
		ExpectStatus(200).
		ExpectJson("0.todoitemId", "10000").
		ExpectJson("0.userId", testUserID).
		ExpectJson("0.description", "Carpe Diem")
}

func omittingMandatoryCommandParameterFails() {
	frisby.Create("Omitting mandatory field fails.").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json, text/plain, */*").
		Post(apiURI+"/TodoItem/AddItem").
		SetJson(todoItem{UserID: "123456", Description: "This will fail.", TodoitemID: ""}).
		Send().
		ExpectStatus(400).
		ExpectJson("message.1.field", "todoitemId").
		ExpectJson("message.1.msg", "todoitemId is a required field.").
		ExpectJsonLength("message", 2) // ... because of the non-existing userId 123456. That's a 2nd validation error, but we aren't interested in it in this test.
}

type user struct {
	UserID            string `json:"userId"`
	NotificationEmail string `json:"notificationEmail"`
}

type todoItem struct {
	UserID      string `json:"userId"`
	TodoitemID  string `json:"todoitemId"`
	Description string `json:"description"`
}
