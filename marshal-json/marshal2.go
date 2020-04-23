package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type Tenant2 struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password" lark:"confidential"`
}

// MarshalJSON implements json.Marshaller.
func (t Tenant2) MarshalJSON() ([]byte, error) {
	// create a new type to prevent json.Marshal loop
	type tCopy Tenant2
	tCopyPointer := (*tCopy)(&t)
	return omitConfidentialFieldJSONMarshal(tCopyPointer)
}

type Brand2 struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	CEO  string `json:"ceo" lark:"confidential"`
}

// MarshalJSON implements json.Marshaller.
func (b Brand2) MarshalJSON() ([]byte, error) {
	// create a new type to prevent json.Marshal loop
	type bCopy Brand2
	bCopyPointer := (*bCopy)(&b)
	return omitConfidentialFieldJSONMarshal(bCopyPointer)
}

// ===== Helper Utils =====

// omitConfidentialFieldJSONMarshal only accepts struct pointer
// it removes certain internal field (with special custom tag "lark=staging_only" value) from json Marshal result
// works for int, string, pointer, map, slice field types
func omitConfidentialFieldJSONMarshal(structPointer interface{}) ([]byte, error) {
	t := reflect.TypeOf(structPointer)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return nil, errors.New("param structPointer is not a pointer to a struct")
	}

	vCopy := reflect.ValueOf(structPointer).Elem()
	tCopy := vCopy.Type()

	// set field to nil if field has staging_only tag
	for i := 0; i < tCopy.NumField(); i++ {
		tag := tCopy.Field(i).Tag
		//fmt.Printf("scanning tag %v \n",tag)
		//fmt.Printf("   json tag has value %v \n",tag.Get("json"))
		//fmt.Printf("   lark tag has value %v \n",tag.Get("lark"))
		if tagValue, ok := tag.Lookup("lark"); ok {
			fmt.Printf("removing confidential field value \n")
			if strings.Contains(tagValue, "confidential") {
				//uncomment the line below to see the difference
				vCopy.Field(i).Set(reflect.Zero(tCopy.Field(i).Type))
			}
		}
	}
	return json.Marshal(vCopy.Interface())
}

func main() {

	fmt.Printf("\nhappy marshalling... \n")

	t := Tenant2{
		ID:       1,
		Name:     "some_tenant_name",
		Password: "not_to_be_revealed",
	}

	b := Brand2{
		ID:   1,
		Name: "some_brand_name",
		CEO:  "not_to_be_revealed",
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

	fmt.Printf("this method goes further to directly control marshal/unmarshal on a specific struct \n")
	fmt.Printf("go allows us to tag anything on a struct, such as the common json:xxx tag. \n")
	fmt.Printf("we use reflection to find special tags, so that we can treat those special fields differntly during marshal/unmarshal operations \n")
	fmt.Printf("in our case, we use 'lark' tag to denote some information should not be marshalled (and sent to the caller in json form) \n")
}
