package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/colonyos/cpm/pkg/domain"
)

type JSONStateService struct {
	path string
	mu   sync.RWMutex
}

func NewJSONStateService(cpmHome string) (*JSONStateService, error) {
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(cpmHome, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cpm home directory: %w", err)
	}

	return &JSONStateService{
		path: filepath.Join(cpmHome, "state.json"),
	}, nil
}

func (s *JSONStateService) load() ([]*domain.Release, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []*domain.Release{}, nil
		}
		return nil, err
	}

	var releases []*domain.Release
	if err := json.Unmarshal(data, &releases); err != nil {
		return nil, fmt.Errorf("failed to parse state file: %w", err)
	}
	return releases, nil
}

func (s *JSONStateService) save(releases []*domain.Release) error {
	data, err := json.MarshalIndent(releases, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0644)
}

func (s *JSONStateService) Save(release *domain.Release) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	releases, err := s.load()
	if err != nil {
		return err
	}

	// Update existing or append
	found := false
	for i, r := range releases {
		if r.Name == release.Name {
			releases[i] = release
			found = true
			break
		}
	}
	if !found {
		releases = append(releases, release)
	}

	return s.save(releases)
}

func (s *JSONStateService) List() ([]*domain.Release, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.load()
}

func (s *JSONStateService) Get(name string) (*domain.Release, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	releases, err := s.load()
	if err != nil {
		return nil, err
	}

	for _, r := range releases {
		if r.Name == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("release %s not found", name)
}

func (s *JSONStateService) Delete(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	releases, err := s.load()
	if err != nil {
		return err
	}

	var newReleases []*domain.Release
	for _, r := range releases {
		if r.Name != name {
			newReleases = append(newReleases, r)
		}
	}

	return s.save(newReleases)
}
