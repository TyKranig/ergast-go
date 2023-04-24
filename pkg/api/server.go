package api

import (
	"fmt"

	"github.com/tylerkranig/ergast-api/pkg/ergast"
)

var DEFAULTURL string = "https://ergast.com/api/f1"

// top level api struct
type FormulaAPI struct {
	client ergast.ClientWithResponses
}

func NewFormulaAPI(url ...string) (*FormulaAPI, error) {
	serverURL := DEFAULTURL
	if len(url) > 0 && len(url[0]) > 0 {
		serverURL = url[0]
	}

	apiClient, err := ergast.NewClientWithResponses(serverURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %w", err)
	}

	return &FormulaAPI{
		client: *apiClient,
	}, nil
}
