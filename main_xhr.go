package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	res, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))
}
