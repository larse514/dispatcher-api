package repository

import (
	"sync"

	"github.com/larse514/dispatcher-api/handlers"
)

// SourceRepositoryInMemory struct containing in memory list
type SourceRepositoryInMemory struct {
	Sources map[string][]handlers.Route
	Lock    *sync.Mutex
}

// CreateRoute method to create a route
func (repo SourceRepositoryInMemory) CreateRoute(source handlers.Source) error {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()

	repo.Sources[source.Name] = append(repo.Sources[source.Name], source.Routes...)
	return nil
}

// GetRoutes method to get a slice of routes for a Source
func (repo SourceRepositoryInMemory) GetRoutes(source handlers.Source) (handlers.Source, error) {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()
	source.Routes = repo.Sources[source.Name]
	return source, nil
}
