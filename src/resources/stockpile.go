package resources

import (
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
	url := "https://www.simcompanies.com/api/v2/resources/"
	err = hu.AuthFetchFromJSONData(url, &t)
	return
}
