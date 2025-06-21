package main

import "fmt"

var input int

func main() {
	fmt.Print("Input: ")
	fmt.Scan(&input)
	result := Sloane(input)

	fmt.Println("Output:")
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i < len(result)-1 {
			fmt.Print("-")
		}
	}

}
func Sloane(input int) []int {
	arr := []int{}

	for i := 0; i < input; i++ {
		res := i*(i+1)/2 + 1

		arr = append(arr, res)
	}
	return arr
}
