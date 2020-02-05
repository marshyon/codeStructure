package architecture

import "fmt"

type Person struct {
	First string
}

type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}


func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no person or persons with id of %d", n)
	}
	return p, nil
}

func (ps PersonService) Save(n int, p Person) {
	ps.a.Save(n, p)
}
