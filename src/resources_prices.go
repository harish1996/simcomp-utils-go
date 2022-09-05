package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	dbtools "hg33.com/simcomp/db"
	res "hg33.com/simcomp/resources"
)

func main() {
	db_file := flag.String("db", "resources.db", "Resource database path")
	// res_id := flag.Int("res", 10, "Resource Id")

	flag.Parse()

	// t, _ := res.ExtractPriceList(*res_id)
	t, _ := res.GetMaterialList()

	err := dbtools.CreateDBFileQuiet(*db_file)
	if err != nil {
		fmt.Println("Create DB File failed due to %w", err)
		os.Exit(1)
	}

	db, err := dbtools.OpenDB(*db_file)

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
		isexist := db.IsTableExists(v.Name)
		if isexist != nil {
			fmt.Printf("%s doesnt exist \n %w", v.Name, isexist)
		} else {
			fmt.Println(v.Name, "Exists")
		}
	}

	sche := []dbtools.DataAndType{{"name", "text"}, {"q0", "real"}}

	_ = db.CreateTable("newtable", &sche)

	db.Close()
}
