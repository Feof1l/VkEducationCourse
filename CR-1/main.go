package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	m := make([]int, n, n)
	for i, _ := range m {
		fmt.Scan(&m[i])
	}
	res := make([]int, n, n)
	res = swap(n, m)
	for i, _ := range res {
		fmt.Print(res[i], " ")
	}

}
func swap(n int, mas []int) []int {
	evenIndex := 0
	for i, _ := range mas {
		if mas[i]%2 == 0 {
			elem := mas[evenIndex]
			mas[evenIndex] = mas[i]
			mas[i] = elem
			evenIndex++
		}
	}
	return mas
}

//Дан не отсортированный массив целых чисел. Необходимо перенести в начало массива все четные числа.
//При этом последовательность четных чисел должна быть сохранена. То есть если 8 стояла после 6, то после переноса в начало, она по-прежнему должна стоять после 6.
