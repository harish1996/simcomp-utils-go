package resources

import (
	"strings"

	hu "hg33.com/simcomp/utils"
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
