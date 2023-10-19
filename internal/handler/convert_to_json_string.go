package handler

import (
	"encoding/json"
	"log"
)

// ToJsonString converts struct to JSON string
func ToJsonString(i interface{}) string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}
	return string(jsonBytes)
}

// ToStruct converts JSON string to struct
func ToStruct(s string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		log.Fatalf("Error unmarshaling: %v", err)
	}
	return i
}
