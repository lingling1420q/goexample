package main

import (
	"log"
)

type Person struct {
	Name string
}

type City struct {
	person *Person
}

var (
	flag   string
	person *Person
	//city   = &City{person}
)

func main() {
	//log.Println(city)
	log.Println("person", person)
	if person == nil {
		person = &Person{"hello word"}
	}
	city := &City{person}
	log.Printf("person %v\n", person)
	//city.person = person
	log.Printf("city %v\n", city.person)
	person.Name = "hell yan"
	log.Printf("city %v\n", city.person)
	person = nil
	log.Printf("city %v\n", city.person)
	flag = "yan"
	log.Printf("flag %v\n", flag)

	flag = "gao"

}
