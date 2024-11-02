package main

func main() {
	todo := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todo)
	cmdFlag := NewCmdFlags()
	cmdFlag.Execute(&todo)
	storage.Save(todo)
}
