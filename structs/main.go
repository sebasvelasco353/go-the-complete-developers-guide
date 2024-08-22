package main

// NOTE: &variable gives me the address of the value the variable points at
// NOTE: *pointer gives me the value this pointer is referencing

import "fmt"

type contactInfo  struct {
  email string
  zipCode int
}

type person struct {
  firstName string
  lastName string
  contactInfo
}

func main() {
  alex := person{
    firstName: "Alex", 
    lastName: "Anderson",
    contactInfo: contactInfo{
      email: "test@email.com",
      zipCode: 12345,
    },
  }
  alex.updateName("pancho")
  alex.print()
}

func (pointerToPerson *person)updateName(newfirstName string) {
  (*pointerToPerson).firstName = newfirstName
}

func (p person)print() {
  fmt.Printf("%+v", p)
}
