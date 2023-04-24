package main

import (
	"context"
	"fmt"

	"github.com/tylerkranig/ergast-api/pkg/api"
)

func main() {
	fapi, err := api.NewFormulaAPI()
	if err != nil {
		fmt.Printf("failed to create new client: %s", err.Error())
	}

	driverResponse, err := fapi.GetDriversByYear(context.Background(), "2004")
	if err != nil {
		fmt.Printf("failed to get drivers: %s", err.Error())
	}

	for _, driver := range *driverResponse.MRData.DriverTable.Drivers {
		fmt.Println(*driver.FamilyName, *driver.GivenName)
	}
}
