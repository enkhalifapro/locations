// Code generated by goa v3.2.5, DO NOT EDIT.
//
// locations HTTP server types
//
// Command:
// $ goa gen locations/api/design

package server

import (
	locations "locations/gen/locations"
)

// NowResponseBody is the type of the "locations" service "now" endpoint HTTP
// response body.
type NowResponseBody struct {
	// User country
	Country string `form:"country" json:"country" xml:"country"`
	// User current date
	Date string `form:"date" json:"date" xml:"date"`
	// User current time
	Time string `form:"time" json:"time" xml:"time"`
}

// NewNowResponseBody builds the HTTP response body from the result of the
// "now" endpoint of the "locations" service.
func NewNowResponseBody(res *locations.Location) *NowResponseBody {
	body := &NowResponseBody{
		Country: res.Country,
		Date:    res.Date,
		Time:    res.Time,
	}
	return body
}

// NewNowPayload builds a locations service now endpoint payload.
func NewNowPayload(xForwardedFor *string) *locations.NowPayload {
	v := &locations.NowPayload{}
	v.XForwardedFor = xForwardedFor

	return v
}
