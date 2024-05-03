package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// open oasis.json
	file, err := os.Open("oasis.json")
	check(err)
	defer file.Close()

	// decode json
	decoder := json.NewDecoder(file)
	var oasis map[string]interface{}
	err = decoder.Decode(&oasis)
	check(err)

	// print json
	fmt.Println(oasis)
}
