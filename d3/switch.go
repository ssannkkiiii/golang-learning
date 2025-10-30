package main

import "fmt"

func main() {
	var day string

	fmt.Print("Enter day (Mon, Tue, Wed, Thu, Fri, Sat, Sun): ")
	fmt.Scan(&day)

	switch day {
	case "Mon":
		fmt.Println("Start of the week")
	case "Fri":
		fmt.Println("Weekend soon!")
	case "Sat", "Sun":
		fmt.Println("Weekend 🎉")
	default:
		fmt.Println("Midweek grind 😎")
	}
}
