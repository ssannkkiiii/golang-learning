package main

import "fmt"

const name string = "Oleksandr"
const surname string = "Hredil"
const age int = 19

func name_check() {
	if name == "Oleksandr" {
		fmt.Println("Yes it's my name!")
	} else {
		fmt.Println("Error!")
	}
}

func age_check() {
	if age == 19 {
		fmt.Println("Yes it's me age!")
	} else {
		fmt.Println("Error!")
	}
}

func new_datas() {
	if n := len(name); n < 10 {
		fmt.Println("All done")
	} else {
		fmt.Println("Error")
	}
}

// 1 type of For
func type_1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func dead_inside() {
	for i := 1000; i > 7; i -= 7 {
		fmt.Println(i)
	}
}

// 2 type of For, analog while
func type_2() {
	i := 5
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

// 3 type of For, infinity
func type_3() {
	var a string = "Text"
	for {
		a = a + a
		fmt.Println("Forever")
	}
}

func main() {
	//name_check()
	//age_check()
	//new_datas()

	//type_1()
	//type_2()
	//type_3()
	//dead_inside()
}
