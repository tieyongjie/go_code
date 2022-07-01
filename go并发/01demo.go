package main

import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Println(s)
	}

}

func main() {
	go say("World")
	say("hello")
}
