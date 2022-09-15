package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_map := make(map[string]string)
	json_map["a"] = "a"
	json_map["b"] = "b"
	c, _ := json.Marshal(json_map)
	fmt.Println(string())

	fmt.Println(string(c))
}
