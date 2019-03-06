package main

import (
	"encoding/json"
	"fmt"
)

func MapToJson() {

	map1 := make(map[string]interface{})
	map1["name"] = "martin"
	map1["age"] = 18

	//return []byte
	str, err := json.Marshal(map1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))

}

func JsonToMap() {
	jsonStr := `{"age":18,"name":"martin"}`

	map2 := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &map2)
	if err != nil {
		fmt.Println(err);
	}

	fmt.Println("json to map",map2)
}

func main() {

	MapToJson()
	JsonToMap()
}
