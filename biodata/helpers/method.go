package helpers

import (
	"fmt"
	"reflect"
)

// Method to print Peserta with format of field: value
func (p Peserta) PrintPeserta() {
	// Get the data values
	v := reflect.ValueOf(p)
	// Get the type of struct
	t := v.Type()
	// Loop through the number of field for the struct
	for i := 0; i < v.NumField(); i++ {
		// Get field from the type
		field := t.Field(i)
		// Get value for the current field
		value := v.Field(i)
		// Print the name of the field and its value
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
		// value.Interface(): Get the value of the field as an interface{}
		// This allows printing the actual value and not reflect.Value representation
	}

}
