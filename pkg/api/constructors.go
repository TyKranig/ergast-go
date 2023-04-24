package api

import (
	"context"

	"github.com/tylerkranig/ergast-api/pkg/ergast"
)

func (fapi *FormulaAPI) GetConstructorsByYear(ctx context.Context, year string) (*ergast.ConstructorsByYearResponse, error) {
	return ergast.CallErgast[ergast.ConstructorsByYearResponse](ctx, func() (res *ergast.GetConstructorsByYearResponse, err error) {
		return fapi.client.GetConstructorsByYearWithResponse(ctx, year, &ergast.GetConstructorsByYearParams{})
	})
}

func (fapi *FormulaAPI) GetConstructorStandingsByYear(ctx context.Context, year string) (*ergast.ConstructorStandingsByYearResponse, error) {
	return ergast.CallErgast[ergast.ConstructorStandingsByYearResponse](ctx, func() (res *ergast.GetConstructorStandingsByYearResponse, err error) {
		return fapi.client.GetConstructorStandingsByYearWithResponse(ctx, year, &ergast.GetConstructorStandingsByYearParams{})
	})
}

func (fapi *FormulaAPI) GetConstructorResultsByYear(ctx context.Context, year, constructor string) (*ergast.ResultsByYearResponse, error) {
	return ergast.CallErgast[ergast.ResultsByYearResponse](ctx, func() (res *ergast.GetResultsByYearAndConstructorResponse, err error) {
		return fapi.client.GetResultsByYearAndConstructorWithResponse(ctx, year, constructor, &ergast.GetResultsByYearAndConstructorParams{})
	})
}
