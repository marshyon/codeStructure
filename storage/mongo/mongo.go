package mongo

import "github.com/marshyon/codeStructure"

// Db map of Persons by integer key
type Db map[int]architecture.Person

// Save method for mongo backend
func (m Db) Save(n int, p architecture.Person) {
	m[n] = p
}

// Retrieve method for mongo backend
func (m Db) Retrieve(n int) architecture.Person {
	return m[n]
}
