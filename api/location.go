package api

import (
	"context"
	"fmt"
	"locations/gen/locations"
)

// Location api implementation.
type Location struct {
}

// NewUser returns the users service implementation.
func NewLocation() *Location {
	// Build and return service implementation.
	return &Location{}
}

// Now
func (s *Location) Now(context.Context) (res *locations.Location, err error) {
	fmt.Println("in nowwwww")
	return &locations.Location{
		Country: "Egypt",
		Date:    "date",
		Time:    "time",
	}, nil
}
