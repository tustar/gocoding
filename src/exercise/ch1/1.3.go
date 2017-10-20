package main

import (
	"os"
	"fmt"
)

func main()  {
	for index, arg := range os.Args[0:] {
		fmt.Println(index, arg)
	}
}
