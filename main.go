package main

import (
	"dualshockcharge/bcheck"
	"fmt"
	"log"
)

func main() {
	list, err := bcheck.FindControllers()
	if err != nil {
		log.Fatal(err)
	}

	for _, elem := range list {
		cstat, err := elem.Capacity()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Capacity:", cstat)
	}
}
