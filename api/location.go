package api

import (
	"context"
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
func (s *Location) Now(ctx context.Context, p *locations.NowPayload) (res *locations.Location, err error) {
	if p.XForwardedFor == nil {
		// todo: enable the next line when finish debug
		// return nil, errors.New("IP not found")
	}
	// todo: enable the next line when finish debug
	// ip := *p.XForwardedFor

	// todo: remove the next line later
	ip := "207.46.197.32"

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
