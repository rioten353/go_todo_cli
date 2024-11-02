package main

// import "fmt"

func main() {
	todo := Todos{}
  storage := NewStorage[Todos]("todos.json")
  storage.Load(&todo)
	todo.Add("Coding")
	todo.Add("Rio")
  todo.Toggle(0)
  todo.Print()
  storage.Save(todo)
	// fmt.Printf("%v", todo)
	// todo.Delete(1)
	// fmt.Printf("%v", todo)
}
