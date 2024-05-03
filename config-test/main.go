package main

import (
	"encoding/json"
	"fmt"
	"os"
	_ "os/exec"
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
	// fmt.Println(oasis)

	build, ok := oasis["build"].(map[string]interface{})
	if !ok {
		panic("build is not a map[string]interface{}")
	}

	mode, ok := build["mode"].(string)
	if !ok {
		panic("mode is not a string")
	}

	fmt.Println("mode:", mode)
}
