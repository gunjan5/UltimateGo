//Declare an interface named speaker with a method named speak. Declare a struct named english that represents a person who
//speaks english and declare a struct named chinese for someone who speaks chinese. Implement the speaker interface for each
//struct using a value receiver and these literal strings "Hello World" and "你好世界". Declare a variable of type speaker and
//assign the address of a value of type english and call the method. Do it again for a value of type chinese.

//Add a new function named sayHello that accepts a value of type speaker. Implement that function to call the speak
//method on the interface value. Then create new values of each type and use the function.

package main

import (
	"fmt"
)

type speaker interface {
	speak()
}

type english struct{}

func (e english) speak() {
	fmt.Println("Yooo")
}

type chinese struct{}

func (c chinese) speak() {
	fmt.Println("你好世界")
}

func sayHello(s speaker) {
	s.speak()
}

func main() {
	e := english{}
	c := chinese{}
	var s speaker

	s = e

	s.speak()

	e.speak()
	c.speak()

	println("********")

	sayHello(s)
	sayHello(e)
	sayHello(c)

}
