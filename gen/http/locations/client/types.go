// Code generated by goa v3.2.5, DO NOT EDIT.
//
// locations HTTP client types
//
// Command:
// $ goa gen locations/api/design

package client

import (
	locations "locations/gen/locations"

	goa "goa.design/goa/v3/pkg"
)

// NowResponseBody is the type of the "locations" service "now" endpoint HTTP
// response body.
type NowResponseBody struct {
	// User country
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// User current date
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// User current time
	Time *string `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

// NewNowLocationOK builds a "locations" service "now" endpoint result from a
// HTTP "OK" response.
func NewNowLocationOK(body *NowResponseBody) *locations.Location {
	v := &locations.Location{
		Country: *body.Country,
		Date:    *body.Date,
		Time:    *body.Time,
	}

	return v
}

// ValidateNowResponseBody runs the validations defined on NowResponseBody
func ValidateNowResponseBody(body *NowResponseBody) (err error) {
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Date == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date", "body"))
	}
	if body.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "body"))
	}
	return
}
