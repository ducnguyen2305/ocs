package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("A---------")
	<-time.After(2 * time.Second)
	fmt.Println("B---------")
}
