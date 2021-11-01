package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard" //here we using third party packages
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()
	flavour := make(map[int]string)
	flavour[1] = "chocolate"
	flavour[2] = "bhadam"
	flavour[3] = "Pineapple"
	flavour[4] = "Mint"
	flavour[5] = "Lemon"

	fmt.Println("Menu")
	fmt.Println("---------")
	fmt.Println("1 --- chocolate")
	fmt.Println("2 --- bhadam")
	fmt.Println("3 --- Pineapple")
	fmt.Println("4 --- Mint")
	fmt.Println("5 --- lemon")
	fmt.Println("Q --- Quit menu")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'Q' || char == 'q' {
			break
		}
		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("you choosed %s", flavour[i])
		fmt.Println(t)

	}

	fmt.Println("Program exiting...")
}
