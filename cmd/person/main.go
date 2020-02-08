package main

import (
	 "fmt"
	 "github.com/marshyon/codeStructure"
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
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(2))
	fmt.Println(ps.Get(3))
}
