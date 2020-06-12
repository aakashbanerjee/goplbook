package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	var s, sep string
	fmt.Println(len(os.Args))
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	//ch1 ex1.3 print execution time
	fmt.Println(time.Since(t1))
	t2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(t2))
	fmt.Println(os.Args[1:])
	//ch1 ex1.1 print os.Args[0] that is the name of the command that invoked it
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[0:], " "))
	//ch1 ex1.2 print the index and value of each of the arguments one per line
	t3 := time.Now()
	for index, value := range os.Args {
		fmt.Printf("%v\t%v\n", index, value)
	}
	fmt.Println(time.Since(t3))
}
