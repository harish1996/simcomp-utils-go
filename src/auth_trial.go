package main

import (
	"fmt"

	tran "hg33.com/simcomp/resources"
)

type cashflow struct {
	Data []struct {
		Id          int
		Money       float32
		Description string
	}
}

func main() {
	// c, err := hu.GetAuthenticatedSession()
	// if err != nil {
	// 	fmt.Printf("Authenticated session creation failed: %w", err)
	// 	os.Exit(1)
	// }

	// c.AddHeader(hu.Defaultheaders)

	// resp, err := c.Httpget("https://www.simcompanies.com/api/v2/companies/me/cashflow/recent/", nil)
	// if err != nil {
	// 	fmt.Printf("HTTP get failed: %w", err)
	// 	os.Exit(2)
	// }

	// var cf cashflow

	// decoder := json.NewDecoder(resp.Body)
	// err = decoder.Decode(&cf)
	// if err != nil {
	// 	fmt.Printf("Readall failed \n %w", err)
	// 	os.Exit(3)
	// }

	// fmt.Println(cf)

	// res_id := flag.Int("res", 10, "Resource Id")

	// flag.Parse()

	// t, _ := tran.ExtractResourceTransactions(*res_id)
	// for _, v := range t {
	// 	ti, _ := time.Parse("2006-01-02T15:04:05.000000-07:00", v.Datetime)
	// 	ti2 := (ti.UnixNano() / 1000000) + 5402
	// 	fmt.Println(v.Datetime, time.Unix(ti2/1000, 0).UTC())
	// }

	res, _ := tran.GetAllResourceCounts()
	for _, v := range res {
		fmt.Println(v)
	}
	// fmt.Println(t)

}
