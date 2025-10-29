package main

import "fmt"

//len - Функція для отримання кількості елементів слайса або масиву
//cap - Показує загальну місткість слайса без виділення нового масиву.

func lenANDcap() {
	sl := []int{1, 2, 3, 4, 5}

	all := make([]int, 3, 5)

	lens := len(all)
	caps := cap(all)

	// очистка слайсу
	sl = sl[:0]
	sl_cap := cap(sl)

	sl2 := []int{1, 2, 3, 4, 5}
	sl2 = nil
	sl_cap2 := cap(sl2)

	fmt.Println("Slice len: ", lens)
	fmt.Println("Slice cap: ", caps) //Якщо append перевищує cap, Go виділяє новий масив і копіює елементи.

	fmt.Println(sl, sl_cap, sl2, sl_cap2)

	s := []int{1, 2, 3, 4, 5}
	sub := s[1:4] // [2 3 4]
	sub2 := s[:3] // [1 2 3]
	sub3 := s[2:] // [3 4 5]

	fmt.Println(sub, sub2, sub3)

}

func slices() {
	sl := []int{1, 2, 3} //slice
	sl2 := append(sl, 4, 5)
	fmt.Println(sl2)
}

func arrays() {
	arr := [3]int{1, 2, 3} //array

	fmt.Println(arr)
}

func main() {
	//arrays()
	//slices()
	lenANDcap()
}
