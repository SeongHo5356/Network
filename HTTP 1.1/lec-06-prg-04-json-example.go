package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	superHeroes := map[string]interface{}{
		"squadName":  "Super hero squad",
		"homeTown":   "Metro City",
		"formed":     2016,
		"secretBase": "Super tower",
		"active":     true,
		"members": []map[string]interface{}{
			{
				"name":           "Molecule Man",
				"age":            29,
				"secretIdentity": "Dan Jukes",
				"powers": []string{
					"Radiation resistance",
					"Turning tiny",
					"Radiation blast",
				},
			},
			{
				"name":           "Madame Uppercut",
				"age":            39,
				"secretIdentity": "Jane Wilson",
				"powers": []string{
					"Million tonne punch",
					"Damage resistance",
					"Superhuman reflexes",
				},
			},
			{
				"name":           "Eternal Flame",
				"age":            1000000,
				"secretIdentity": "Unknown",
				"powers": []string{
					"Immortality",
					"Heat Immunity",
					"Inferno",
					"Teleportation",
					"Interdimensional travel",
				},
			},
		},
	}

	fmt.Println(superHeroes["homeTown"])
	fmt.Println(superHeroes["active"])
	fmt.Println(superHeroes["members"].([]map[string]interface{})[1]["powers"].([]string)[2])

	file, err := os.Create("lec-06-prg-04-json-example.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(superHeroes)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

}
