package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" "max:10"`
}
type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	Sample := Sample{"Fachrul Aziz"}
	var sampleType = reflect.TypeOf(Sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("Max"))

	Sample.Name = ""
	fmt.Println(IsValid(Sample))

	contoh := ContohLagi{"", ""}
	fmt.Println(contoh)
}
