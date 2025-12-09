package domain

import "time"

// Release represents an installed package instance
type Release struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	ColonyID    string    `json:"colonyId"`
	InstallTime time.Time `json:"installTime"`
	// We might add Status or Manifest copy later
}

// StateService defines operations for local state management
type StateService interface {
	Save(release *Release) error
	List() ([]*Release, error)
	Get(name string) (*Release, error)
	Delete(name string) error
}
