package main

import (
	"fmt"

	"github.com/SoliDeoHonertGloria/mydict"
)

func main() {

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}
