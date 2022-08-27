package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

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

	res_id := flag.Int("res", 10, "Resource Id")

	flag.Parse()

	t, _ := tran.ExtractResourceTransactions(*res_id)
	for _, v := range t {
		fmt.Println(v)
	}

	res, err := tran.GetAllResourceCounts()

	if err != nil {
		fmt.Printf("Error %w", err)
		os.Exit(3)
	}
	var aero tran.ResourceUnit

	for _, v := range res {
		if v.Kind.Db_letter == 100 {
			aero = v
		}
	}

	db, err := sql.Open("sqlite3", "./transactions.db")
	if err != nil {
		fmt.Printf("Error %w", err)
		os.Exit(1)
	}

	defer db.Close()

	stmt := "insert into aerospace_research_stock values( ?, ?, ?, ?, ?, ?, ?, ?, ? )"

	_, err = db.Exec(stmt, aero.Id, time.Now().String(), aero.Amount, aero.Quality, aero.Cost.Workers, aero.Cost.Admin, aero.Cost.Material1, aero.Cost.Material2, aero.Cost.Market)
	// fmt.Println(t)
	if err != nil {
		fmt.Printf("Error %w", err)
		os.Exit(2)
	}

}
