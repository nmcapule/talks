package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	dat, err := ioutil.ReadFile("./main.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
}
