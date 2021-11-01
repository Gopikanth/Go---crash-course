package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const enter = "and press enter" //reuse of code is done here

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstnumber = rand.Intn(8) + 2 //here we are using 8 +2 because  we need values from two  only
	var secondnumber = rand.Intn(8) + 2
	var subraction int = rand.Intn(8) + 2
	var answer = firstnumber*secondnumber - subraction
	game(firstnumber, secondnumber, subraction, answer)
}

func game(firstnumber, secondnumber, subraction, answer int) {

	space := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome To the Game")
	fmt.Println("------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 to 10", enter)
	space.ReadString('\n')
	fmt.Println("Multiply the number by ", firstnumber, enter)
	space.ReadString('\n')
	fmt.Println("Multiply the result by ", secondnumber, enter)
	space.ReadString('\n')
	fmt.Println("Divide the number by what you actually thought", enter)
	space.ReadString('\n')
	fmt.Println("Now subract", subraction, enter)
	space.ReadString('\n')
	fmt.Println("The answer is", answer)

}

//conclusion
//one function one purpose
