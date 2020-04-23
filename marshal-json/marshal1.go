package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Int64Str int64

func (i Int64Str) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(i), 10))
}

func (i *Int64Str) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		// unmarshal as nil for a given empty string
		if s == "" {
			i = nil
			return nil
		}
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*i = Int64Str(value)
		return nil
	}

	// fallback to number
	return json.Unmarshal(b, (*int64)(i))
}

//this func is not used here, but this kind of convertion is often helpful in real projects
//to prepare a []int64 instance in the form of []Int64Str for final output in json
func convertInt64ToInt64Str(nums []int64) []Int64Str {
	numstr := []Int64Str{}
	for _, n := range nums {
		numstr = append(numstr, Int64Str(n))
	}
	return numstr
}

type Tenant1 struct {
	ID   Int64Str `json:"id"`
	Name string   `json:"name"`
}

type Brand1 struct {
	ID       Int64Str   `json:"id"`
	Name     string     `json:"name"`
	Products []Int64Str `json:"products"`
}

func main() {

	fmt.Printf("\nhappy marshalling... \n\n")

	t := Tenant1{
		ID:   0,
		Name: "some_tenant_name",
	}

	b := Brand1{
		ID:       0,
		Name:     "some_brand_name",
		Products: []Int64Str{Int64Str(1), Int64Str(2), Int64Str(3), Int64Str(4), Int64Str(5)},
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

	unmarshalledB := &Brand1{}
	json.Unmarshal(bjson, unmarshalledB)
	fmt.Printf("unmarshalled structs: %v \n\n", unmarshalledB)
	for _, pid := range unmarshalledB.Products {
		fmt.Printf("product id: %v \n", pid)
	}

	fmt.Printf("note that here we dont need the 'string' annotation as illustrated in marshal0.go \n")
	fmt.Printf("all numbers marked as Int64Str are automatically encoded as strings")

	fmt.Printf("annotating with string on an integer type enables us to convert a int field in its string format in marshalling \n")
	fmt.Printf("this method also support []int64 type well, both marshaling and unmarshaling \n")
}
