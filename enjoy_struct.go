package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

type Address struct {
	addrid string
	addr   string
}

type VCard struct {
	name    string
	address *Address
	birth   string
	image   []byte
}

func formatPerson(ps person) {
	ps.name = strings.ToUpper(ps.name)
	ps.age = int(ps.age)
}

func main() {
	// p1 is point
	p1 := new(person)

	// p1 is person
	// var p1 person

	// p1 is point
	// p1 := &person{"sdfh", 5}
	p1.name = "cc"
	p1.age = 10
	formatPerson(*p1)
	fmt.Println(p1)

	addr1 := new(Address)
	addr1.addr = "HZ"
	addr1.addrid = "001"
	card := new(VCard)

	card.birth = "2016-03-03 00:00:00"
	card.name = "aaa"
	card.address = addr1
	fmt.Printf("card: %v\n", card)
}
