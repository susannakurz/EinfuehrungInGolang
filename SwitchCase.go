package main

import "fmt"

var zaehler int = 0;

func main() {
	for{
		switch zaehler{
		case 1
			fmt.Println("Eins")
		case 2
			fmt.Println("Zwei")
		case 3
			fmt.Println("Drei")
		case 4
			fmt.Println("Vier")
		}
	}
}
