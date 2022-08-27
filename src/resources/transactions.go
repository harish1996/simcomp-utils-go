package resources

import (
	"fmt"

	hu "hg33.com/simcomp/utils"
)

type Transaction struct {
	Id       int
	Datetime string
	Category string
	Amount   int
	Quality  int
	Kind     struct {
		Db_letter int
	}
	Cost       float32
	Details    map[string]interface{}
	Otherparty struct {
		Id      int
		Company string
	}
}

func ExtractResourceTransactions(resource_id int) (t []Transaction, err error) {
	url := fmt.Sprintf("https://www.simcompanies.com/api/v2/resources-transactions/%d/0/", resource_id)
	err = hu.AuthFetchFromJSONData(url, &t)
	return
}
