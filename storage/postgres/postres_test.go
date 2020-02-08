package postgres

import (
	"testing"

	architecture "github.com/marshyon/codeStructure"
)

// DbMock map of Persons by integer key
type DbMock map[int]architecture.Person

// Save method for postgres backend
func (m DbMock) Save(n int, p architecture.Person) {
	m[n] = p
}

// Retrieve method for postgres backend
func (m DbMock) Retrieve(n int) architecture.Person {
	return m[n]
}
func TestSomething(t *testing.T) {
	dbmMock := DbMock{}

	p1 := architecture.Person{
		First: "Jane",
	}

	p2 := architecture.Person{
		First: "John",
	}

	ps := architecture.NewPersonService(dbmMock)
	ps.Save(1, p1)
	ps.Save(2, p2)

	name, err := ps.Get(1)
	if err != nil {
		t.Errorf("error returned storing %s : %s\n", name, err)
	}
	name, err = ps.Get(2)
	if err != nil {
		t.Errorf("error returned storing %s : %s\n", name, err)
	}
	name, err = ps.Get(3)
	if err == nil {
		t.Errorf("something went wrong retrieving a non existen record %s : %s\n", name, err)
	}
}
