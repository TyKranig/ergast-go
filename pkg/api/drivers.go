package api

import (
	"context"

	"github.com/tylerkranig/ergast-api/pkg/ergast"
)

func (fapi *FormulaAPI) GetAllDrivers(ctx context.Context) (*ergast.DriversResponse, error) {
	return ergast.CallErgast[ergast.DriversResponse](ctx, func() (res *ergast.GetAllDriversResponse, err error) {
		return fapi.client.GetAllDriversWithResponse(ctx)
	})
}

func (fapi *FormulaAPI) GetDriversByYear(ctx context.Context, year string) (*ergast.DriversResponse, error) {
	return ergast.CallErgast[ergast.DriversResponse](ctx, func() (res *ergast.GetDriversByYearResponse, err error) {
		return fapi.client.GetDriversByYearWithResponse(ctx, year)
	})
}

func (fapi *FormulaAPI) GetDriverStandingsByYear(ctx context.Context, year string) (*ergast.DriverStandingsByYearResponse, error) {
	return ergast.CallErgast[ergast.DriverStandingsByYearResponse](ctx, func() (res *ergast.GetDriverStandingsByYearResponse, err error) {
		return fapi.client.GetDriverStandingsByYearWithResponse(ctx, year, &ergast.GetDriverStandingsByYearParams{})
	})
}
