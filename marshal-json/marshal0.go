package main

import (
	"encoding/json"
	"fmt"
)

type Tenant0 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Brand0 struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

func main() {

	fmt.Printf("\nhappy marshalling... \n\n")

	t := Tenant0{
		ID:   0,
		Name: "some_tenant_name",
	}

	b := Brand0{
		ID:   0,
		Name: "some_brand_name",
	}

	tjson, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("marshaling error is %v", err)
	}
	fmt.Printf("marshalled json: %v \n\n", string(tjson))

	bjson, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("marshaling error is %v", err)
	}
	fmt.Printf("marshalled json: %v \n\n", string(bjson))

	fmt.Printf("note the differnt in json output, there is a string double quotes on brand2's id \n")
	fmt.Printf("annotating with string on an integer type enables us to convert a int field in its string format in marshalling \n")
	fmt.Printf("but this annotation is limited, it does not help in marshalling a []int64 type \n")
	fmt.Printf("even worse it failed to unmarshal a json and throws an error when some one feed us a json with an empty string id \n")

}
