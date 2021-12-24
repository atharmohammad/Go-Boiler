package main

import "fmt"

func main() {
	colors := map[string]string{
		"White": "#fff",
		"Black": "#0000",
	}
	bottle := make(map[int]string)
	bottle[1] = "Nayasa"
	bottle[2] = "HackerEarth"
	delete(bottle, 1)
	print(colors)
	fmt.Println(bottle)
}

func print(m map[string]string) {
	for col, hex := range m {
		fmt.Println("Color is ", col, "and its hex is ", hex)
	}
}
