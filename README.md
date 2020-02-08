# codeStructure

taken from a course by Todd McLeod on Udemy - Exploring the Go Programming Language

This is a pattern for using an interface not only to 'mock' a backend for testing
but also to abstract a database back end.

Dummy modules for Mongo and Postgress are available under 'storage/{$dbname}' directories
where each has it's own tests. Each implement the interace used to access the database 
storage backends. 

# Running the app

To run the app:

```
go run cmd/person/main.go
```

or 

```
cd cmd/person
go build
./person
```

# Testing

run tests ( from the application root folder ) with 

```
go test  ./...
```

