package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type DBWrapper struct {
	*sql.DB
}

var (
	ErrDBExist      = fmt.Errorf("database file already exists")
	ErrTableNoExist = fmt.Errorf("table doesnt exist")
)

func CreateDBFile(name string) error {

	files, err := os.Open(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			files, err = os.Create(name)
			if err != nil {
				return fmt.Errorf("file create failed due to %w", err)
			}
		} else {
			return fmt.Errorf("file open failed due to %w", err)
		}
	} else {
		files.Close()
		return ErrDBExist
	}

	return nil
}

func CreateDBFileQuiet(name string) error {
	err := CreateDBFile(name)
	if err == ErrDBExist {
		err = nil
	}
	return err
}

func OpenDB(name string) (db *DBWrapper, err error) {
	db = &DBWrapper{}
	db.DB, err = sql.Open("sqlite3", name)
	return
}

func (db *DBWrapper) IsTableExists(name string) error {

	var n string

	query := fmt.Sprintf("select name from pragma_table_list where name=\"%s\" ", name)
	result, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("query failed due to %w", err)
	}

	defer result.Close()

	if result.Next() {
		err = result.Scan(&n)
		if err != nil {
			return fmt.Errorf("scan failed due to %w", err)
		}

		if n == name {
			return nil
		}
	}

	return ErrTableNoExist
}

type DataAndType struct {
	Name string
	Type string
}

func (db *DBWrapper) CreateTable(name string, schema *[]DataAndType) error {

	var result string
	for _, v := range *schema {
		if result == "" {
			result = fmt.Sprintf("%s %s", v.Name, v.Type)
			continue
		}
		result = result + fmt.Sprintf(", %s %s ", v.Name, v.Type)
	}

	query := fmt.Sprintf("create table %s ( %s )", name, result)
	fmt.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("exec failed due to %w", err)
	}
	return nil
}
