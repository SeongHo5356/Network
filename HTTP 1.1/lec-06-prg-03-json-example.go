package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("lec-06-prg-03-json-example.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	var superHeroes map[string]interface{}
	dec := json.NewDecoder(file)
	err = dec.Decode(&superHeroes)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(superHeroes["homeTown"])
	fmt.Println(superHeroes["active"])
	fmt.Println(superHeroes["members"].([]interface{})[1].(map[string]interface{})["powers"].([]interface{})[2])
}
