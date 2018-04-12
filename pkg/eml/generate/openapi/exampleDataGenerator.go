package openapi

import (
	"math/rand"
	"strings"

	randomdata "github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
)

// Example data
type Example struct {
	Parameter func(string, string) string
}

var streamIDs map[string]string

func parameterGenerator(parameter string, streamName string) string {
	if isStreamID(parameter, streamName) {
		if streamIDs == nil {
			streamIDs = make(map[string]string)
		}
		if streamID, ok := streamIDs[parameter]; ok {
			return streamID
		}
		streamIDs[parameter] = guid()
		return streamIDs[parameter]
	}
	if strings.Contains(strings.ToLower(parameter), "email") {
		return randomdata.Email()
	}
	if strings.Contains(strings.ToLower(parameter), "date") {
		return randomdata.FullDate()
	}
	if strings.Contains(strings.ToLower(parameter), "country") {
		return randomdata.Country(randomdata.ThreeCharCountry)
	}
	if strings.Contains(strings.ToLower(parameter), "firstname") {
		maleOrFemale := rand.Intn(1)
		return randomdata.FirstName(maleOrFemale)
	}
	if strings.Contains(strings.ToLower(parameter), "lastname") || strings.Contains(strings.ToLower(parameter), "surname") {
		return randomdata.LastName()
	}
	if strings.Contains(strings.ToLower(parameter), "address") {
		return randomdata.Address()
	}
	if strings.Contains(strings.ToLower(parameter), "street") {
		return randomdata.Street()
	}
	if strings.Contains(strings.ToLower(parameter), "postalcode") {
		return randomdata.PostalCode("CA")
	}
	if strings.Contains(strings.ToLower(parameter), "city") {
		return randomdata.City()
	}
	return randomdata.SillyName()
}

func isStreamID(parameter string, aggregateType string) bool {
	streamName := aggregateType + strings.ToLower(parameter[len(parameter)-2:])
	return streamName == aggregateType+"id"
}

func guid() string {
	guid := uuid.NewV4().String()
	return strings.Replace(guid, "-", "", -1)
}
