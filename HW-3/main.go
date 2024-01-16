package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	m := make([]int, n, n)
	for i, _ := range m {
		fmt.Scan(&m[i])
	}
	res := insertSort(n, m)
	for i, _ := range res {
		fmt.Print(res[i], " ")
	}

}
func insertSort(n int, m []int) []int {
	for i := 0; i < len(m); i++ {
		key := m[i]
		j := i - 1
		for j >= 0 && m[j] > key {
			m[j+1] = m[j]
			j--
		}
		m[j+1] = key
	}
	return m
}
