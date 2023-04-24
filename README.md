# ergast-go
golang wrapper for ergast formula-1 api

```go
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
```

# What is Ergast

Ergast is a free API for developers to access historic formula 1 data. This wrapper attempts to make an easy interface for golang developers to utilize formula 1 shenanigans.

# Installation

Just import the project into your golang project and run a ``go mod tidy``
```
import "github.com/tylerkranig/ergast-api/pkg/api"
```

# Makefile

There are a few make commands to help development
`make generate` ---> will regenerate the oapi-codegen definitions
`make clean` ---> delete all *.gen.go files, remove the oapi-codegen files
`make lint` ---> lint go code with golangci-lint as well as the yaml file for ergast with spectral
`make format` ---> formats go source code

# Usage

To call the api you first need to create a new Client struct
```go
	fapi, err := api.NewFormulaAPI()
	if err != nil {
		fmt.Printf("failed to create new client: %s", err.Error())
	}
```
Passing in a different url is optional, default is ``https://ergast.com/api/f1``
```go
	fapi, _ := api.NewFormulaAPI("{some other url}")
```

Then using the Client struct you can call any of the endpoints you need

# Credits

This package depends on the free api provided at [Ergast](http://ergast.com/mrd/). For full information about the API and it's responsible use, please refer to their website. [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen) was used to create a client wrapper from the yaml definition.
