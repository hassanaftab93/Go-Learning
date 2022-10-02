package structs

/*
- What are structures?

A structure is a composite object
- Organizational purposes
- For example we need a struct "Person"
*/

type Person struct {
	name  string
	addr  string
	phone string
}

func Called() string {
	return "structs package called"
}

func InitializePerson(Name, Address, Phone string) Person {
	p1 := new(Person)
	//p1 := Person (name: "Joe",addr: "??",phone: "00092")
	p1.name = Name
	p1.addr = Address
	p1.phone = Phone

	return *p1
}
