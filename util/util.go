package util

import "reflect"

// ToMap converts a struct to a map
func ToMap(s interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		// Check if the field is not a zero value
		if !field.IsZero() {
			result[fieldName] = field.Interface()
		}
	}
	return result
}
