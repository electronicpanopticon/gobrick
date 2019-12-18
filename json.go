package gobrick

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ToJSONBytes takes in an entity and returns formatted JSON as a byte array.
func ToJSONBytes(v interface{}) []byte {
	b, err := json.MarshalIndent(v, "", "  ")
	if string(b) == "null" {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}

// ToJSON takes in an entity and returns a formatted JSON string.
func ToJSON(v interface{}) string {
	b := strings.TrimSpace(string(ToJSONBytes(v)))
	if b == "" || b == "\"\"" {
		return "{}"
	}
	return b
}
