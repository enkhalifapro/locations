package api

import (
	"errors"
	"locations/api/mocks"
	"locations/gen/locations"
	"locations/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	// Arrange
	ip := "2.2.2.2"
	locationMgrMock := &mocks.LocationManager{}
	locationMgrMock.On("GetIPLocation", ip).Return(&pkg.Location{
		Country: "Egypt",
		Date:    "14-11-2020",
		Time:    "15:20:10"}, nil)

	// Act
	loc := NewLocation(locationMgrMock)
	res, err := loc.Now(nil, &locations.NowPayload{XForwardedFor: &ip})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Egypt", res.Country)
	assert.Equal(t, "14-11-2020", res.Date)
	assert.Equal(t, "15:20:10", res.Time)

	// test negative path
	ip = "3.3.3.3"
	locationMgrMock.On("GetIPLocation", ip).Return(nil, errors.New("unknown location"))

	// Act
	loc = NewLocation(locationMgrMock)
	res, err = loc.Now(nil, &locations.NowPayload{})

	// Assert
	assert.Error(t, err)
}
