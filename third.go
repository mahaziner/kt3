package main

import (
	"fmt"
)

func main() {
	chane := make(chan int) 

	chane <- 42

	fmt.Println("Это сообщение никогда не будет напечатано")
}
