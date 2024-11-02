package main

import "fmt"

func main() {
	todo := Todos{}
	todo.Add("Coding")
	todo.Add("Rio")

	fmt.Printf("%v", todo)
}
