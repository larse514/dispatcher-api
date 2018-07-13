package repository

import (
	"sync"
	"testing"

	"github.com/larse514/dispatcher-api/handlers"
)

func TestSourceRepositoryInMemory_AddOneSourceGetOneSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	source, err := repo.GetRoutes(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}

func TestSourceRepositoryInMemory_AddTwoSourceGetTwoSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 2)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	routes[1] = handlers.Route{URL: "https://www.msn.com"}
	sources[1] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	source, err := repo.GetRoutes(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
	if actual[1] != excepted[1] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}

func TestSourceRepositoryInMemory_AddTwoDifferentSourceGetTwoDifferentSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	sources[1] = handlers.Source{Name: "Service2", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	source, err := repo.GetRoutes(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}

	source, err = repo.GetRoutes(handlers.Source{Name: sources[1].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted = routes
	actual = source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}
