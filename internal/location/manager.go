package location

import (
	"locations/pkg"
)

// Manager ...
type Manager struct {
	ipManager IPManager
}

// IPManager ...
type IPManager interface {
	GetLocation(ip string) (*pkg.Location, error)
}

// NewManager contains location service core functionalities
func NewManager(mgr IPManager) *Manager {
	return &Manager{
		ipManager: mgr,
	}
}

// GetIPLocation ...
func (m *Manager) GetIPLocation(ip string) (*pkg.Location, error) {
	return m.ipManager.GetLocation(ip)
}
