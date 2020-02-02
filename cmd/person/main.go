package main

import (
	 "fmt"
	 "github.com/marshyon/codeStructue"
	 "github.com/marshyon/codeStructue/storage/mongo"
	 "github.com/marshyon/codeStructue/storage/postgres"
)

func main() {
	fmt.Println("hi")

	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Person{
		First: "Jane",
	}

	p2 := architecture.Person{
		First: "John",
	}

	p3 := architecture.Person{
		First: "Mellanie",
	}

	ps := architecture.NewPersonService(dbm)


	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.get(dbm, 2))

	ps.Save(3, p3)
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(2))
	fmt.Println(ps.Get(3))
	fmt.Println(ps.Get(4))

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}
