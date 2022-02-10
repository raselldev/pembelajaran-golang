package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	raka := Man{"Raka"}
	raka.Married()
	fmt.Println(raka.Name)

	rasell := Man{"Rasell"}
	rasell.MarriedPointer()
	fmt.Println(rasell.Name)
}
