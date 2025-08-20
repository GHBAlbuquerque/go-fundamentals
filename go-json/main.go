package main

import (
	"encoding/json"
	"log"
)

type Idol struct {
	StageName string `json:"stage_name"`
	RealName  string `json:"real_name"`
	HairColor string `json:"hair_color"`
	Group     string `json:"group"`
}

func main() {

	myJson := `
	[
		{
			"stage_name":"Karina",
			"real_name":"Yu Jimin",
			"hair_color":"black",
			"group":"aespa"
		}, 
		{
			"stage_name":"Winter",
			"real_name":"Kim Min-jeong",
			"hair_color":"blonde",
			"group":"aespa"
		},
		{
			"stage_name":"NingNing",
			"real_name":"Ning Yizhuo",
			"hair_color":"brown",
			"group":"aespa"
		},
		{
			"stage_name":"Giselle",
			"real_name":"Aeri Uchinaga",
			"hair_color":"pink",
			"group":"aespa"
		}
	]
	`

	// going from json to object
	var unmarshalled []Idol

	err := json.Unmarshal([]byte(myJson), &unmarshalled) //bytes + pointer to store

	if err != nil {
		log.Println("Error unmarshalling json", err.Error())
	}

	log.Println("Unmarshalled:", unmarshalled)

	// going from object to json
	idols := []Idol{}
	idols = append(idols, Idol{StageName: "Karina", RealName: "Yu Jimin", HairColor: "black", Group: "aespa"})
	idols = append(idols, Idol{StageName: "Winter", RealName: "Kim Min-jeong", HairColor: "blonde", Group: "aespa"})
	idols = append(idols, Idol{StageName: "NingNing", RealName: "Ning Yizhuo", HairColor: "brown", Group: "aespa"})
	idols = append(idols, Idol{StageName: "Giselle", RealName: "Aeri Uchinaga", HairColor: "pink", Group: "aespa"})

	newJson, err := json.MarshalIndent(idols, "", "") // marshal with identatition, good for debugging. otherwise use marshal

	if err != nil {
		log.Println("Error marshalling json", err.Error())
	}

	log.Println(string(newJson)) // needs to convert otherwise will print byte array
}
