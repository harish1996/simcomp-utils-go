package main

import (
	"encoding/json"
	"fmt"
	"os"

	hu "hg33.com/simcomp/utils"
)

type cashflow struct {
	Data []struct {
		Id          int
		Money       float32
		Description string
	}
}

func main() {
	c, err := hu.GetAuthenticatedSession()
	if err != nil {
		fmt.Printf("Authenticated session creation failed: %w", err)
		os.Exit(1)
	}

	resp, err := c.Httpget("https://www.simcompanies.com/api/v2/companies/me/cashflow/recent/", hu.Defaultheaders, nil, nil)
	if err != nil {
		fmt.Printf("HTTP get failed: %w", err)
		os.Exit(2)
	}

	var cf cashflow

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&cf)
	if err != nil {
		fmt.Printf("Readall failed \n %w", err)
		os.Exit(3)
	}

	fmt.Println(cf)
}
