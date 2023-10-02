package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func Channa(src map[string]interface{}) string {
	var dest = make(map[string]interface{})

	helper(src, dest, "")

	b, _ := json.MarshalIndent(dest, " ", " ")

	return string(b)
}

func helper(src, dest map[string]interface{}, prefix string) {
	for k, v := range src {
		switch reflect.ValueOf(v).Kind() {
		case reflect.Map:
			if val, ok := v.(map[string]interface{}); ok {
				helper(val, dest, prefix+k+"__")
			}
		case reflect.Slice:
			switch child := v.(type) {
			case []interface{}:
				for i, val := range child {
					value := fmt.Sprintf("%v", val)
					if _, ok := val.(string); ok {
						value = fmt.Sprintf("\"%v\"", val)

					}

					dest[prefix+k+"__"+strconv.Itoa(i)] = value
				}
			}
		default:
			value := fmt.Sprintf("%v", v)
			if _, ok := v.(string); ok {
				value = fmt.Sprintf("\"%v\"", v)

			}

			dest[prefix+k] = value
		}
	}
}
