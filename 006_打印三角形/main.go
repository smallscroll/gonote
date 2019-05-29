package main

import "fmt"

func main() {

	/*
			*
		   ***
		  *****
		 *******
		*********

	*/

	line := 5
	for i := 0; i <= line; i++ {
		for k := i; k < line; k++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i*2-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for m := 0; m < line-1; m++ {
		for o := 0; o < m+1; o++ {
			fmt.Printf(" ")
		}
		for n := m*2 - 1; n < line*2-4; n++ {
			fmt.Printf("*")
		}
		fmt.Println("")

	}

}
