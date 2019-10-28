package gobrick

import (
	"encoding/json"
	"fmt"
)

func ToJSONBytes(v interface{}) []byte {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

func ToJSON(v interface{}) string {
	b := ToJSONBytes(v)
	if b == nil {
		return "{}"
	}
	return string(b)
}