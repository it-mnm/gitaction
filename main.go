package main

import(
	"fmt"
	"time"
)

func main(){
	for{
		fmt.Println("Hello Cloud")
		time.Sleep(10*time.Second)
	}
}