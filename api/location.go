package api

import (
	"context"
	"errors"
	"locations/gen/locations"
	"locations/pkg"
)

// Location api implementation.
type Location struct {
	locationManager LocationManager
}

// LocationManager ...
type LocationManager interface {
	GetIPLocation(ip string) (*pkg.Location, error)
}

// NewLocation returns the location service implementation.
func NewLocation(mgr LocationManager) *Location {
	// Build and return service implementation.
	return &Location{
		locationManager: mgr,
	}
}

// Now
func (s *Location) Now(_ context.Context, p *locations.NowPayload) (res *locations.Location, err error) {
	if p.XForwardedFor == nil {
		return nil, errors.New("IP not found")
	}
	ip := *p.XForwardedFor

	loc, err := s.locationManager.GetIPLocation(ip)
	if err != nil {
		return nil, err
	}

	return &locations.Location{
		Country: loc.Country,
		Date:    loc.Date,
		Time:    loc.Time,
	}, nil
}
