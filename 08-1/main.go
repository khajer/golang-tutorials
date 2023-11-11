package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Home string `json:"home"`
}
type MainJSON struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`

	Bill struct {
		Address string `json:"address"`
	} `json:"bill"`
}

func main() {
	fmt.Println("test json")
	bill := MainJSON{
		Id:   2,
		Name: "itsara",
		Address: Address{
			Home: "lampang",
		},
	}
	fmt.Println(bill)
	// jsonString, err := json.Marshal(bill) //jsonString, err := json.Marshal(bill)
	jsonString, err := json.MarshalIndent(bill, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonString))

	// convert back
	jsonData := `
	{
		"id":5,
		"name":"test",
		"address":{
			"home":"chiangmai"
		},
		"bill":{
			"address": "bkk"
		}
	}
	`
	fmt.Println("data json string -> struct")
	fmt.Println(jsonData)
	var dataJson1 MainJSON
	err = json.Unmarshal([]byte(jsonData), &dataJson1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(dataJson1)

}
