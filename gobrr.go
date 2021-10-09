package main
import (
	"fmt"
)

func brr() {
	for {
		fmt.Println("brr");
	}
}

func main() {
	go brr()
	for {
	}
}
