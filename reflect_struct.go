package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	fieldA string
	fieldB int
	fieldC []string
	fieldD map[string]string
}

func main() {
	fmt.Println("Go Learninig for Reflect Struct")

	foo := &Foo{
		fieldA: "Value of FieldA",
		fieldB: 123,
		fieldC: []string{"Value", "Of", "FieldC"},
		fieldD: make(map[string]string),
	}

	fmt.Printf("%V\n", foo)

	reflectFoo := reflect.ValueOf(foo)
	fmt.Printf("%v\n", reflectFoo.Type())

	fmt.Printf("-- Show Name of Fields --\n")
	elemFoo := reflectFoo.Elem()
	fmt.Printf("%v\n", elemFoo)

	typeOfFoo := elemFoo.Type()
	fmt.Printf("%v\n", typeOfFoo)

	for i := 0; i < elemFoo.NumField(); i++ {
		f := elemFoo.Field(i)
		fmt.Printf("%d: %s %s\n", i,
			typeOfFoo.Field(i).Name, f.Type())
	}

	fmt.Printf("\n")
	fmt.Printf("\n")

	reflectVal := reflect.ValueOf(*foo)
	fmt.Printf("%v\n", reflectVal.Type())

	fmt.Printf("-- Show Value of Fields --\n")
	numFields := reflectVal.NumField()
	for idx := 0; idx < numFields; idx++ {
		fieldVal := reflectVal.Field(idx)
		fmt.Printf("%v\n", fieldVal.Type())
		fmt.Printf("%v\n", fieldVal)
		fmt.Printf("%v\n", fieldVal.Kind())
		fmt.Printf("\n")
	}

}
