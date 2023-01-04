package main

import "fmt"

func main() {
	var palabra string

	fmt.Scan(&palabra)

	fmt.Println(len(palabra))
	for _, letra := range palabra {
		fmt.Println(string(letra))

	}

}
