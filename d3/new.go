package main

import "fmt"

func con_brek() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // пропускає 2
		}
		if i == 4 {
			break // вихід з циклу
		}
		fmt.Println("Number: ", i)
	}
}

func for_range() {
	//Перебір елементів масиву, зрізу, рядка, мапи або каналу.
	// i = index, v = значення
	nums := []int{100, 200, 300, 400, 500}
	for i, v := range nums {
		fmt.Println(i, v)
	}
}

func ignore() {
	nums := []int{1, 2, 3, 4, 5}
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	con_brek()
	for_range()
	ignore()
}
