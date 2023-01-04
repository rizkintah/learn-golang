package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func (s Student) SetName(inputName string) {
	fmt.Println("Please type your name", inputName)
	s.Name = inputName
}
