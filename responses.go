package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	type blk struct {
		Kind     int16
		Id       int64
		Quantity int32
		Price    float32
		Seller   struct {
			Company string
		}
	}

	r, err := http.Get("https://www.simcompanies.com/api/v3/market/0/42/")
	if err == nil {
		var f []blk
		js := json.NewDecoder(r.Body)
		js.Decode(&f)

		fmt.Println(f)

	} else {
		fmt.Println(err)
	}
}
