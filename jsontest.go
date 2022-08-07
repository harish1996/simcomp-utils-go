package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Message struct {
	Id int64
	/* The field name has to start with a capital letter, otherwise the field wont be exported ( wont be visible to external libraries. ) */
	Kind interface{}
}

func main() {
	b := []byte(`[
	    {"id": 38234935, "kind": 42, "quantity": 775, "quality": 0, "price": 4.5, "seller": {"id": 2383955, "company": "Asian LTD", "realmId": 0, "logo": "", "certificates": 0, "contest_wins": 0, "npc": false, "courseId": null, "ip": "5a433d"}, "posted": "2022-07-31T10:20:06.608723+00:00", "fees": 105},
	    {"id": 38234643, "kind": 42, "quantity": 14485, "quality": 2, "price": 4.55, "seller": {"id": 645907, "company": "PREMIER.CO", "realmId": 0, "logo": "https://d1fxy698ilbz6u.cloudfront.net/logo/509bb009d70092ab77d879b7548ad603489ed1f8.png", "certificates": 1, "contest_wins": 0, "npc": false, "courseId": null, "ip": "3ba40e"}, "posted": "2022-07-31T10:14:48.839992+00:00", "fees": 0},
		{"id": 38234644, "kind": "43", "quantity": 14485, "quality": 2, "price": 4.55, "seller": {"id": 645907, "company": "PREMIER.CO", "realmId": 0, "logo": "https://d1fxy698ilbz6u.cloudfront.net/logo/509bb009d70092ab77d879b7548ad603489ed1f8.png", "certificates": 1, "contest_wins": 0, "npc": false, "courseId": null, "ip": "3ba40e"}, "posted": "2022-07-31T10:14:48.839992+00:00", "fees": 0}
	    ]`)

	var m []Message

	json.Unmarshal(b, &m)

	for i, v := range m {
		if _, ok := v.Kind.(string); ok {
			m[i].Kind, _ = strconv.ParseInt(v.Kind.(string), 10, 16)
			m[i].Kind = int16(m[i].Kind.(int64))
		}
		if _, ok := v.Kind.(float64); ok {
			m[i].Kind = int16(v.Kind.(float64))
		}
		fmt.Printf("%T %d\n", v.Kind, m[i].Kind.(int16)+5)
	}
	fmt.Println(m)
}
