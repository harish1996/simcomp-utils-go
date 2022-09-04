package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	res "hg33.com/simcomp/resources"
)

func main() {
	db_file := flag.String("db", "resources.db", "Resource database path")
	// res_id := flag.Int("res", 10, "Resource Id")

	flag.Parse()

	// t, _ := res.ExtractPriceList(*res_id)
	t, _ := res.GetMaterialList()

	// fmt.Printf("%v", t)

	files, err := os.Open(*db_file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			files, err = os.Create(*db_file)
			if err != nil {
				fmt.Printf("Cant create %s due to %w", *db_file, err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Something else")
			os.Exit(2)
		}
	}
	files.Close()

	db, err := sql.Open("sqlite3", *db_file)

	// for _, v := range t {
	// 	v.Name = strings.Trim(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(v.Name), " ", "_"), "-", "_"), " \n\t\r")
	// 	_, err := db.Exec(fmt.Sprintf(" CREATE TABLE %s ( time string, q0 real, q1 real, q2 real, q3 real, q4 real, q5 real, q6 real, q7 real, q8 real )", v.Name))

	// 	if err != nil {
	// 		fmt.Println("Table creation failed at %s with error %w", v.Name, err)
	// 		os.Exit(1)
	// 	}

	// }

	for _, v := range t {
		v.Name = res.CanonicalizeName(v.Name)
		result, err := db.Query(fmt.Sprintf("pragma table_info('%s')", v.Name))
		if err != nil {
			fmt.Println("Query failed for some reason %w", err)
			os.Exit(2)
		}

		for result.Next() {
			var arg1, arg2, arg3, arg4, arg5, arg6 interface{}
			err = result.Scan(&arg1, &arg2, &arg3, &arg4, &arg5, &arg6)
			if err != nil {
				fmt.Println("Scan failed due to %w", err)
				os.Exit(3)
			}
			fmt.Println(arg1, arg2, arg3, arg4, arg5, arg6)
		}

		result.Close()
	}

	db.Close()
}
