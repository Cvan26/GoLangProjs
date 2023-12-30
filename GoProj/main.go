package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type fashion struct {
	top    string
	bottom string
}

type people interface {
	stdin() string
}

func main() {

	// var cvan person
	// cvan.firstname = "anh"
	// cvan.lastname = "viet"
	// fmt.Println(cvan)
	// cvanPointer := &cvan
	// cvanPointer.UpdateFirstName("Ngoc")
	// fmt.Printf("%+v", cvan)
	// fmt.Println(&cvanPointer)
	// colors := make(map[string]string)
	// // colors["Red"] = "white"
	// colors = map[string]string{
	// 	"Red":    "ff0000",
	// 	"greeen": "#4bf745",
	// }

	// // fmt.Println(colors["Red"])
	// printMap(colors)
	// nguoi1 := person{}
	nguoi2 := "dsadsađâs"
	// describePeople(nguoi1)
	describePeople(nguoi2)

}

func describePeople(p people) {
	fmt.Println(p.stdin())
}

func (p person) stdin() string {
	p.firstname = "ngoc"
	p.lastname = "nhu"
	fullname := p.firstname + p.lastname
	return fullname
}

func (m string) stdin() string {

	m = "occhjo"
	return m
}

func (p *person) UpdateFirstName(newFirstName string) {
	p.firstname = newFirstName
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
