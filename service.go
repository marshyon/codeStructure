package architecture

import "fmt"

// Accessor interface is used to access and abstract storage back-ends
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

// Person struct is the main
// data componenent
type Person struct {
	First string
}

// PersonService uses accessor interface
type PersonService struct {
	a Accessor
}

// Get method used to access data through
// Person service and the Retrieve method
// the Retrieve method is implemented by the storage backend
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no person or persons with id of %d", n)
	}
	return p, nil
}

// Save method used to access data through
// Person service and the Save method
// the Save method is implemented by the storage backend
func (ps PersonService) Save(n int, p Person) {
	ps.a.Save(n, p)
}

// NewPersonService creates a new service to action
// save and retrieve operations
func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}
