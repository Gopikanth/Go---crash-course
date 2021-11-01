package packageone

import "fmt"

var PackageVar = "This is in other package"

func Printme(s1, s2 string) {
	fmt.Println(s1, s2)
}
