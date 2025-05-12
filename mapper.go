package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Address struct {
	city   string
	street string
}

func (a Address) GetCity() string {
	return a.city
}

func (a Address) GetStreet() string {
	return a.street
}

func (a *Address) SetCity(c string) {
	a.city = c
}

func (a *Address) SetStreet(s string) {
	a.street = s
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,string"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func (a *Address) UnmarshalJSON(data []byte) error {
	var addrStr string
	if err := json.Unmarshal(data, &addrStr); err != nil {
		return err
	}
	parts := strings.SplitN(addrStr, ",", 2)
	if len(parts) == 2 {
		a.city = strings.TrimSpace(parts[0])
		a.street = strings.TrimSpace(parts[1])
	}
	return nil
}

func toJSONMap(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	tp := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := tp.Field(i)
		jsonTag := fieldType.Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]

		if jsonName == "" || jsonName == "-" {
			continue
		}

		if fieldType.Type.Name() == "Address" {
			address := field.Interface().(Address)
			result[jsonName] = map[string]interface{}{
				"city":   address.GetCity(),
				"street": address.GetStreet(),
			}
		} else {
			result[jsonName] = field.Interface()
		}
	}
	return result
}

func main() {
	jsonStr := `{
		"name": "Vasya",
		"age": "34",
		"email": "vasya@example.com",
		"address": "Spb, MyStreet"
	}`

	var person Person
	if err := json.Unmarshal([]byte(jsonStr), &person); err != nil {
		panic(err)
	}

	result := toJSONMap(&person)
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
