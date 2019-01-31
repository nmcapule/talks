package main

import (
	"fmt"
	"syscall/js"
)


func myFunc(v []js.Value) {
	ele := js.Global().Get("document").Call("getElementById", "txt")
	fmt.Println("Input has value:", ele.Get("value"))
}

func main() {
	btn := js.Global().Get("document").Call("createElement", "button")
	btn.Set("textContent", "Hello my button")
	btn.Set("id", "btn")
	btn.Set("onclick", js.NewCallback(myFunc))

	txt := js.Global().Get("document").Call("createElement", "input")
	txt.Set("value", "")
	txt.Set("id", "txt")

	js.Global().Get("document").Get("body").Call("appendChild", txt)
	js.Global().Get("document").Get("body").Call("appendChild", btn)

	// wait, or else the callbacks won't execute.
	wait := make(chan bool)
	<-wait
}
