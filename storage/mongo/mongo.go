package mongo

import "github.com/marshyon/codeStructure"

type Db map[int]architecture.Person

func (m Db) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m Db) retrieve(n int) architecture.Person {
	return m[n]
}
