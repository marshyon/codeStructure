package main

import (
	"fmt"

	architecture "github.com/marshyon/codeStructure"
	"github.com/marshyon/codeStructure/storage/mongo"
)

func main() {
	dbm := mongo.Db{}

	p1 := architecture.Person{
		First: "Jane",
	}

	p2 := architecture.Person{
		First: "John",
	}

	ps := architecture.NewPersonService(dbm)
	ps.Save(1, p1)
	ps.Save(2, p2)
	res, err := ps.Get(1)
	if err != nil {
		fmt.Printf("Error getting result for key 1 : %s\n", err)
	} else {
		fmt.Printf("got %s from key 1\n", res.First)
	}
	res, err = ps.Get(2)
	if err != nil {
		fmt.Printf("Error getting result for key 2 : %s\n", err)
	} else {
		fmt.Printf("got %s from key 2\n", res.First)
	}
	res, err = ps.Get(3)
	if err != nil {
		fmt.Printf("Error getting result for key 3 : %s\n", err)
	}
}
