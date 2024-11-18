package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) {
	// Convert to JSON with indentation
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(string(b))
}

func ByteToJson(input []byte) (string, error) {
	var data interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return "", err
	}

	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}
