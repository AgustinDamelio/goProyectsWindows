package main

import (
	"fmt"
)

func main() {
	newStudent1 := Student{}
	newStudent1.Name = "Agustin"
	newStudent1.LastName = "Damelio"
	newStudent1.Id = "123456"
	newStudent1.admissionDate = "10-12-2022"
	newStudent1.printStudent()
}

type Student struct {
	Name          string
	LastName      string
	Id            string
	admissionDate string
}

func (s Student) printStudent() {
	fmt.Printf("%s %s id: %s entered the school on the %s", s.Name, s.LastName, s.Id, s.admissionDate)
}
