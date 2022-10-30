package main

import "fmt"

func main() {
	fmt.Println("### Basic Concepts ###")

	//variables
	var city string = "New York"
	var country string = "USA"
	var age int = 24
	var height float32 = 1.91
	var isRain bool = false

	fmt.Println("COUNTRY:")
	fmt.Println(country)
	fmt.Println("CITY:")
	fmt.Println(city)
	fmt.Println("AGE:")
	fmt.Println(age)
	fmt.Println("HEIGHT")
	fmt.Println(height, "m")
	fmt.Println("is raining")
	fmt.Println(isRain)

	//constants
	const pi = 3.14
	const color = "red"
}
