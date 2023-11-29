package main

//import "syscall/js"

func add(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
}

func main() {
	println("Hello world")
}
