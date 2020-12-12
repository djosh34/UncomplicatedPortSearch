package main

import (
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
		case "list": list()
		case "find": find()
		case "killall": killall()
	}



}

func list() {
	fmt.Println("list")
}

func find() {
	fmt.Println("find")

}

func killall() {
	fmt.Println("killall")

}
