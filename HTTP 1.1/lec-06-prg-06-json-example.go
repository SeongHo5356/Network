package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	superHeroes_Source := map[string]interface{}{
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

	superHeroes_Source_mid, err := json.MarshalIndent(superHeroes_Source, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var superHeroes map[string]interface{}
	err = json.Unmarshal(superHeroes_Source_mid, &superHeroes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(superHeroes["homeTown"])
}
