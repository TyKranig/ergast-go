package api

import (
	"context"

	"github.com/tylerkranig/ergast-api/pkg/ergast"
)

func (fapi *FormulaAPI) GetResultsByYearAndRoundWithResponse(ctx context.Context, year string, round string) (*ergast.ResultsByYearResponse, error) {
	return ergast.CallErgast[ergast.ResultsByYearResponse](ctx, func() (res *ergast.GetResultsByYearAndRoundResponse, err error) {
		return fapi.client.GetResultsByYearAndRoundWithResponse(ctx, year, round, &ergast.GetResultsByYearAndRoundParams{})
	})
}

func (fapi *FormulaAPI) GetResultsByYear(ctx context.Context, year string) (*ergast.ResultsByYearResponse, error) {
	return ergast.CallErgast[ergast.ResultsByYearResponse](ctx, func() (res *ergast.GetResultsByYearResponse, err error) {
		return fapi.client.GetResultsByYearWithResponse(ctx, year, &ergast.GetResultsByYearParams{})
	})
}

func (fapi *FormulaAPI) GetResultsByYearAndConstructor(ctx context.Context, year, constructor string) (*ergast.ResultsByYearResponse, error) {
	return ergast.CallErgast[ergast.ResultsByYearResponse](ctx, func() (res *ergast.GetResultsByYearAndConstructorResponse, err error) {
		return fapi.client.GetResultsByYearAndConstructorWithResponse(ctx, year, constructor, &ergast.GetResultsByYearAndConstructorParams{})
	})
}
