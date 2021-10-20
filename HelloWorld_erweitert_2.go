package main
import "fmt"

var hello string ="Hello world"

func main() {
	for true {
		for i :=len(hello); i >0; i-- {
			for x:=0; x<i; x++{
				fmt.Print(string(hello[x]))
			}
			fmt.Println("!")
		}
		for i :=0; i<len(hello); i++ {
			for x:=0; x<i; x++{
				fmt.Print(string(hello[x]))
			}
			fmt.Println("!")
		}
	}
}
