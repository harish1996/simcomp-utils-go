package resources

import (
	"fmt"
	"time"

	hu "hg33.com/simcomp/utils"
)

type Prices struct {
	Id       int
	Kind     int
	Quantity int
	Quality  int
	Price    float32
	Seller   struct {
		Id      int
		Company string
	}
	Posted string
	Fees   float32
}

type PriceList struct {
	Time  string
	Kind  int
	Price []float32
}

func ExtractResourcePrices(res_id int) (prices []Prices, err error) {
	url := fmt.Sprintf("https://www.simcompanies.com/api/v3/market/all/0/%d/", res_id)
	err = hu.NoAuthFetchFromJSONData(url, &prices)
	return
}

func getPriceListfromPrices(prices []Prices) (list PriceList, err error) {
	list.Kind = prices[0].Kind
	list.Time = time.Now().String()

	for _, v := range prices {
		for len(list.Price) <= v.Quality+1 {
			list.Price = append(list.Price, v.Price)
		}
	}

	return
}

func ExtractPriceList(res_id int) (list PriceList, err error) {
	prices, err := ExtractResourcePrices(res_id)
	if err != nil {
		err = fmt.Errorf("Extracting resource price failed %w", err)
		return
	}

	list, err = getPriceListfromPrices(prices)
	if err != nil {
		err = fmt.Errorf("Price list conversion failed %w", err)
		return
	}

	return
}
