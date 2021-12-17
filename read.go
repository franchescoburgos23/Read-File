package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type User struct {
	fname string
	lname string
}

func Space() {
	fmt.Println("")
	fmt.Println("    *****************************************************************************   ")
	fmt.Println("")
}

func Regars() {
	fmt.Println("Hi I am divide the first name and last name")
}

func Goodbye() {
	fmt.Println("Thank you for used me bye :D")
}

func main() {

	Space()
	Regars()
	Space()
	var url string
	fmt.Println("Please enter a url")
	fmt.Scan(&url)
	Space()

	fil, err := os.Open(url)

	if err != nil {
		log.Fatal(err.Error())

	}

	scanner := bufio.NewScanner(fil)

	for scanner.Scan() {

		line := strings.Fields(scanner.Text())

		if len(line) > 20 {
			log.Fatal(err.Error())
		}
		structuser := make([]User, 0, 3)

		structuser = append(structuser, User{

			fname: line[0],
			lname: line[1],
		})

		for _, nrange := range structuser {
			fmt.Println("the first name is", nrange.fname, "and the last name is ", nrange.lname)
		}

	}
	Space()
	Goodbye()

}
