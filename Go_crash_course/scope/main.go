// CASE - 1
/*package main

import "fmt"

func main() {
	var a = "i am a boy"
	fmt.Println(a)
	next()
}
func next() {
	var a = "i am a girl"
	fmt.Println(a)
}

 the variable a in main and  variable a in next are not same
 They are called BLOCK LEVEL VARIABLES

*/
//CASE -2
/*package main

import "fmt"
var a = "i am a boy"

func main() {

	fmt.Println(a)
	next()
}
func next() {
	fmt.Println(a)
 now the variable a is known as PACKAGE LEVEL VARIABLES
}*/
//CASE -3

/*package main

import "fmt"

var a = "i am a boy"

func main() {
	a := " this is block level variable"
	fmt.Println(a)
	next()
}
func next() {
	fmt.Println(a)
}*/

//CASE -4

/*package main

import "fmt"

var a = "i am a boy"

func main() {
	var blockLevel = "HELLO BLOCK" //using camelCase is convention for variables
	fmt.Println(blockLevel)
	next()
}
func next() {
	fmt.Println(a)
}

//In case-4  we use blockLevel variable because it avoids shadowing of variables*/

//CASE -5
package main

import (
	"fmt"
	"myapp/packageone"
)

var a = "i am a boy"

func main() {
	var b = packageone.B
	fmt.Println(b)
	next()
}
func next() {
	fmt.Println(a)
}

//in this case 5 we created our own package
//we can conlude that that a in package is not exported because of lowercase where B is exported due to package is exported
