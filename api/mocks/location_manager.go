// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	pkg "locations/pkg"

	mock "github.com/stretchr/testify/mock"
)

// LocationManager is an autogenerated mock type for the LocationManager type
type LocationManager struct {
	mock.Mock
}

// GetIPLocation provides a mock function with given fields: ip
func (_m *LocationManager) GetIPLocation(ip string) (*pkg.Location, error) {
	ret := _m.Called(ip)

	var r0 *pkg.Location
	if rf, ok := ret.Get(0).(func(string) *pkg.Location); ok {
		r0 = rf(ip)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pkg.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
