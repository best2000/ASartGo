package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type Config struct {
	Mode string `json:"Mode"`
	DefaultBrowser  string `json:"DefaultBrowser"`
}

func main() {
	/*set := Config{Mode: "def", DefaultBrowser: "ffx"}
    byteArray, _ := json.MarshalIndent(set,""," ")
	fmt.Println(string(byteArray))*/

	bf, _ := ioutil.ReadFile("config.json")
	//ioutil.WriteFile("config.json", byteArray, 0644)

	jsonString := string(bf)
	var set Config
	json.Unmarshal([]byte(jsonString), &set)
	fmt.Printf("%+v\n", set)
	if set.DefaultBrowser == "" {
		fmt.Println("plz set ur debrowse!")
	}

	

	
}

