package repository

import (
	"fmt"
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
func (repo SourceRepositoryInMemory) GetSource(source handlers.Source) (handlers.Source, error) {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()
	source.Routes = repo.Sources[source.Name]
	return source, nil
}

// GetAllSources method to get all sources
func (repo SourceRepositoryInMemory) GetAllSources() ([]handlers.Source, error) {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()
	sources := make([]handlers.Source, 0)
	for key, value := range repo.Sources {
		fmt.Println("Key:", key, "Value:", value)
		sources = append(sources, handlers.Source{Name: key, Routes: value})
	}

	return sources, nil
}
