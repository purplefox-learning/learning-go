package main

import (
	"encoding/json"
	"fmt"
)

type Tenant0 struct {
	ID   int64  `json:"id"` //`json:"xxx" someOtherTag:"yyy"` is called struct tag
	Name string `json:"name"`
}

type Brand0 struct {
	ID              int64  `json:"id,string"`
	Name            string `json:"name"`
	ExportedField   string
	unexportedField string
}

//referred by the comment below
type MyStruct struct {
	SomeField       string `json:"some_field,omitempty"`
	NeverGoOutField string `json:"-"`
}

func main() {

	fmt.Printf("\nhappy marshalling... \n\n")

	t := Tenant0{
		ID:   0,
		Name: "some_tenant_name",
	}

	b := Brand0{
		ID:              0,
		Name:            "some_brand_name",
		ExportedField:   "some_value",
		unexportedField: "some_other_value",
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
	fmt.Printf("annotating as json:xxxName,string on an integer typed field enables us to \n")
	fmt.Printf("- convert a integer field in its string format in marshalling \n")
	fmt.Printf("- parse a string back to integer in unmarshalling \n")
	fmt.Printf("but this annotation is limited, \n")
	fmt.Printf("- it can not marshal a []int64 type or other complex types where subset of fields to be stringfied in json \n")
	fmt.Printf("- even worse it failed to unmarshal a field with an empty string value, it throws an error \n\n")

	fmt.Printf("on a side note, besides the 'string' flag, there is also a ',omitempty' supported by json package \n")
	fmt.Printf("which allows json to skip a field while marshalling if it has a zero-value \n")
	fmt.Printf(`if some_field=="" in MyStruct, with omitempty the JSON value would be {}，without omitempty the JSON value would be {"some_field": ""}`)

	fmt.Printf("\n\n")
	fmt.Printf(`the json:"-" tag is used to skip marshalling a particular field in all cases`)

	fmt.Printf("\n\n")
	fmt.Printf("last but not least, and more fundamentally, unexported field will never get marshalled in any case \n")
}
