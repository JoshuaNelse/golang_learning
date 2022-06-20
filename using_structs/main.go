package main

import "fmt"

type MyType struct {
	myAttribute string
}

func (p *MyType) myPointerMethod() string {
	return p.myAttribute
}

func (v MyType) myValueMethod() string {
	return v.myAttribute
}

func main() {
	var myInstance MyType = MyType{"attr"}
	fmt.Println(myInstance.myPointerMethod())

	// testing object reference/dereference
	fmt.Println(myInstance)
	a := &myInstance
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(**b)
	fmt.Println((**b).myPointerMethod())
	fmt.Println((&(**b)).myPointerMethod())

	// testing auto derefence/reference on different recievers
	newInstance := MyType{"newAttr"}
	// by value (they both work)
	fmt.Println(newInstance.myPointerMethod())
	fmt.Println(newInstance.myValueMethod())

	// by pointer (also works)
	var newPointer *MyType = &newInstance
	fmt.Println(newPointer.myPointerMethod())
	fmt.Println(newPointer.myValueMethod())
}
