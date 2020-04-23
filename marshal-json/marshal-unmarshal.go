package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type FruitBasket struct {
	Id              int64 `json:"ref"` //check out the effect of "ref"
	Name            string
	Fruit           []string
	Created         time.Time
	unexportedField string //unexported field is not encoded.
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func main() {

	//marshaling of simple types

	aNumber, _ := json.Marshal(3.14)
	fmt.Println(string(aNumber))

	aString, _ := json.Marshal("pie")
	fmt.Println(string(aString))

	someFruits := []string{"apple", "orange", "pear"}
	fruitsJson, _ := json.Marshal(someFruits)
	fmt.Println(string(fruitsJson))

	someMap := map[string]int{"apple": 5, "orange": 7, "pear": 9}
	mapJson, _ := json.Marshal(someMap)
	fmt.Println(string(mapJson))
	fmt.Println("\n")

	//marshaling of struct obj

	basket := FruitBasket{
		Id:              999,
		Name:            "Standard",
		Fruit:           []string{"Apple", "Banana", "Orange"},
		Created:         time.Now(),
		unexportedField: "some_secret",
	}

	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	jsonData, err = json.MarshalIndent(basket, "", "    ")
	fmt.Println(string(jsonData))
	fmt.Println("\n")

	//marshalling of complex structures

	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Bitcoin in a Nutshell", Author: author}
	fmt.Printf("book in struct: %+v \n", book)
	bookBytes, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("book in json: %v \n", string(bookBytes))

	//unmarshal

	jsonData = []byte(`
		{
			"ref": 999,
			"Name": "Standard",
			"Fruit": [
				"Apple",
				"Banana",
				"Orange"
			],
			"Created": "2011-01-01T11:11:11Z"
		}`)

	var basketParsed FruitBasket
	err = json.Unmarshal(jsonData, &basketParsed)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basketParsed.Name, basketParsed.Fruit, basketParsed.Id)
	fmt.Println(basketParsed.Created)
	fmt.Println("\n")

	//unmarshal to a generic type
	//The encoding/json package unmarshal any valid JSON data into a plain interface{} value
	//- map[string]interface{} to store arbitrary JSON objects
	//- []interface{} to store arbitrary JSON arrays

	jsonData = []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var genericData map[string]interface{}
	json.Unmarshal(jsonData, &genericData)

	//k - string; v - interface{}
	for k, v := range genericData {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}

}
