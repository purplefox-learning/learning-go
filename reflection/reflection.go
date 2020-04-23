package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Tenant struct {
	ID        int64  `json:"id"`
	ArticleID int64  `json:"article_id"`
	count     int64  `json:"count"`
	Name      string `json:"name"`
}

func main() {

	tt := Tenant{
		ID:        1,
		ArticleID: 2333,
		count:     6666,
		Name:      "some_name",
	}

	fmt.Printf("printing each field through reflection \n\n")

	vCopy := reflect.ValueOf(&tt).Elem()
	tCopy := vCopy.Type()

	for i := 0; i < tCopy.NumField(); i++ {
		fmt.Printf("Num Field #%v name: %v \n", i, tCopy.Field(i).Name)

		tag := tCopy.Field(i).Tag
		fmt.Printf("--- it has tags: %v \n", tag)

		//if we find the "json" annotation tag...
		if tagValue, ok := tag.Lookup("json"); ok {
			//if the tag has a value of 'id' or any other values containing 'id'...
			if strings.Contains(tagValue, "id") {
				fmt.Printf("--- found a tag named json and with 'id' value \n")
				//we can get the int value, and manipulate it anyway we want
				//int64field := vCopy.Field(i).Int()
				//stringfield := strconv.FormatInt(int64field, 10)
			}
		}
	}
}
