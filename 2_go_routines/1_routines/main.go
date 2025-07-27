package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello(("Hello"))
	sayHello("World")
}

func sayHello(msg string) {
	for range 5 {
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}
