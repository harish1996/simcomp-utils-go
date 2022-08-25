package resources

import (
	"encoding/json"
	"fmt"

	hu "hg33.com/simcomp/utils"
)

type ResourceUnit struct {
	Id      int
	Amount  int
	Quality int
	Kind    struct {
		Db_letter int
	}
	Cost struct {
		Workers   float32
		Admin     float32
		Material1 float32
		Material2 float32
		Material3 float32
		Material4 float32
		Material5 float32
		Market    float32
	}
	Materials []string
}

func GetAllResourceCounts() (t []ResourceUnit, err error) {

	c, err := hu.GetAuthenticatedSession()
	if err != nil {
		err = fmt.Errorf("Authenticated session creation failed: %w", err)
		return t, err
	}

	c.AddHeader(hu.Defaultheaders)

	resp, err := c.Httpget("https://www.simcompanies.com/api/v2/resources/", nil)
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
