package resources

import (
	"encoding/json"
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

	c, err := hu.GetAuthenticatedSession()
	if err != nil {
		err = fmt.Errorf("Authenticated session creation failed: %w", err)
		return t, err
	}

	c.AddHeader(hu.Defaultheaders)

	resp, err := c.Httpget(fmt.Sprintf("https://www.simcompanies.com/api/v2/resources-transactions/%d/0/", resource_id), nil)
	if err != nil {
		err = fmt.Errorf("HTTP get failed: %w", err)
		return t, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)
	if err != nil {
		err = fmt.Errorf("Readall failed \n %w", err)
		return t, err
	}

	return t, nil
}
