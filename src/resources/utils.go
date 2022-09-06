package resources

import (
	"database/sql"
	"fmt"
	"strings"

	dbtools "hg33.com/simcomp/db"
	hu "hg33.com/simcomp/utils"
)

var (
	ErrEmptyDB = fmt.Errorf("the Database is NIL")
	ErrTableExist = fmt.Errorf("the resource_list table is already present")
)


type Material struct {
	Name           string
	Db_letter      int
	Transportation int
	Retailable     bool
	Research       bool
}

func GetMaterialList() (m []Material, err error) {
	url := "https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/"
	err = hu.NoAuthFetchFromJSONData(url, &m)
	return
}

func CanonicalizeName(in string) string {
	return strings.Trim(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(in), " ", "_"), "-", "_"), " \n\t\r")
}

func createMaterialListTable( db *dbtools.DBWrapper ) error {
	schema := []dbtools.DataAndType{ 
		{"name","text"}, 
		{"res_id","integer"}, 
		{"transport","real"}, 
		{"retailable","integer"}, 
		{"research","integer"}
	}

	return db.CreateTable( "resource_list", &schema )
	
}

/* Outside facing function of createMaterialListTable with error checking and all the good stuff. */
func CreateMaterialListTable( db *dbtools.DBWrapper ) error {
	if db == nil {
		return ErrEmptyDB
	}
	if db.IsTableExists("resource_list") == nil {
		return ErrTableExist
	} 

	return createMaterialListTable( db )
}

func addSingleMaterial( db *dbtools.DBWrapper, m Material ) error {
	query:="insert into resource_list values ( ?, ?, ?, ?, ? )"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("Error while preparing %w",err)
	}
	defer stmt.Close()

	result, err := stmt.Exec( m.Name, m.Db_letter, m.Transportation, m.Retailable, m.Research )
	if err != nil {
		return fmt.Errorf("Error while Executing %w", err)
	}
}

func UpdateMaterialListTable( db *dbtools.DBWrapper ) error {
	if db == nil {
		return ErrEmptyDB
	}
	
	if errors.Is(db.IsTableExists("resource_list"),dbtools.ErrTableNoExist) {
		err := createMaterialListTable( db )
		if err != nil {
			return fmt.Errorf("error while creating material list table %w",err)
		}		
	}

	list, err := GetMaterialList()
	if err != nil {
		return fmt.Errorf("error while fetching material list %w",err)
	}

	var statements map[string]*sql.Stmt

	gen_update_queries_text := func (col string) string {
		return fmt.Sprintf("update resource_list set %s=\"?\" where res_id=?",col)
	}
	gen_update_queries_number := func (col string) string {
		return fmt.Sprintf("update resource_list set %s=? where res_id=?",col)
	}
	var queries map[string]string {
		"find" : "select * from resource_list",
		"insert" : "insert into resource_list values ( ?, ?, ?, ?, ? )",
		"update_name" : gen_update_queries_text("name"),
		"update_transport" : gen_update_queries_number("transport"),
		"update_retailable" : gen_update_queries_number("retailable"),
		"update_research" : gen_update_queries_number("research"),
	}


	for k,v := range queries {
		statements[k], err := db.Prepare( v )
		if err != nil {
			return fmt.Errorf("error while preparing %s statement : %w",v,err)
		}

		defer statements[k].Close()
	}

	
	rows, err := statements["find"].Query()
	if err != nil {
		return fmt.Errorf("error while querying resource_list table %w",err)
	}

	defer rows.Close()

	/* Extracting all the materials from the list and storing it in a map to check later */
	var m Material
	var mmap map[int]Material

	for rows.Next() {
		err := rows.Scan( &m.Name, &m.Db_letter, &m.Transportation, &m.Retailable, &m.Research )
		if err != nil {
			return fmt.Errorf("error while scanning : %w",err)
		}

		mmap[m.Db_letter] = m
	}









}