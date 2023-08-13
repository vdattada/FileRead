package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	filecontents, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("Couldnt read file %v", err)
	}

	//str := string(filecontents)
	//fmt.Println(str)

	a := bytes.Split(filecontents, []byte("."))

	//fmt.Println(len(string(a[0])))
	//fmt.Println(string(a[0]))

	// for i := 0; i < len(a); i++ {
	// 	i := i
	// 	s := string(a[i])
	// 	fmt.Println(s)
	// }

	for _, e := range a {

		x := bytes.Split(e, []byte("\n"))
		for _, r := range x {
			for _, p := range r {
				fmt.Println(string(p))
			}
			fmt.Println(string(r))
		}
		fmt.Println(string(e))
	}

}
