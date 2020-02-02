package mongo

import "github.com/marshyon/codeStructure"

type Db map[int]architecture.Person

func (pg Db) Save(n int, p architecture.Person) {
	pg[n] = p
}

func (pg Db) Retrieve(n int) architecture.Person {
	return pg[n]
}
